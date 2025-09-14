# 短链服务

一个简洁有力的描述，说明你的项目是什么以及它解决了什么问题。

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Python Version](https://img.shields.io/badge/python-3.8%2B-brightgreen)
![Build Status](https://img.shields.io/github/actions/workflow/status/your_username/your_repo/ci.yml?branch=main)

![image-20250914232857028](/Users/ouyangyi/Library/Application Support/typora-user-images/image-20250914232857028.png)



![image-20250914232820346](/Users/ouyangyi/Library/Application Support/typora-user-images/image-20250914232820346.png)

> **提示**: 这是一个附加的、更详细的说明或“电梯演讲”。


## ✨ 特性



## 📦 安装



## 🚀 快速开始



## 📖 使用指南





## 🤝 如何贡献

我们非常欢迎各种贡献！请阅读我们的贡献指南。

1. Fork 本项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开一个 Pull Request

请阅读 [CONTRIBUTING.md](https://contributing.md/) 了解详细的行为准则和提交流程。

## 📜 许可证

本项目基于 **MIT** 许可证发行。详情请见 [LICENSE](https://license/) 文件。

## 👥 作者/贡献者

- **Orion Young** - *初始工作* - [YourUsername](https://github.com/yourusername)

感谢所有为此项目做出贡献的人！ ([contributors](https://github.com/yourusername/yourrepo/contributors)).

## 🙏 致谢

- 灵感来源 [某个项目](https://github.com/someproject)
- 感谢 [某个人](https://github.com/someone) 提出的宝贵建议。
- 使用的图标来自 [某网站](https://somewebsite.com/)。

```

shortURL/
├── main.go          # 程序入口
├── models/          # 数据模型
│   └── shortlink.go
├── controllers/     # 控制器（处理业务逻辑）
│   └── shortlink_controller.go
├── storage/         # 数据库层
│   └── storage.go
└── go.mod
```

ToDo添加过期时间

页面 按小时过期，后端存戳，比对时间
时间格式 2006-01-02 15 （仅精确到小时）
ISO 8601 字符串

如 2024-06-01T12:00:00Z 或 2024-06-01T12:00:00+08:00
-[ ] 添加过期时间

-[ ] 添加密码验证

### 版本v1.0914

前端更新：

- 用户未填写密码时，不会提交 password 字段。
- 展示短链时，只有有密码才显示密码信息。

后端修改逻辑：

1. 查询数据库时，where 条件要加上 password 是否为空（或 has_password 字段为 false）。
2. 如果查到的记录 password 不为空（即有密码），则不返回，直接新建一条 has_password=false 的短链。
