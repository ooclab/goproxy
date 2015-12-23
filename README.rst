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

  goproxy http -l 127.0.0.1:9000

启动 http(s) 代理，但使用一个 socks v5 服务作为后端::

  goproxy http -l 127.0.0.1:9000 -b 127.0.0.1:1080

测试 HTTP 代理::

  curl --proxy http://127.0.0.1:9000 http://httpbin.org/ip


SOCKS
-----------

启动纯 socks v5 代理::

  goproxy socks -l 127.0.0.1:9000

测试 SOCKS 代理::

  curl --socks5 127.0.0.1:9000 http://httpbin.org/ip


其他资源
================

- qtunnel
- ssh
