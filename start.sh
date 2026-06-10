#!/bin/bash


# Sammy AI 启动脚本

echo "🚀 启动 Sammy AI 云养娃陪伴产品..."

# 检查 Go 是否安装
if ! command -v go &> /dev/null; then
    echo "❌ 错误：Go 未安装，请先安装 Go 1.22+"
    exit 1
fi

# 检查前端目录是否存在
if [ ! -d "frontend" ]; then
    echo "❌ 错误：frontend 目录不存在"
    exit 1
fi

# 下载依赖
echo "📦 下载依赖..."
go mod tidy

# 构建项目
echo "🔨 构建项目..."
go build -o sammy-ai

# 检查构建是否成功
if [ ! -f "sammy-ai" ]; then
    echo "❌ 错误：构建失败"
    exit 1
fi

# 启动服务
echo "✅ 启动成功！"
echo ""
echo "🌐 访问地址："
echo "   - 前端页面：http://localhost:8090"
echo "   - 管理后台：http://localhost:8090/_"
echo "   - API 文档：http://localhost:8090/api"
echo ""
echo "👤 默认管理员账号："
echo "   - 邮箱：admin@example.com"
echo "   - 密码：admin123"
echo ""

# 启动 PocketBase
./sammy-ai serve
