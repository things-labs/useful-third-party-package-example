# 编译环境
FROM golang:1.16.3-alpine3.12 AS builder

ENV GOPROXY https://goproxy.io,direct

WORKDIR /release

COPY . .

# building
RUN CGO_ENABLED=0 go build -o helloworld .


# 运行环境
FROM frolvlad/alpine-glibc:alpine-3.12_glibc-2.31

MAINTAINER thinkgos "gb.jiang@fjwanrun.com"

WORKDIR /root

COPY --from=builder /release/helloworld .

EXPOSE 80

ENTRYPOINT ["./helloworld"]