# lost-and-found-backend
精弘试用期大作业后端
lost-and-found-backend/     # 项目根目录
├── app/
│   ├── controllers/        # 控制器：接收http请求，调用service，返回响应
│   │   └── userController/ # 用户模块控制器：登录、注册
│   ├── errs/               # 全局错误定义：统一业务码、HTTP码与错误信息
│   ├── middlewares/        # 中间件：JWT解析、全局异常捕获与统一响应
│   ├── models/             # 数据库模型结构体（对应数据库表）
│   ├── services/           # 业务逻辑层，处理核心业务
│   └── utils/              # 工具函数：密码加密、JWT生成、成功响应封装
├── configs/                # 保存、读取配置与初始化
│   ├── config/             # 配置加载代码
│   ├── database/           # 数据库初始化连接代码
│   └── router/             # Gin路由注册
├── .gitignore
├── config.example.yaml     # 配置模板，真正配置叫config.yaml（git不提交真实配置）
├── go.mod
├── go.sum
├── main.go                 # 唯一入口，启动
└── README.md