FROM alpine:3.21

# 使用阿里云的Alpine镜像源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 更新包列表并安装Redis和Supervisor
RUN apk update && \
    apk add --no-cache redis supervisor curl && \
    mkdir -p /root/gewe/base && \
    rm -rf /var/cache/apk/* /tmp/* /usr/share/man
    
# 添加Redis的Supervisor配置文件
COPY ./*.ini /etc/supervisor.d/

#复制long,pact及配置文件 到运行目录
COPY ./bin/* /root/gewe/base/
COPY ./config/* /root/gewe/base/


# 添加Supervisor主配置文件
COPY supervisord.conf /etc/supervisord.conf

# 启动Supervisor
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]
