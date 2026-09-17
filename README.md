# lost-and-found-backend
精弘试用期大作业后端
lost‑and‑found‑back/  # 项目根目录
├── app/
│   ├── controllers/   # 控制器：接收http请求，调用service，返回响应
│   ├── middlewares/   # 中间件：登录校验、跨域、日志、权限
│   ├── models/       # 数据库模型结构体（对应数据库表）
│   ├── services/      # 业务逻辑层，处理核心业务
│   └── utils/         # 工具函数：加密、时间、字符串处理等
├── configs/
│   ├── config/        # 配置加载代码
│   └── database/      # 数据库初始化连接代码
├── .gitignore
├── config.yaml.example # 配置模板，真正配置叫config.yaml（git不提交真实配置）
├── go.mod
├──main.go              #唯一入口，启动
└── README.md