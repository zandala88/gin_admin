FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
ENV GOOS linux
ENV GOARCH amd64

WORKDIR /app
COPY . /app

RUN go build .
EXPOSE 9090
ENTRYPOINT ["./gin_admin"]