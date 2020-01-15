===============================
goproxy - 简单的代理小工具
===============================

简介
===========

开发过程中，有些工具需要使用代理才能使用。本工具可以让大家快速创建一个代理服务。

注意：使用过程中你可能需要使用 qtunnel, ssh 等工具，具有一个不能本地的 vps。

功能
============

- http(s) 代理
- socks v5 代理
- socks v5 转 http(s) 代理


快速使用
=========

使用 Docker 运行服务：

```
docker run --rm -it -d -p 9000:8000 ooclab/goproxy:v1.0.0
```

测试：

```
curl -s --socks5 127.0.0.1:9000 http://httpbin.org/ip
```

编译
===========

进入 goproxy 主目录::

  make // or make install


帮助
============

查看命令帮助手册::

  goproxy --help
  goproxy http --help
  goproxy socks --help


使用
===========

HTTP
-----------

启动纯 http(s) 代理::

  goproxy http --listen 127.0.0.1:9000

启动 http(s) 代理，但使用一个 socks v5 服务作为后端::

  goproxy http --listen 127.0.0.1:9000 --backend 127.0.0.1:1080

测试 HTTP 代理::

  curl --proxy http://127.0.0.1:9000 http://httpbin.org/ip


SOCKS
-----------

启动纯 socks v5 代理::

  goproxy socks --listen 127.0.0.1:9000

测试 SOCKS 代理::

  curl --socks5 127.0.0.1:9000 http://httpbin.org/ip


其他资源
================

- qtunnel
- ssh
- socat


常用工具设置代理
======================

git
----------

http(s) 协议
~~~~~~~~~~~~~~~~~~~~~~

参考:

- `Getting git to work with a proxy server <http://stackoverflow.com/questions/783811/getting-git-to-work-with-a-proxy-server>`_

如果 git repo 地址开头为 http 或 https , 如::

  git clone https://github.com/ooclab/goproxy

配置全局设置::

  git config --global http.proxy http://proxyuser:proxypwd@proxy.server.com:8080
  git config --global https.proxy https://proxyuser:proxypwd@proxy.server.com:8080

取消代理设置::

  git config --global --unset http.proxy
  git config --global --unset https.proxy

git 协议
~~~~~~~~~~~~~~~~~~

参考:

- `How to Use the Git Protocol Through a HTTP CONNECT Proxy <http://www.emilsit.net/blog/archives/how-to-use-the-git-protocol-through-a-http-connect-proxy/>`_

http 代理::

  exec socat STDIO PROXY:$_proxy:$1:$2,proxyport=$_proxyport

socks 代理::

  exec socat STDIO SOCKS4:$_proxy:$1:$2,socksport=$_proxyport

tips
~~~~~~

查看 git config 设置::

  git config -l
