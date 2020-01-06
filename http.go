package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
	"h12.io/socks"
)

// 启动本地 http(s) 代理
func startHttpProxy(listen string) {
	log.Printf("try start HTTP(s) %s, no backend\n", listen)

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(listen, proxy))
}

// 启动 http(s) 代理，使用一个 socks v5 服务器作为后端
func startHttpProxyByBackend(listen string, backend string) {
	log.Printf("try start HTTP(s) %s, use backend %s\n", listen, backend)

	var tlsClientSkipVerify = &tls.Config{InsecureSkipVerify: true}
	dialSocksProxy := socks.DialSocksProxy(socks.SOCKS5, backend)

	proxy := goproxy.NewProxyHttpServer()
	proxy.Tr = &http.Transport{
		TLSClientConfig: tlsClientSkipVerify,
		Proxy:           http.ProxyFromEnvironment,
		Dial:            dialSocksProxy}
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(listen, proxy))
}
