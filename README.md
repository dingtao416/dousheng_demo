# dousheng_demo

## 极简抖音

**注：目前仅考虑在 Ubuntu 环境下运行**

项目依赖 go 1.17，docker 与 docker-compose 运行

[docker 安装](https://yeasy.gitbook.io/docker_practice/install/ubuntu)

docker-compose 安装

```shell
sudo apt-get install docker-compose 
```

自动化一键启动需要的容器 （现在仅需 mysql）

```shell
docker-compose up -d
```

（如果不是在 WSL 环境下运行，还需修改 config.ini）

项目编译与运行

```shell
go build && ./dousheng_demo
```