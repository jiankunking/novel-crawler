FROM golang:1.10.1-alpine3.7
MAINTAINER jiankunking(jiankunking@163.com)

# repo
RUN cp /etc/apk/repositories /etc/apk/repositories.bak
RUN echo "http://mirrors.aliyun.com/alpine/v3.6/main/" > /etc/apk/repositories
RUN echo "http://mirrors.aliyun.com/alpine/v3.6/community/" >> /etc/apk/repositories

# timezone
RUN apk update
RUN apk add --no-cache tzdata \
    && echo "Asia/Shanghai" > /etc/timezone \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# move to GOPATH
RUN mkdir -p /go/src/github.com/jiankunking/novel-crawler
COPY . $GOPATH/src/github.com/jiankunking/novel-crawler/
WORKDIR $GOPATH/src/github.com/jiankunking/novel-crawler

# build
RUN mkdir -p /app
RUN go build -o /app/novel-crawler cmd/main.go

WORKDIR /app
CMD ["/app/novel-crawler"]
