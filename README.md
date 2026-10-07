<div align="center">

# lost-and-found-backend

**精弘试用期大作业 · 校园失物招领系统后端**

<img src="https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go">
<img src="https://img.shields.io/badge/Gin-v1.9-4EC08D?style=flat-square&logo=gin&logoColor=white" alt="Gin">
<img src="https://img.shields.io/badge/GORM-v1.25-6C5CE7?style=flat-square&logo=gorm&logoColor=white" alt="GORM">
<img src="https://img.shields.io/badge/MySQL-5.7+-4479A1?style=flat-square&logo=mysql&logoColor=white" alt="MySQL">

校园失物招领系统后端服务：用户注册登录、失物/招领帖发布与审核、认领申请、评论、联系人、公告、禁言与系统管理。

共 **35 个接口**，与 Apifox 文档保持一致。

</div>

---

## 快速开始

1. **环境要求**：Go 1.20+，MySQL 5.7+
2. **创建数据库**（表结构启动时自动迁移创建）：

   ```sql
   CREATE DATABASE lostfound CHARACTER SET utf8mb4;
   ```

3. **配置**：复制 `config.example.yaml` 为 `config.yaml`，填写数据库账号密码、`jwt.key`、管理员邀请码、`upload.base_url`
4. **启动**：

   ```bash
   go run main.go
   # 或编译后运行
   go build -o server main.go && ./server
   ```

5. 服务默认监听 `8080`，上传图片保存在本地 `./images` 目录，通过 `/images` 路径静态访问

## 目录结构

```
lost-and-found-backend/
├── app/
│   ├── controllers/   # 控制器：接收请求 → 调用 service → 返回响应
│   │   ├── user_controller/          # 用户
│   │   ├── post_controller/          # 帖子、图片上传
│   │   ├── claim_controller/         # 认领申请
│   │   ├── comment_controller/       # 评论
│   │   ├── contact_controller/       # 联系人
│   │   ├── announcement_controller/  # 公告
│   │   ├── admin_controller/         # 失物招领管理员
│   │   └── sysadmin_controller/      # 系统管理员
│   ├── middlewares/   # JWT 校验、角色权限、禁言校验、统一错误响应、登录限流
│   ├── models/        # 数据库模型（对应数据表）
│   ├── services/      # 业务逻辑层
│   ├── errs/          # 统一业务错误码定义
│   └── utils/         # 工具：JWT、密码加密、时间等
├── configs/
│   ├── config/        # 配置加载
│   ├── database/      # 数据库初始化与自动建表
│   └── router/        # 路由注册
├── images/            # 上传图片存储目录
├── config.example.yaml # 配置模板（真实配置 config.yaml 不入库）
├── main.go            # 唯一入口
└── README.md
```

## 角色与权限

| 角色 | 说明 | 注册方式 |
| --- | --- | --- |
| 普通用户 | 发布帖子、提交认领、评论、管理联系人 | 直接注册 |
| 失物招领管理员 | 审核帖子、删除帖子、修改帖子解决状态 | 注册时携带管理员邀请码 |
| 系统管理员 | 用户角色管理、禁言/解禁、公告、系统数据 | 注册时携带管理员邀请码 |

注册时 `role` 与 `invite_code` 共同决定身份：普通用户无需邀请码，管理员角色必须提供邀请码。

## API 一览

统一前缀 `/api`。除标注"公开"的接口外，均需请求头 `Authorization: Bearer <token>`。

### 公开接口

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| POST | `/api/register` | 注册 |
| POST | `/api/login` | 登录 |
| GET | `/api/posts` | 公开帖子列表（仅已通过） |
| GET | `/api/posts/:post_id/comments` | 获取评论 |
| GET | `/api/posts/:post_id` | 帖子详情（**可选登录**：游客看已通过帖，作者可看自己待审核帖） |
| GET | `/api/announcements` | 查看公告 |

### 用户

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| GET | `/api/user/profile` | 获取个人信息 |
| PATCH | `/api/user/profile` | 修改个人信息 |
| PATCH | `/api/user/password` | 修改密码 |

### 失物招领帖子

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| POST | `/api/upload` | 上传图片 |
| POST | `/api/posts` | 发布帖子（默认待审核） |
| GET | `/api/my/posts` | 查询自己发布的帖子 |
| PUT | `/api/my/posts/:post_id` | 编辑并重新提交（回到待审核） |
| DELETE | `/api/my/posts/:post_id` | 删除自己的帖子 |

### 认领申请

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| POST | `/api/posts/:post_id/claims` | 提交认领申请 |
| GET | `/api/posts/:post_id/claims` | 查看自己帖子收到的申请 |
| GET | `/api/my/claims` | 查看我提交的申请 |
| PUT | `/api/claims/:claim_id/audit` | 处理认领申请（同意/拒绝） |
| DELETE | `/api/claims/:claim_id` | 撤回申请（仅待处理可撤回） |

### 联系人

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| POST | `/api/contact` | 添加联系人 |
| GET | `/api/my/contacts` | 查询联系人列表 |
| DELETE | `/api/my/contacts/:contact_id` | 删除联系人 |

### 评论

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| POST | `/api/posts/:post_id/comments` | 发布评论 |
| DELETE | `/api/posts/:post_id/comments/:comment_id` | 删除自己的评论 |

