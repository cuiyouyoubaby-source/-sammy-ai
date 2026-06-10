# Sammy AI - 云养娃陪伴产品

基于 PocketBase 的 AI 云养娃陪伴产品，为用户提供沉浸式虚拟宝宝陪伴体验。

## 项目概述

Sammy AI 是一款全平台沉浸式 AI 云养娃陪伴产品，聚焦当代年轻人独居孤独、晚婚不育、情绪内耗，以及老年人空巢寂寞、缺乏情感寄托的核心痛点。

## 核心功能

### 1. 个性化人设自定义系统
- 自定义性别、初始年龄、外貌风格
- 自主选择性格特质：活泼开朗、温柔乖巧、调皮可爱、安静内敛、治愈软萌等
- 可微调五官、穿搭、造型样式

### 2. 现实同步真实成长引擎
- 以真实日历、天数为成长基准
- 随时间推移逐步增长年龄、变化心智
- 行为、语言、情绪、认知能力随年龄同步迭代

### 3. 全动态拟人化交互体验
- 实时语音对话
- 动态表情系统
- 肢体动态效果
- 主动陪伴机制

### 4. 双端全场景适配
- 移动端：轻便随时玩
- PC端：高清大屏动态展示
- 数据云端互通

## 技术栈

- **后端**: PocketBase (Go)
- **前端**: HTML5 + JavaScript + ECharts
- **数据库**: SQLite
- **实时通信**: WebSocket

## 快速开始

### 安装依赖

```bash
# 安装 Go 依赖
go mod tidy

# 构建项目
go build -o sammy-ai
```

### 运行项目

```bash
# 启动 PocketBase 服务
./sammy-ai serve

# 访问管理后台
# http://localhost:8090/_/ (admin@example.com / admin123)

# 访问前端页面
# http://localhost:8090
```

## 项目结构

```
sammy-ai/
├── main.go              # 主入口文件
├── go.mod               # Go 模块配置
├── go.sum               # Go 依赖锁定
├── pb_data/             # PocketBase 数据目录
├── pb_migrations/       # 数据库迁移文件
├── pb_hooks/            # 自定义钩子
├── frontend/            # 前端页面
│   ├── index.html       # 主页面
│   ├── css/
│   └── js/
└── README.md
```

## API 接口

### Sammy 成长状态
- **GET** `/api/sammy/growth/:userId`
- 获取 Sammy 的成长状态

### Sammy 互动
- **POST** `/api/sammy/interact`
- 与 Sammy 进行互动

## 数据库集合

### users (系统自带)
- 用户认证集合

### sammy_profiles
- userId: 关联用户
- name: Sammy 名字
- gender: 性别
- birthDate: 出生日期
- personality: 性格特质
- appearance: 外貌风格
- growthPoints: 成长点数
- mood: 当前心情
- lastInteraction: 最后互动时间

### interactions
- userId: 关联用户
- sammyId: 关联 Sammy
- type: 互动类型 (chat/voice/play/sleep)
- content: 互动内容
- moodChange: 心情变化
- timestamp: 时间戳

## 开发计划

- [x] 项目初始化
- [ ] 数据库集合创建
- [ ] 用户认证系统
- [ ] Sammy 人设系统
- [ ] 成长引擎
- [ ] 互动系统
- [ ] 前端页面
- [ ] 语音交互
- [ ] 移动端适配

## 许可证

MIT
