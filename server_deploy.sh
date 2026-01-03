#!/bin/bash

# 0. 拉取最新代码 (Pull latest code)
echo "Pulling latest code from Gitee..."
git pull origin master || { echo "Git pull failed! Continuing anyway..."; }

# 1. 停止旧容器 (Stop old containers)
echo "Stopping old containers..."
docker stop smart-backend smart-frontend || true
docker rm smart-backend smart-frontend || true

# 2. 重新构建镜像 (Rebuild Images - as requested)
echo "Building Backend Image..."
docker build -t smart-backend:v1 .

echo "Building Frontend Image..."
cd frontend
docker build -t smart-frontend:v1 .
cd ..

# 3. 启动容器 (Start Containers manually)
echo "Starting Backend..."
# Map Host 8082 -> Container 8081. Mount config.
docker run -d \
  --name smart-backend \
  -p 8082:8081 \
  -e APP_ENV=prod \
  -v $(pwd)/config/prod.yaml:/app/config/prod.yaml \
  smart-backend:v1

echo "Starting Frontend..."
# Map Host 81 -> Container 81. Mount nginx config.
docker run -d \
  --name smart-frontend \
  -p 81:81 \
  -v $(pwd)/nginx.prod.conf:/etc/nginx/conf.d/default.conf \
  smart-frontend:v1

echo "---------------------------------------"
echo "Deployment Complete!"
echo "Frontend: http://43.138.85.114:81"
echo "Backend:  http://43.138.85.114:8082"
echo "---------------------------------------"
