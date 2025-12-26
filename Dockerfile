## 使用 Ubuntu 20.04 作为基础镜像
#FROM ubuntu:20.04
#
## 创建工作目录 /app
#RUN mkdir -p /app
#
## 将本地编译好的二进制文件 (假设叫 main) 复制到容器的 /app 目录下
## 注意：这一步要求你目录下必须先有一个编译好的 main 文件
#COPY main /app/main
#
## (可选) 如果你有配置文件目录，比如 config/，也需要复制进去
## COPY config /app/config
#
## 设置工作目录
#WORKDIR /app
#
## 赋予执行权限 (虽然通常 COPY 进去就有权限，但为了保险)
#RUN chmod +x /app/main
#
## 设置启动入口
#ENTRYPOINT ["/app/main"]