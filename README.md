<div align="center">

# lost-and-found-backend

**精弘试用期大作业 · 校园失物招领系统后端**

<p>
<img src="https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go&logoColor=white" alt="Go">
<img src="https://img.shields.io/badge/Gin-Web%20Framework-008ECF" alt="Gin">
<img src="https://img.shields.io/badge/GORM-ORM-00A6FB" alt="GORM">
<img src="https://img.shields.io/badge/MySQL-5.7%2B-4479A1?logo=mysql&logoColor=white" alt="MySQL">
<img src="https://img.shields.io/badge/JWT-Auth-000000?logo=jsonwebtokens&logoColor=white" alt="JWT">
</p>

</div>

基于 **Go + Gin + GORM + MySQL** 的失物招领服务端，覆盖用户注册登录、帖子发布与审核、认领申请、评论、联系人、公告与系统管理全流程。数据表由 GORM `AutoMigrate` 自动创建，开箱即跑。

---

## 核心能力

<table>
  <tr>
    <td align="center" style="background:#eaf6ff;border-radius:8px;padding:12px;width:33%">
      <b>👤 用户体系</b><br>
      <span style="color:#555">注册 · 登录 · JWT 鉴权<br>bcrypt 加密 · 登录限流</span>
    </td>
    <td align="center" style="background:#eafaf0;border-radius:8px;padding:12px;width:33%">
      <b>📦 帖子闭环</b><br>
      <span style="color:#555">发布 → 审核 → 认领 → 解决<br>图片上传 · 公开列表</span>
    </td>
    <td align="center" style="background:#fff4e5;border-radius:8px;padding:12px;width:33%">
      <b>🔒 认领防重</b><br>
      <span style="color:#555">事务 + 行锁<br>并发提交不产生重复申请</span>
    </td>
  </tr>
  <tr>
    <td align="center" style="background:#f3ecff;border-radius:8px;padding:12px;width:33%">
      <b>💬 评论 · 联系人</b><br>
      <span style="color:#555">发布/删除评论（归属校验）<br>联系人增删查</span>
    </td>
    <td align="center" style="background:#ffeef0;border-radius:8px;padding:12px;width:33%">
      <b>🛡 三权分立</b><br>
      <span style="color:#555">普通用户 · 失物招领管理员<br>系统管理员（角色中间件）</span>
    </td>
    <td align="center" style="background:#eef6f6;border-radius:8px;padding:12px;width:33%">
      <b>📢 公告 · 数据</b><br>
      <span style="color:#555">公告发布/删除<br>系统数据统计</span>
    </td>
  </tr>
</table>

> 注：上表 emoji 仅作视觉标识，接口文档见下方 API 一览。

---

## 快速开始

1. **环境要求**：Go 1.20+，MySQL 5.7+
2. **创建数据库**（表结构启动时自动创建）：

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

5. 服务默认监听 **8080** 端口，静态图片目录 `/images`

---

## 目录结构

```
lost-and-found-backend/
├── app/
│   ├── controllers/   # 控制器：接收 HTTP 请求，调用 service，返回响应
│   │   ├── user_controller/          # 用户相关
│   │   ├── post_controller/          # 帖子、图片上传
│   │   ├── claim_controller/         # 认领申请
│   │   ├── comment_controller/       # 评论
│   │   ├── contact_controller/       # 联系人
│   │   ├── announcement_controller/  # 公告
│   │   ├── admin_controller/         # 失物招领管理员
│   │   └── sysadmin_controller/      # 系统管理员
│   ├── middlewares/   # 中间件：JWT 校验、角色权限、统一错误响应、登录限流
│   ├── models/        # 数据库模型（对应数据表）
│   ├── services/      # 业务逻辑层
│   ├── errs/          # 统一业务错误码定义
│   └── utils/         # 工具：JWT、密码加密、时间等
├── configs/
│   ├── config/        # 配置加载（viper）
│   ├── database/      # 数据库初始化与自动建表
│   └── router/        # 路由注册
├── images/            # 上传图片存储目录（静态开放）
├── config.example.yaml # 配置模板（真实配置为 config.yaml，不入库）
├── main.go            # 唯一入口
└── README.md
```

---

## 角色与权限

<table>
  <tr>
    <td align="center" style="background:#eaf6ff;border-radius:8px;padding:12px;width:33%">
      <b>👤 普通用户</b><br>
      <span style="color:#555">发布帖子 · 认领 · 评论 · 联系人</span><br>
      <span style="color:#888">直接注册</span>
    </td>
    <td align="center" style="background:#fff4e5;border-radius:8px;padding:12px;width:33%">
      <b>🛠 失物招领管理员</b><br>
      <span style="color:#555">审核帖子 · 删除违规帖子</span><br>
      <span style="color:#888">注册时携带管理员邀请码</span>
    </td>
    <td align="center" style="background:#ffeef0;border-radius:8px;padding:12px;width:33%">
      <b>⚙️ 系统管理员</b><br>
      <span style="color:#555">用户角色 · 公告 · 系统数据</span><br>
      <span style="color:#888">注册时携带管理员邀请码</span>
    </td>
  </tr>
