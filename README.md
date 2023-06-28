# minigin
基于gin开发的MVC微型Go框架～

## 目录结构
```
minigin
├── conf // 项目配置
│   └── app.ini
├── controllers // Controller控制层
│   ├── activity.go // 秒杀活动demo控制层
│   ├── controller.go
│   └── sonbusiness // 按业务分Controller控制层
│       └── index.go
├── dao // 操作基类
│   └── database.go
├── docs // 文档文本
├── go.mod
├── go.sum
├── library // 基础库
│   ├── e   // 错误码及信息
│   │   ├── code.go
│   │   └── msg.go
│   ├── file // 文本
│   │   └── file.go
│   ├── logging // 日志
│   │   └── log.go
│   ├── middleware // 中间件
│   │   └── jwt.go
│   ├── redis // Redis
│   │   └── redis.go
│   ├── setting // 基础设置
│   │   └── setting.go
│   └── util // 常用方法
│       ├── md5.go
│       └── response.go
├── main.go // 主入口
├── models // 模型
│   └── article.go
├── README.md
├── routers // 路由分发
│   ├── router.go
│   └── sonrouter.go // 子路由分发
├── runtime // 系统运行日志
│   └── logs
├── script // 脚本
├── static // 静态资源文件
├── templates // 模版文件
│   └── index.html
└── tmp // Air相关日志与产出
```