### 失物招领管理员

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| GET | `/api/admin/posts` | 查询所有帖子（含未审核） |
| PATCH | `/api/admin/posts/:post_id/audit` | 审核帖子（通过/拒绝） |
| DELETE | `/api/admin/posts/:post_id` | 删除帖子 |
| PATCH | `/api/admin/posts/:post_id/resolve` | 修改帖子的解决状态 |

### 系统管理员

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| GET | `/api/sys/users` | 查询所有用户 |
| PATCH | `/api/sys/users/:user_id/role` | 修改用户角色 |
| PATCH | `/api/sys/users/:user_id/mute/:mute_second` | 禁言用户（按秒） |
| PATCH | `/api/sys/users/:user_id/unmute` | 解除禁言 |
| POST | `/api/sys/announcements` | 发布公告 |
| DELETE | `/api/sys/announcements/:announcement_id` | 删除公告 |
| GET | `/api/sys/stats` | 获取系统数据 |

## 统一响应格式

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

`code = 0` 表示成功；非 0 为业务错误码。错误时同时返回 HTTP 状态码与业务错误码。

## 错误码

### 用户相关（1xxx）

| HTTP | 业务码 | 说明 |
| --- | --- | --- |
| 400 | 1001 | 数据获取失败（参数缺失/格式错误/用户名或密码为空） |
| 409 | 1002 | 用户名或手机号已存在 |
| 500 | 1003 | 用户信息查询失败 |
| 500 | 1004 | 密码加密失败 |
| 500 | 1005 | 数据库出错 |
| 403 | 1006 | 没有权限 |
| 404 | 1007 | 用户不存在 |
| 401 | 1008 | 密码错误 |
| 500 | 1009 | 登录令牌生成失败 |
| 401 | 1010 | 未登录或无效的 token |
| 401 | 1011 | 登录过期 |
| 400 | 1012 | 手机号格式不正确 |
| 400 | 1013 | 新旧密码相同 |
| 400 | 1014 | 旧密码错误 |
| 400 | 1015 | 用户名限制长度为 50 字符 |
| 429 | 1016 | 登录失败次数过多，账号已锁定，请 15 分钟后再试 |
| 403 | 1017 | 账号处于禁言状态，禁止发布内容 |

### 帖子相关（2xxx）

| HTTP | 业务码 | 说明 |
| --- | --- | --- |
| 400 | 2001 | 未接收到上传文件 |
| 400 | 2002 | 文件上传失败 |
| 400 | 2003 | 文件格式不正确，仅支持 jpg、jpeg、png |
| 400 | 2004 | 图片大小不能超过 5MB |
| 400 | 2005 | 路径参数格式错误 |
| 404 | 2006 | 帖子不存在 |
| 403 | 2007 | 这不是你的帖子 |
| 409 | 2008 | 当前状态不允许此操作 |
| 409 | 2009 | 当前状态不允许编辑 |

### 联系人相关（3xxx）

| HTTP | 业务码 | 说明 |
| --- | --- | --- |
| 404 | 3001 | 联系人不存在 |
| 403 | 3002 | 这不是你的联系人 |

### 评论相关（4xxx）

| HTTP | 业务码 | 说明 |
| --- | --- | --- |
| 404 | 4001 | 评论不存在 |
| 403 | 4002 | 这不是你的评论 |

### 公告相关（5xxx）

| HTTP | 业务码 | 说明 |
| --- | --- | --- |
| 404 | 5001 | 公告不存在 |

### 认领相关（6xxx）

| HTTP | 业务码 | 说明 |
| --- | --- | --- |
| 409 | 6001 | 该物品已被认领/解决 |
| 403 | 6002 | 不能认领自己发布的帖子 |
| 409 | 6003 | 您已提交过认领申请 |
| 404 | 6004 | 认领申请不存在 |

## 配置说明（config.yaml）

| 配置段 | 字段 | 说明 |
| --- | --- | --- |
| `database` | `name` / `host` / `port` / `user` / `pass` | MySQL 连接信息 |
| `jwt` | `key` | JWT 签名密钥 |
| `jwt` | `issuer` | 令牌签发者 |
| `register` | `admin_invite_code` | 管理员注册邀请码 |
| `upload` | `base_url` | 图片访问基础地址（返回给前端的图片完整 URL 前缀） |

## 核心业务规则

- **帖子状态流**：`待审核` → 管理员审核 → `已通过` / `已拒绝`；编辑重新提交后回到 `待审核`
- **详情可见性**：游客仅可查看 `已通过` 帖；作者携带 token 可查看自己的 `待审核` 帖；管理员可查看全部
- **认领规则**：仅 `已通过` 且 `未解决` 的帖子可提交认领；同一用户对同一帖子只能有一条 `待处理` 申请；不能认领自己发布的帖子；并发提交由数据库事务 + 行锁保证不产生重复申请
- **认领状态流**：`待处理` → 发布者同意 → 帖子标记为 `已解决`；被拒绝后可重新提交
- **禁言**：系统管理员可按秒禁言/解禁；被禁言用户不能上传图片、发帖、认领、评论；到期自动解除
- **删除评论**：路径中的 `post_id` 必须与评论实际所属帖子一致，否则返回 `400/2005`
- **登录限流**：连续登录失败达到阈值后账号锁定 15 分钟（`429/1016`）

## 部署提示

- `config.yaml` 已在 `.gitignore` 中，真实配置不会提交到仓库
- 上传图片保存在本地 `./images` 目录，部署到服务器时注意目录的读写权限与磁盘空间