</table>

> 注册时 `role` 字段与 `invite_code` 决定身份：`普通用户` 无需邀请码，`失物招领管理员` / `系统管理员` 必须提供邀请码。

---

## API 一览

统一前缀 `/api`；除「公开接口」外均需请求头 `Authorization: Bearer <token>`。

### 📖 公开接口

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| POST | `/api/register` | 注册 |
| POST | `/api/login` | 登录 |
| GET | `/api/posts` | 查询所有帖子（公开列表） |
| GET | `/api/posts/:post_id/comments` | 获取评论 |
| GET | `/api/announcements` | 查看公告 |

### 👤 用户（需登录）

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| GET | `/api/user/profile` | 获取用户个人信息 |
| PATCH | `/api/user/profile` | 修改用户个人信息 |
| PATCH | `/api/user/password` | 修改密码 |

### 📦 失物招领帖子（需登录）

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| POST | `/api/upload` | 上传图片 |
| POST | `/api/posts` | 发布帖子（默认待审核） |
| GET | `/api/my/posts` | 查询自己发布的帖子 |
| DELETE | `/api/my/posts/:post_id` | 删除自己的帖子 |
| GET | `/api/posts/:post_id` | 查看帖子详情 |
| PUT | `/api/my/posts/:post_id` | 编辑并重新提交帖子（进入待审核） |

### 🔑 认领申请（需登录）

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| POST | `/api/posts/:post_id/claims` | 提交认领申请 |
| DELETE | `/api/claims/:claim_id` | 撤回认领申请（仅待处理可撤回） |
| GET | `/api/my/claims` | 查看我提交的认领申请 |
| GET | `/api/posts/:post_id/claims` | 查看我的帖子收到的认领申请 |
| PUT | `/api/claims/:claim_id/audit` | 处理认领申请（同意/拒绝） |

### 📇 联系人（需登录）

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| POST | `/api/contact` | 添加联系人 |
| GET | `/api/my/contacts` | 查询联系人列表 |
| DELETE | `/api/my/contacts/:contact_id` | 删除联系人 |

### 💬 评论（需登录）

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| POST | `/api/posts/:post_id/comments` | 发布评论 |
| DELETE | `/api/posts/:post_id/comments/:comment_id` | 删除自己的评论 |

### 🛠 失物招领管理员

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| GET | `/api/admin/posts` | 查询所有帖子（含未审核） |
| PATCH | `/api/admin/posts/:post_id/audit` | 审核帖子（通过/拒绝） |
| DELETE | `/api/admin/posts/:post_id` | 删除帖子 |

### ⚙️ 系统管理员

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| GET | `/api/sys/users` | 查询所有用户 |
| PATCH | `/api/sys/users/:user_id/role` | 修改用户角色 |
| POST | `/api/sys/announcements` | 发布公告 |
| DELETE | `/api/sys/announcements/:announcement_id` | 删除公告 |
| GET | `/api/sys/stats` | 获取系统数据 |

---

## 统一响应格式

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

- `code = 0` 表示成功；非 0 为业务错误码（见下表）
- 错误时 HTTP 状态码与业务错误码同时返回

---

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

---

## 配置说明（config.yaml）

| 配置段 | 字段 | 说明 |
| --- | --- | --- |
| `database` | `name` / `host` / `port` / `user` / `pass` | MySQL 连接信息 |
| `jwt` | `key` | JWT 签名密钥（**上线前务必修改**） |
| `jwt` | `issuer` | 令牌签发者 |
| `register` | `admin_invite_code` | 管理员注册邀请码（**上线前务必修改**） |
| `upload` | `base_url` | 图片访问基础地址（如 `http://localhost:8080`，返回给前端的图片完整 URL 前缀） |

---

## 核心业务规则

- **帖子状态流**：`待审核` → 管理员审核 → `已通过` / `已拒绝`；已通过的帖子才能被查看详情、被认领、被评论
- **认领规则**：仅 `已通过` 且 `未解决` 的帖子可提交认领；同一用户对同一帖子只能有一条 `待处理` 申请；不能认领自己发布的帖子；并发提交已通过数据库事务 + 行锁保证不产生重复申请
- **认领状态流**：`待处理` → 发布者同意 → 帖子标记为 `已解决`；被拒绝后可重新提交
- **删除评论**：路径中的 `post_id` 必须与评论实际所属帖子一致，否则返回 `400/2005`
- **登录限流**：连续登录失败达到阈值后账号锁定 15 分钟（`429/1016`）

## 注意事项

- `config.yaml` 已加入 `.gitignore`，请勿提交真实配置
- 演示环境 `jwt.key`、数据库密码、管理员邀请码使用同一调试值，**上线前必须分别改为独立强口令**，否则邀请码泄露等于令牌与数据库同时失守
- 上传图片保存在本地 `./images` 目录，正式部署建议改用对象存储
