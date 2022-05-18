## 注意事项

* 热更新需要先 git config --global user.email "you@example.com"

## 更新日志
### 4-25
* 适配 Nark 

## 安装教程 

xdd-plus安装教程

第一步：下载go

`cd /usr/local && wget https://dl.google.com/go/go1.17.4.linux-amd64.tar.gz -O go1.17.4.linux-amd64.tar.gz`

第二步：解压go

`tar -xvzf go1.17.4.linux-amd64.tar.gz`

第三步：设置环境变量 

`vi /etc/profile`

将文本复制到最后一行

```
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
export GOROOT=/usr/local/go
export GOPATH=/usr/local/go/path
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

第五步：先按 esc 然后输入 :wq 

保存文件后 

`source /etc/profile`

第六步：检查go安装

`go env`

第七步：拉xdd-plus的库

`cd ~ && git clone https://ghproxy.com/https://github.com/xp9477/xdd-plus.git`

第八步：编译xdd-plus

`cd /root/xdd-plus && go build && chmod 777 xdd-plus`

