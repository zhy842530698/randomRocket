# randomRocket

基于Socket5的轻量级代理服务工具

## 快速开始

进入根目录下

需要本地或远程安装golang编译环境

服务端任务编译服务端任务

```shell
>cd randomRocket/cmd && go build -o server server_main.go
> ./server
```

此时会产生一个

![image-20240719105421155](/Users/zhanghaoyun/Library/Application Support/typora-user-images/image-20240719105421155.png)

将服务器ip、port、密码copy一份到本地

```shell
>cd randomRocket/cmd && go build  -o client client_main.go 
> ./client -p <服务端生成的密码> -h <服务器ip+":"+port>
```

随后开启客户端监听

## 性能测试

相比与传统的ShadowSocksR,其特点在加密端做了简化处理(抛弃整体的aes、des加密过程),这样使得整个代理网络传输的更快,整个代码用golang实现

## 讨论

声明:只提供技术服务,非商业性质用途,非灰色产业链,出现任何法律与博主无关

私人邮件:842530698@qq.com