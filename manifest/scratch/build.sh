#!/bin/bash

echo `pwd`

# 获取 Git 信息
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null) || GIT_COMMIT="unknown"
GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD 2>/dev/null) || GIT_BRANCH="unknown"
GIT_TAG=$(git describe --tags --exact-match 2>/dev/null) || GIT_TAG="none"
BUILD_TIME=$(date -u '+%Y-%m-%dT%H:%M:%SZ')

docker build \
          -f `pwd`/manifest/scratch/Dockerfile \
          --build-arg ARCH="amd64" \
          --build-arg GIT_COMMIT="$GIT_COMMIT" \
          --build-arg GIT_BRANCH="$GIT_BRANCH" \
          --build-arg GIT_TAG="$GIT_TAG" \
          --build-arg BUILD_TIME="$BUILD_TIME" \
          --output . \
          --target export-stage \
          -t ysp-agent-build \
          .