#!/bin/bash

if [ "$1" == "update_restart" ]; then
    go build -o ../bin/server .
    pip3 install -r python/requirements.txt | grep -v "Requirement already satisfied"

    port=8888
    pid=$(netstat -nlp | grep :$port | awk '{print $7}' | awk -F"/" '{ print $1 }');
    if [ -n "$pid" ]; then
        kill -1 $pid;
    fi
    exit
fi
