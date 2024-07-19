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

```shell
2024/07/11 10:03:16 服务监听地址：[::]:56302
2024/07/11 10:03:16 密码：
Jp5uubL9vRs9wHg8RiiT7Fo2XDFYIrWPtl282UMEM5nqAqiv+PL149ZPrmd5f+AnABpWBSSnl4GknzUPoZhTsVvkUZArVQaAFRzl3k7fZUq/GKVZq0UyDe/BQNHJg6y4JYXw/GiO6/vPEC9jmrCb3AwWCsXMwmbzcIzKOA5sfMugvsgdLPHtw/bao4pHP7dfls5NYXdJ54YL+SMSBx6inSEBzfq6iJVyA0j06Omzrd0Jc9Vk29I6h5JSgmrEXtMt9xdX7ip64RFv2BOqqcZEtDR7cWBLLlScCKb+UBlM1EHQ/xTibXW7xzcflCAwaZEpdnSJYot+QtfmjX07Pms5hA==
```

将服务器ip(能访问到的网卡均可)、端口port(56302)、密码复制一份到本地

```shell
>cd randomRocket/cmd && go build  -o client client_main.go 
> ./client -p <服务端生成的密码> -h <服务器ip+":"+port>
```

随后开启客户端监听

```shell
2024/07/11 10:13:27 服务监听地址：[::]:7448
2024/07/11 10:13:27 密码：
Jp5uubL9vRs9wHg8RiiT7Fo2XDFYIrWPtl282UMEM5nqAqiv+PL149ZPrmd5f+AnABpWBSSnl4GknzUPoZhTsVvkUZArVQaAFRzl3k7fZUq/GKVZq0UyDe/BQNHJg6y4JYXw/GiO6/vPEC9jmrCb3AwWCsXMwmbzcIzKOA5sfMugvsgdLPHtw/bao4pHP7dfls5NYXdJ54YL+SMSBx6inSEBzfq6iJVyA0j06Omzrd0Jc9Vk29I6h5JSgmrEXtMt9xdX7ip64RFv2BOqqcZEtDR7cWBLLlScCKb+UBlM1EHQ/xTibXW7xzcflCAwaZEpdnSJYot+QtfmjX07Pms5hA==
```

客户端默认监听7448端口,(如果被占用,可以修改一下),后面只需要将应用软件的数据从7448端口发送即可

<h2/>端口监听

**1、chrome插件从应用商店下载即可:**

https://chromewebstore.google.com/detail/proxy-switchyomega/padekgcemlokbadohgkifijomclgjgif?hl=zh-CN&utm_source=ext_sidebar

**2、新建代理服务器**

*情景模式*->新建情景模式->情景模型选择代理->协议端选择SOCKS5,输入代理服务器IP和端口

**3、更新auto switch**

在规则列表规则中选择**步骤2**新建的代理服务器名称,默认情景模式

下方规则列表格式选择AutoProxy,规则输入网址:

```json
"https://raw.githubusercontent.com/gfwlist/gfwlist/master/gfwlist.txt"
```

随后点击更新情景模式

**4、保存与退出**

点击左下方应用选项即可关闭,插件栏选择规则即可实现代理

## 性能分析

相比与传统的ShadowSocksR,其特点在加密端做了简化处理(抛弃整体的aes、des加密过程),这样使得整个代理网络传输的更快

## 讨论

声明:只提供技术交流,非商业性质用途,非灰色产业链,出现任何非法活动与盈利与作者无关

私人邮件:842530698@qq.com

haoyun zhang