#! /bin/bash

rm -f ./config.yaml
# 生成生产环境config.yaml文件
cp .docker-compose/shell/config.yaml ./config.yaml

pip3 install -r python/requirements.txt
go build -o ../../bin/server .
../../bin/server
