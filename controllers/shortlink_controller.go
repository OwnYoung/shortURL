package controllers

import (
	"crypto/sha256"
	"encoding/base64"

	// "strconv"

	"strings"

	"math/rand"
	"time"

	"github.com/OwnYoung/shortURL/models"
	"github.com/OwnYoung/shortURL/storage"
	"github.com/gin-gonic/gin"
)

// CreateShortLink 函数：处理创建请求，生成短码，存入数据库。
func CreateShortLink(c *gin.Context) {
	// POST /post?url=xxxx HTTP/1.1
	link := c.PostForm("url")
	password := c.PostForm("password")
	expiresAt := c.PostForm("expires_at") // 直接传入iso格式过期时间戳

	if link == "" || expiresAt == "" {
		c.JSON(400, gin.H{"error": "Missing required parameters"})
		return
	}

	// 查询原链接是否已存在
	var existingLink models.ShortLink
	flag := storage.DB.Where("original_url = ?", link).Take(&existingLink)
	if flag.Error == nil {
		c.JSON(200, gin.H{
			"message":      "Link already exists",
			"short_code":   existingLink.ShortCode,
			"original_url": existingLink.OriginalURL,
			"expires_at":   existingLink.ExpiresAt,
		})
		return
	}

	// 不存在，生成唯一短码
	shortCode := generateShortCode(link)
	var existing models.ShortLink
	result := storage.DB.Where("short_code = ?", shortCode).Take(&existing)
	for result.Error == nil {
		shortCode = generateShortCode(link)
		result = storage.DB.Where("short_code = ?", shortCode).Take(&existing)
	}

	newLink := models.ShortLink{
		ShortCode:   shortCode,
		OriginalURL: link,
		Password:    password,
		ExpiresAt:   expiresAt,
	}

	storage.DB.Create(&newLink)

	c.JSON(200, gin.H{
		"short_code":   shortCode,
		"original_url": link,
		"expires_at":   newLink.ExpiresAt,
		"created_at":   time.Now().Format("2006-01-02 15"),
	})
}

// RedirectShortLink 函数：根据短码查找原链接并重定向。
func RedirectShortLink(c *gin.Context) {
	// 根据短码找原链接
	shortCode := c.Param("shortCode")
	password := c.Query("password")

	// TODO: 添加密码验证
	var link models.ShortLink
	result := storage.DB.Where("short_code = ?", shortCode).Take(&link)
	if result.Error != nil {
		// 未找到对应的短码
		c.JSON(404, gin.H{"error": "Short link not found"})
		return
	}
	// 找到对应的短码，检查是否过期
	expireTime, err := time.Parse(time.RFC3339, link.ExpiresAt)
	if err != nil {
		c.JSON(500, gin.H{"error": "Invalid expires_at format"})
		return
	}
	if time.Now().After(expireTime) {
		c.JSON(403, gin.H{"error": "Short link has expired"})
		return
	}
	// 检查密码是否正确
	if link.Password != "" && link.Password != password {
		c.JSON(401, gin.H{"error": "Incorrect password"})
		return
	}

	// 校验结束，在有效期内，密码正确，可以302 重定向
	// c.Writer.WriteString(fmt.Sprintf("%s \n", "302重定向")) // 打印日志
	if !strings.HasPrefix(link.OriginalURL, "http") {
		link.OriginalURL = "http://" + link.OriginalURL
	}
	c.Redirect(302, link.OriginalURL)
}

// 生成短码的辅助函数
func generateShortCode(originalURL string) string {
	//
	salt := generateSalt(16) // You can adjust the salt length
	saltedURL := salt + originalURL
	hash := sha256.Sum256([]byte(saltedURL))
	base64Encoded := base64.URLEncoding.EncodeToString(hash[:]) // Encode to base64
	shortLink := base64Encoded[:8]                              // Take the first 8 characters (adjust length as needed)
	return shortLink
}

// Function to generate a random salt
func generateSalt(length int) string {
	// rand.Seed(time.Now().UnixNano()) go1.20后弃用
	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
