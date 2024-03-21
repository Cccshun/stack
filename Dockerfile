# 从go镜像上开始构建
FROM golang:1.19 AS BUILDER

# 设置go env
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 创建工作目录
WORKDIR /build

# 将上下文目录拷贝至容器工作目录中
COPY . .

# 编译成二进制文件app
RUN go build -o ../dist/app .

# 创建用于存放二进制文件的目录
#WORKDIR /dist

# 将二进制文件从build移动到dist
#RUN cp /build/app .

FROM scratch

COPY --from=BUILDER /dist/app /

# 声明服务端口
EXPOSE 8080

# 容器运行时启动go服务
CMD /app
