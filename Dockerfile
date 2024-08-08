# 使用 golang:1.20-alpine 作为基础镜像
FROM golang:1.22-alpine

# 创建一个非 root 用户和用户组
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# 设置工作目录并确保非 root 用户有权限
WORKDIR /app
RUN chown -R appuser:appgroup /app

# 切换到非 root 用户
USER appuser

# 启动时进入 bash，方便进行开发工作
CMD ["sh"]
