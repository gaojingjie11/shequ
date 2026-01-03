#!/bin/bash

# 0. 拉取最新代码 (Pull latest code)
echo "Pulling latest code from Gitee..."
git pull origin master || { echo "Git pull failed! Continuing anyway..."; }

# 1. 停止旧容器 (Stop old containers)
echo "Stopping old containers..."
docker-compose -f docker-compose.prod.yml down || true

# 2. 重新构建镜像 (Rebuild Images)
echo "Building Backend Image..."
docker build -t smart-backend:v1 .

echo "Building Frontend Image..."
cd frontend
docker build -t smart-frontend:v1 .
cd ..

# 3. 启动 Docker Compose (Start Services)
echo "Starting services..."
docker-compose -f docker-compose.prod.yml up -d

echo "---------------------------------------"
echo "Deployment Complete!"
echo "Frontend: http://43.138.85.114:81"
echo "Backend:  http://43.138.85.114:8082"
echo "---------------------------------------"
