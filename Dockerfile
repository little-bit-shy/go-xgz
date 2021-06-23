FROM alpine:3.14.0

ARG ENV

COPY ./Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone

RUN mkdir /lib64 \
    && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

RUN mkdir /data
WORKDIR /data

ADD ./cmd/server /data
ADD ./configs/${ENV} /data/configs
ADD ./language /data/language

RUN chmod 777 /data/server

EXPOSE 8000
EXPOSE 9000

ENTRYPOINT ["./server", "-conf", "./configs", "-language", "./language"]