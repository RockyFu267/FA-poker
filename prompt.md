## 项目基础结构
aquila
├── .DS_Store
├── .gitignore
├── .idea
│   ├── .gitignore
│   ├── bot-server.iml
│   ├── modules.xml
│   ├── vcs.xml
│   └── workspace.xml
├── .vscode
│   └── settings.json
├── LICENSE
├── README.md
├── api
│   ├── bot
│   │   ├── qianxun.go
│   │   ├── spider.go
│   │   └── weixin.go
│   ├── common
│   │   ├── auth.go
│   │   ├── dto.go
│   │   └── vo.go
│   ├── minio
│   │   └── file.go
│   ├── system
│   │   ├── dept.go
│   │   ├── dto.go
│   │   ├── menu.go
│   │   ├── role.go
│   │   └── user.go
│   ├── wechat
├── config
│   ├── app.go
│   ├── config.go
│   ├── gorm_mysql.go
│   ├── gorm_pgsql.go
│   ├── minio.go
│   ├── redis.go
│   └── zap.go
├── config.release.yaml
├── config.yaml
├── constants
│   └── constants.go
├── enum
│   ├── accountTypeEnum.go
│   ├── codeEnum.go
│   ├── messageEnum.go
│   └── orderEnum.go
├── global
│   └── global.go
├── initialize
│   ├── createInitTable.go
│   ├── gorm.go
│   ├── gorm_mysql.go
│   ├── gorm_pgsql.go
│   ├── minio.go
│   ├── redis.go
│   ├── server.go
│   ├── viper.go
│   └── zap.go
├── middleware
│   ├── AuthMiddleWare.go
│   ├── CorsMiddleWare.go
│   ├── LoggerMiddleWare.go
│   └── RecoverMiddleWare.go
├── model
│   ├── dept.go
│   ├── menu.go
│   ├── role.go
│   ├── role_menu.go
│   ├── user.go
│   └── user_role.go
├── router
│   ├── bot.go
│   ├── common.go
│   ├── login.go
│   ├── menu.go
│   ├── role.go
│   └── user.go
├── scheduler
│   └── scheduler.go
├── utils
│   ├── captcha.go
│   ├── directory.go
│   ├── jwt.go
│   ├── paginate.go
│   ├── response.go
│   ├── strSplit.go
│   ├── utils.go
│   └── validator.go
└── main.go