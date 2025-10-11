FROM alpine:latest

# 先更新包列表，并安装 shadow（提供 groupadd, useradd）
RUN apk add --no-cache shadow

# FROM alpine:V0.0.00.006
### Containers should NOT run as root as a good practice
RUN groupadd -g 2030 nexagroup && useradd -u 2030 -g nexagroup nexa-user -d /home/nexa-user -m

ENV APP_ROOT=/home/nexa-user
ENV PATH=${APP_ROOT}/bin:${PATH} HOME=${APP_ROOT}

RUN mkdir ${APP_ROOT}/bin

# 创建必要的目录
RUN mkdir  ${APP_ROOT}/certs
RUN mkdir  ${APP_ROOT}/conf

# 复制程序和证书文件
# COPY server.key server.crt ${APP_ROOT}/certs/

RUN chmod -R +rx ${APP_ROOT} && \
    chown -R nexa-user:nexagroup ${APP_ROOT}

COPY --chmod=757  ../../cmd/nexa/nexa ${APP_ROOT}/bin/

RUN chown -R  nexa-user:nexagroup ${APP_ROOT}/conf

USER 2030
WORKDIR ${APP_ROOT}

ENTRYPOINT ["/home/nexa-user/bin/nexa httpserver"]