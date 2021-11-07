FROM guoxf/golang-build:1.17.0-alpine3.14 as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

WORKDIR /workspace
COPY . .
RUN go env && go build -gcflags=-G=3 -o usercenter main.go
RUN pwd

FROM guoxf/golang-run:alpine-3.13.5
LABEL MAINTAINER="ants.guoxf@gmail.com"

ENV DUBBO_GO_CONFIG_PATH="./dubbogo.yaml"

WORKDIR /workspace
COPY --from=builder /workspace/usercenter ./

EXPOSE 20000

ENTRYPOINT ./usercenter
