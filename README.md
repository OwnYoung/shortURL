# 短链服务

一个简洁高效的短链接生成与管理服务，支持自定义过期时间和密码保护，适合个人或团队内部使用。

<p align="left">
  <img alt="License" src="https://img.shields.io/badge/license-MIT-blue.svg">
  <img alt="Go Version" src="https://img.shields.io/badge/go-1.18+-brightgreen">
  <img alt="Build Status" src="https://img.shields.io/github/actions/workflow/status/your_username/your_repo/ci.yml?branch=main">
</p>

## 目录

- [特性](#特性)
- [安装](#安装)
- [快速开始](#快速开始)
- [使用指南](#使用指南)
- [目录结构](#目录结构)
- [API 说明](#api-说明)
- [如何贡献](#如何贡献)
- [许可证](#许可证)
- [作者/贡献者](#作者贡献者)
- [致谢](#致谢)

## ✨ 特性

- 支持短链接生成与跳转
- 可设置按小时精度的过期时间（ISO 8601 格式）
- 支持密码保护短链
- 简洁的前端页面
- SQLite 本地存储，易于部署
- RESTful API 设计

## 📦 安装

1. 克隆仓库
   ```bash
   git clone https://github.com/your_username/shortURL.git
   cd shortURL
   ```
2. 构建
   ```bash
   go build -o shortURL
   ```
3. （可选）直接使用已编译的 `shortURL.exe`

## 🚀 快速开始

1. 启动服务
   ```bash
   ./shortURL
   ```
2. 访问 [http://localhost:8080](http://localhost:8080) 打开前端页面

## 📖 使用指南

- **生成短链**：在首页输入原始链接、可选的过期时间和密码，点击生成
- **访问短链**：通过生成的短链访问，若设置了密码则需输入密码
- **过期机制**：短链到期后自动失效，无法访问

## 🗂️ 目录结构

```
shortURL
├─ README.md
├─ controllers
│  └─ shortlink_controller.go
├─ go.mod
├─ go.sum
├─ main.go
├─ models
│  └─ shortlink.go
├─ shortURL.exe
├─ static
│  ├─ index.html
│  ├─ password_prompt.html
│  └─ shortlink.html
├─ storage
│  └─ storage.go
└─ test.db

```

## 🛠️ API 说明

### 创建短链

- `POST /api/shorten`
  - 请求参数（JSON）:
    - `url` (string): 原始链接
    - `expire_at` (string, 可选): 过期时间（ISO 8601 格式）
    - `password` (string, 可选): 访问密码
  - 返回:
    - `short_url` (string): 生成的短链

### 跳转短链

- `GET /s/{shortcode}`
  - 跳转到原始链接
  - 若设置了密码，需通过验证

> 具体参数和返回格式详见代码注释

## 🤝 如何贡献

我们非常欢迎各种贡献！请阅读我们的贡献指南。

1. Fork 本项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开一个 Pull Request

请阅读 `CONTRIBUTING.md` 了解详细的行为准则和提交流程。

## 📜 许可证

本项目基于 MIT 许可证发行。详情请见 LICENSE 文件。

## 👥 作者/贡献者

- Orion Young - 初始工作 

感谢所有为此项目做出贡献的人！(contributors)

## 🙏 致谢

- 灵感来源 某个项目
- 感谢 某个人 提出的宝贵建议
- 使用的图标来自 某网站
