# Sammy AI - 云养娃陪伴产品

🎉 **Sammy AI** 是一个基于 PocketBase 的云养娃陪伴产品，让你可以随时随地与可爱的 AI 宝贝 Sammy 互动。

## ✨ 新特性

### 🎨 全新视觉设计
- **粉色渐变主题**：温馨可爱的视觉风格
- **浮动泡泡动画**：活泼生动的背景效果
- **圆润卡片设计**：现代简约的用户界面
- **响应式布局**：完美适配手机和电脑

### 🎮 丰富的互动功能
- **聊天**：与 Sammy 进行有趣的对话
- **喂食**：给 Sammy 喂食，看它开心地吃饭
- **玩耍**：和 Sammy 一起玩游戏
- **哄睡**：温柔地哄 Sammy 入睡

### 🗣️ 语音交互
- **语音合成**：Sammy 会用可爱的声音和你说话
- **语音输入**：支持语音输入，解放双手
- **实时反馈**：Sammy 会根据互动给出语音反应

### 📊 成长系统
- **成长进度**：实时查看 Sammy 的成长状态
- **互动统计**：记录每次互动的次数和类型
- **心情变化**：Sammy 的心情会随着互动而变化

## 🚀 快速开始

### 本地开发

```bash
# 克隆项目
git clone https://github.com/cuiyouyoubaby-source/-sammy-ai.git
cd -sammy-ai

# 启动服务
./start.sh
```

### 访问地址

- 🌐 **前端页面**: http://localhost:8090
- 🔧 **管理后台**: http://localhost:8090/_
- 📚 **API 文档**: http://localhost:8090/api

## 🛠️ 技术栈

- **后端**: Go + PocketBase
- **前端**: HTML5 + CSS3 + JavaScript
- **语音**: Web Speech API
- **动画**: CSS3 Animations

## 📱 功能特性

### 已实现
- ✅ 可爱的 Sammy 形象展示
- ✅ 丰富的互动功能（聊天、喂食、玩耍、哄睡）
- ✅ 语音合成（Sammy 会说话）
- ✅ 语音输入支持
- ✅ 成长进度追踪
- ✅ 互动历史记录
- ✅ 响应式设计

### 开发中
- 🔄 更多 Sammy 表情和动作
- 🔄 个性化定制
- 🔄 多用户支持
- 🔄 云端数据同步

## 🎨 设计特色

### 视觉风格
- **粉色渐变**：温馨可爱的主色调
- **圆润设计**：友好的视觉体验
- **动画效果**：生动的交互反馈
- **浮动元素**：活泼的页面氛围

### 交互设计
- **即时反馈**：每次互动都有视觉和语音反馈
- **动画过渡**：流畅的页面切换效果
- **语音交互**：自然的对话体验
- **成长系统**：可视化的进度展示

## 📝 API 文档

### 互动 API

```http
POST /api/sammy/interact
Content-Type: application/json

{
  "type": "chat",
  "content": "你好呀"
}
```

### 聊天 API

```http
POST /api/sammy/chat
Content-Type: application/json

{
  "message": "今天天气真好"
}
```

### 语音合成 API

```http
POST /api/sammy/tts
Content-Type: application/json

{
  "text": "你好呀！我是 Sammy"
}
```

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

MIT License
