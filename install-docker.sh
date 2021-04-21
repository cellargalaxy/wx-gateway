#!/usr/bin/env bash

while :
do
    if [ ! -z $configPath ];then
        break
    fi
    read -p "please enter config path(required):" configPath
done

if [ -z $listenPort ];then
    read -p "please enter listen port(default:8990):" listenPort
fi
if [ -z $listenPort ];then
    listenPort="8990"
fi

echo 'configPath:'$configPath
echo 'listenPort:'$listenPort
echo 'input any key go on, or control+c over'
read

echo 'stop container'
docker stop wx_gateway
echo 'remove container'
docker rm wx_gateway
echo 'remove image'
docker rmi wx_gateway
echo 'docker build'
docker build -t wx_gateway .
echo 'docker run'
docker run -d \
--restart=always \
--name wx_gateway \
-p $listenPort:8990 \
-v $configPath:/resources/config.ini \
wx_gateway

echo 'all finish'
