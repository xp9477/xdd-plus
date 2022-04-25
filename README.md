## 注意事项

  1. master:的值即为密码，后面不可带注释，全匹配方可登录，也不要pt_pin 可自定义
  2. 2.9+版本需要配置    cid和secret 在青龙里面系统设置，添加应用后配置

## 更新日志
### 4-25
* 适配 Nark

## 安装教程 

xdd-plus安装教程

第一步：下载go

cd /usr/local && wget https://dl.google.com/go/go1.17.4.linux-amd64.tar.gz -O go1.17.4.linux-amd64.tar.gz

第二步：解压go

tar -xvzf go1.17.4.linux-amd64.tar.gz

第三步：设置环境变量 

vi /etc/profile

将文本复制到最后一行

export GO111MODULE=on

export GOPROXY=https://goproxy.cn

export GOROOT=/usr/local/go

export GOPATH=/usr/local/go/path

export PATH=$PATH:$GOROOT/bin:$GOPATH/bin


第五步：先按 esc 

然后输入 :wq 

保存文件后 

source /etc/profile

第六步：检查go安装

go env

第七步：拉xdd-plus的库

cd ~ && git clone https://ghproxy.com/https://github.com/xp9477/xdd-plus.git

第八步：编译xdd-plus

cd /root/xdd-plus && go build

# 常见问题

编码问题参考
https://blog.csdn.net/qq_29499107/article/details/83583983/usr/lib64/python3.6/http

Token故障请先用官方教程重装  已排查是nginx问题
https://thin-hill-428.notion.site/2-8Faker-QL-pannel-Faker-Repository-environment-Setup-45edcbfe90d74d8abb2d71896eab3be7
请使用官方一键安装 就解决此问题了



1. 如何自动更新短信镜像

```
docker run -d \
    --name watchtower \
    --restart always \
    -v /var/run/docker.sock:/var/run/docker.sock \
    containrrr/watchtower \
    --cleanup
```

