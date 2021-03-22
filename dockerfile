FROM golang:1.15

RUN apt-get update && \
    apt-get install -y \
    net-tools \
    libsm6 \
    libxrender1 \
    libxext-dev \
    libssl-dev \
    libcrypto++-dev \
    python3-pip \
    python3-dev

ENV PYTHON_VERSION 3.6.8
ADD ./python/pip.conf /root/.pip/pip.conf
ADD ./python/requirements.txt /root/.pip/requirements.txt
RUN pip3 install -r /root/.pip/requirements.txt

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

VOLUME ["/go/src/my-gin-go"]

EXPOSE 8888
WORKDIR /go/src/my-gin-go

RUN mkdir -p /go/bin
ADD .docker-compose/shell/run.sh /go/bin
RUN chmod +x /go/bin/run.sh

ADD .docker-compose/shell/control.sh /go/bin
RUN chmod +x /go/bin/control.sh
RUN ln -s /go/bin/control.sh /bin/control

CMD ["/go/bin/run.sh"]
