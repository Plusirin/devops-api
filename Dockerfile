# 打包依赖阶段使用golang作为基础镜像
FROM golang:1.17 as builder

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.io,direct

WORKDIR /app

COPY . .

# 指定OS等，并go build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o devops-api

RUN mkdir -p publish && cp devops-api publish && \
    cp config.yaml publish

## 运行阶段指定 scratch 作为基础镜像
FROM scratch

WORKDIR /app

# 将上一个阶段publish文件夹下的所有文件复制进来
COPY --from=builder /app/publish .

ENTRYPOINT ["./devops-api"]
