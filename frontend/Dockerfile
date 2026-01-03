# 构建阶段：安装依赖并打包前端
FROM node:18-alpine AS builder
WORKDIR /app
# 复制前端依赖清单
COPY package.json package-lock.json ./
# 国内镜像加速安装依赖
RUN npm install --registry=https://registry.npmmirror.com
# 复制所有前端代码
COPY . .
# 打包生产环境静态文件（生成dist目录）
RUN npm run build

# 运行阶段：Nginx托管静态文件
FROM nginx:alpine
# 复制打包产物到Nginx默认目录
COPY --from=builder /app/dist /usr/share/nginx/html
# 复制自定义Nginx配置（覆盖默认配置）
# 复制自定义Nginx配置（覆盖默认配置）
# COPY ../nginx.conf /etc/nginx/conf.d/default.conf
# 注意：我们将在部署脚本中通过 -v 挂载 nginx.prod.conf，所以这里不需要复制
# 暴露前端访问端口
EXPOSE 81
# 启动Nginx（前台运行）
CMD ["nginx", "-g", "daemon off;"]