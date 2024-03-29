#!/bin/bash

DEMO_PATH=`pwd`
GOLANG_DOCKER_VERSION=golang:alpine
DOCKER_CONTAINER=golang-runtime
APP_NAME=dev-tools-go
DOCKER_RUN_PATH=/go/src/$APP_NAME

usage()
{
  echo "USAGE: $CALLER [-h] COMMANDS"
  echo "       $CALLER [ --help ]"
  echo "COMMANDS:"
  echo "    install       创建docker容器"
  echo "    start         启动docker容器"
  echo "    test          运行单元测试"
  echo "    enter         进入docker容器"
  echo "Run '$CALLER COMMAND --help' for more information on a command."
  exit 1
}

installDocker () {
  echo "start install docker container"

  docker run -itd \
    -v $DEMO_PATH:$DOCKER_RUN_PATH \
    --name $DOCKER_CONTAINER \
    $GOLANG_DOCKER_VERSION &> /dev/null

  # docker run -itd \
  #   -v $DEMO_PATH:/go/src/dev-tools-go \
  #   --name golang-runtime \
  #   golang:1.10 &> /dev/null
}

startDocker () {
  docker start $DOCKER_CONTAINER
}

enterDocker () {
  docker exec -it $DOCKER_CONTAINER /bin/sh
}

execTestKata () {
  docker exec -it $DOCKER_CONTAINER sh $DOCKER_RUN_PATH/scripts/init.sh
  # docker exec -it $DOCKER_CONTAINER sh -c 'cp /etc/apk/repositories /etc/apk/repositories.bak;echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories;apk add --no-cache make'
  docker exec -it $DOCKER_CONTAINER sh -c "cd $DOCKER_RUN_PATH;make coverhtml"
  # docker exec -it $DOCKER_CONTAINER sh -c 'apk del git'
}

checkDockerContainerStatus () {
  check_result=`docker ps -a | grep $DOCKER_CONTAINER`

  if [ ! "$check_result" ]; then
    return 0
  fi

  check_exit=`echo $check_result | grep Exited`
  if [ ! "$check_exit" ]; then
    return 2
  fi
  return 1
}

checkDockerInstalled () {
  check_result=`which docker`

  if [ ! "$check_result" ]; then
    echo "没有检测到docker，请先安装docker之后再试"
    exit 1
  fi
}

if [ $# -ne 1 ] ; then
    usage
fi

checkDockerInstalled

checkDockerContainerStatus
DOCKER_START_STATUS=$?

if [ "$1" = "install" ]; then
  if [ "$DOCKER_START_STATUS" != "0" ]; then
    echo "容器已经存在"
    exit 0
  fi
  installDocker
  exit 0
fi


if [ "$1" = "start" ]; then
  if [ "$DOCKER_START_STATUS" = "2" ]; then
    echo "container is working"
    exit 0
  elif [ "$DOCKER_START_STATUS" = "0" ]; then
    echo "container not exist"
    installDocker
  fi
  startDocker
elif [ "$1" = "enter" ]; then
  if [ "$DOCKER_START_STATUS" = "0" ]; then
    echo "container not exist"
    installDocker
  elif [ "$DOCKER_START_STATUS" != "2" ]; then
    startDocker
  fi

  enterDocker
elif [ "$1" = "test" ]; then
  if [ "$DOCKER_START_STATUS" = "0" ]; then
    echo "container not exist"
    installDocker
  elif [ "$DOCKER_START_STATUS" != "2" ]; then
    startDocker
  fi

  execTestKata
fi
