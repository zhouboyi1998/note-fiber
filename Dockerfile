FROM alpine
LABEL maintainer="zhouboyi<1144188685@qq.com>"

WORKDIR /go/note-fiber
COPY ./main /go/note-fiber
COPY ./application-docker.yaml /go/note-fiber

# 设置环境变量
ENV ENVCONFIG docker

EXPOSE 18099
ENTRYPOINT ["./main"]
