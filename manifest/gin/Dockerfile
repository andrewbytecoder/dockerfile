# 构建阶段
FROM golang:1.25.1-alpine AS builder

# 设置工作目录
WORKDIR /build

# 复制源代码
COPY .git .git
COPY . .

# 获取 Git 信息
RUN GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown") && \
    GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD 2>/dev/null || echo "unknown") && \
    GIT_TAG=$(git describe --tags --exact-match 2>/dev/null || echo "none") && \
    BUILD_TIME=$(date -u '+%Y-%m-%dT%H:%M:%SZ') && \
    GO_VERSION=$(go version | awk '{print $3}') && \
    echo "Git Commit: $GIT_COMMIT" && \
    echo "Git Branch: $GIT_BRANCH" && \
    echo "Git Tag: $GIT_TAG" && \
    echo "Build Time: $BUILD_TIME" && \
    echo "Go Version: $GO_VERSION" && \
    CGO_ENABLED=0 GOOS=linux go build \
      -ldflags="-s -w \
        -X main.GitCommit=${GIT_COMMIT} \
        -X main.GitBranch=${GIT_BRANCH} \
        -X main.GitTag=${GIT_TAG} \
        -X main.BuildTime=${BUILD_TIME} \
        -X main.GoVersion=${GO_VERSION}" \
      -a -installsuffix cgo -o main ./cmd/gin/main.go

# 运行阶段，最小化镜像
FROM alpine:V0.0.2

# 可选：设置非 root 用户（提升安全性）
RUN adduser -D -s /bin/sh nmq-user

RUN mkdir -p /home/nmq-user

# 设置工作目录
WORKDIR /home/nmq-user

# 从构建阶段复制编译好的二进制文件
COPY --chmod=757 --from=builder /build/main /home/nmq-user/

# 暴露端口
EXPOSE 8080

# 更改文件所有者
RUN chown nmq-user:nmq-user main

# 切换到非 root 用户
USER nmq-user

# 设置时区
ENV TZ=Asia/Shanghai

# 启动应用
CMD ["/home/nmq-user/main"]


