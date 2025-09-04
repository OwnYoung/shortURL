package controllers

import (
	"crypto/sha256"
	"encoding/base64"

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
	// 处理创建请求，获得gin传来的原链接
	link := c.PostForm("url")

	//* 查询原链接是否存在
	var existingLink models.ShortLink
	flag := storage.DB.Where("original_url = ?", link).Take(&existingLink)
	if flag.Error == nil {
		// 查到数据，error为nil，说明link已存在
		// c.Writer.WriteString(fmt.Sprintf("%s %s\n", existingLink.ShortCode, existingLink.OriginalURL))
		c.JSON(200, gin.H{
			"message":      "Link already exists",
			"short_code":   existingLink.ShortCode,
			"original_url": existingLink.OriginalURL,
		})
		return
	}

	//* 若不存在则生成短码并存入数据库
	// 生成短码
	shortCode := generateShortCode(link)
	// 需要先检查数据库是否存在该短码，若存在则重新生成first/take方法
	var existing models.ShortLink
	result := storage.DB.Where("short_code = ?", shortCode).Take(&existing)

	for result.Error == nil {
		// 已经存在，重新生成
		shortCode = generateShortCode(link)
		result = storage.DB.Where("short_code = ?", shortCode).Take(&existing)
	}

	// 存入数据库
	newLink := models.ShortLink{
		ShortCode:   shortCode,
		OriginalURL: link,
		// CreatedAt: time.Now(),
	}
	storage.DB.Create(&newLink)

	// 返回结果
	c.JSON(200, gin.H{
		"short_code":   shortCode,
		"original_url": link,
	})

}

// RedirectShortLink 函数：根据短码查找原链接并重定向。
func RedirectShortLink(c *gin.Context) {
	// 根据短码找原链接
	shortCode := c.Param("shortCode")
	var link models.ShortLink
	result := storage.DB.Where("short_code = ?", shortCode).Take(&link)
	if result.Error != nil {
		// 未找到对应的短码
		c.JSON(404, gin.H{"error": "Short link not found"})
		return
	}
	// 302 重定向
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
