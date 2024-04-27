#!/bin/bash

# 文件所在目录
HOME="$WORK/com.lh.service"

# 拉取最新代码
git pull

# 打包go脚本
rm -rf ${HOME}/output/*
go build -o ${HOME}/output/lhService.sh ${HOME}/main.go
${HOME}/output/lhService.sh &