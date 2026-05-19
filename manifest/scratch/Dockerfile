# 构建阶段
FROM golang:1.25.1-alpine AS builder

# 设置工作目录
WORKDIR /build

# 复制源代码
COPY .git .git
COPY . .

# 设置构建参数
ARG ARCH
ARG GIT_COMMIT
ARG GIT_BRANCH
ARG GIT_TAG
ARG BUILD_TIME

# 获取 Git 信息
RUN GO_VERSION=$(go version | awk '{print $3}') && \
    echo "Git Commit: $GIT_COMMIT" && \
    echo "Git Branch: $GIT_BRANCH" && \
    echo "Git Tag: $GIT_TAG" && \
    echo "Build Time: $BUILD_TIME" && \
    echo "Go Version: $GO_VERSION" && \
    CGO_ENABLED=0 GOOS=linux go build \
      -ldflags="-s -w \
        -X main.GitCommit=${GIT_COMMIT} \
        -X main.GitBranch=${GIT_BRANCH} \
        -X main.BuildTime=${BUILD_TIME} \
        -X main.GoVersion=${GO_VERSION}" \
      -a -installsuffix cgo -o main ./cmd/gin/main.go


# 阶段 2：导出阶段，仅用于提取二进制文件
FROM scratch AS export-stage

# 从 builder 阶段复制编译好的二进制文件到当前阶段的根目录
COPY --from=builder /build/main ./main