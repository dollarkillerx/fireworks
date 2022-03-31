FROM alpine:latest
WORKDIR /app
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update && \
    apk add tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone
COPY ./cmd/backend/fireworks /app
COPY ./configs/ /app/configs
COPY ./dist/ /app/dist
CMD ["/app/fireworks"]
