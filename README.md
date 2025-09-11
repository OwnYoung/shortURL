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

# ToDo添加过期时间
页面 按小时过期，后端存戳，比对时间
时间格式 2006-01-02 15 （仅精确到小时）
ISO 8601 字符串

如 2024-06-01T12:00:00Z 或 2024-06-01T12:00:00+08:00
-[ ] 添加过期时间

-[ ] 添加密码验证