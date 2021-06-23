FROM ubuntu:20.10

#RUN rm -rf /etc/apt/sources.list
#RUN echo "deb http://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse" >> /etc/apt/sources.list
#RUN echo "deb-src http://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse" >> /etc/apt/sources.list
#RUN echo "deb http://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse" >> /etc/apt/sources.list
#RUN echo "deb-src http://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse" >> /etc/apt/sources.list
#RUN echo "deb http://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse" >> /etc/apt/sources.list
#RUN echo "deb-src http://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse" >> /etc/apt/sources.list
#RUN echo "deb http://mirrors.aliyun.com/ubuntu/ focal-proposed main restricted universe multiverse" >> /etc/apt/sources.list
#RUN echo "deb-src http://mirrors.aliyun.com/ubuntu/ focal-proposed main restricted universe multiverse" >> /etc/apt/sources.list
#RUN echo "deb http://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse" >> /etc/apt/sources.list
#RUN echo "deb-src http://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse" >> /etc/apt/sources.list

RUN apt update
RUN apt install -y wget git

RUN cd /tmp \
    && mkdir /usr/local/go \
    && wget https://golang.google.cn/dl/go1.16.linux-amd64.tar.gz \
    && tar -xzf /tmp/go1.16.linux-amd64.tar.gz -C /usr/local/go --strip-components=1 \
    && rm -f /tmp/go1.16.linux-amd64.tar.gz \
    && ln -s /usr/local/go/bin/go /usr/local/bin/go

RUN go env -w GO111MODULE="on" \
    && go env -w GOPROXY="https://goproxy.io,direct" \
    && go env -w GOOS="linux"

RUN go get github.com/go-kratos/kratos/tool/kratos@v1.0.0 \
    && ln -s /root/go/bin/kratos /usr/local/bin/kratos

RUN go get github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
    && ln -s /root/go/bin/protoc-gen-doc /usr/local/bin/protoc-gen-doc

ENV  GO_PATH=/root/go
ENV  PROTO_PATH=${GO_PATH}/src
ENV  THIRD_PROTO_PATH=${GO_PATH}/pkg/mod/github.com/go-kratos/kratos@v1.0.0/third_party
ENV  PROTOC_GEN_BM=${GO_PATH}/bin/protoc-gen-bm
ENV  PROTOC_GEN_GO=${GO_PATH}/bin/protoc-gen-go
ENV  PROTOC_GEN_GOFAST=${GO_PATH}/bin/protoc-gen-gofast

COPY ./docker/go/protoc /usr/local/protoc
COPY ./docker/go/protoc-gen-bm ${PROTOC_GEN_BM}
COPY ./docker/go/protoc-gen-go ${PROTOC_GEN_GO}
COPY ./docker/go/protoc-gen-gofast ${PROTOC_GEN_GOFAST}

RUN  ln -s ${PROTOC_GEN_BM} /usr/local/bin/protoc-gen-bm
RUN  ln -s ${PROTOC_GEN_GO} /usr/local/bin/protoc-gen-go
RUN  ln -s ${PROTOC_GEN_GOFAST} /usr/local/bin/protoc-gen-gofast
RUN  ln -s /usr/local/protoc/bin/protoc /usr/local/bin/protoc

#私有仓库相关
#RUN git config --global url."ssh://git@gitlab.******.com:62022/".insteadof "https://gitlab.******.com/"
#RUN go env -w GOPRIVATE="gitlab.******.com/*"
#
#COPY ./id_rsa /root/.ssh/id_rsa
#RUN chmod 600 /root/.ssh/id_rsa
#COPY ./ssh_config /etc/ssh/ssh_config

CMD ["/bin/bash"]