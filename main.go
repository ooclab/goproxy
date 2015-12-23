package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "goproxy"
	app.Usage = "非常容易使用的代理工具"

	listenFlag := cli.StringFlag{
		Name:  "listen, l",
		Value: "127.0.0.1:9000",
		Usage: "代理监听地址",
	}

	app.Commands = []cli.Command{
		{
			Name:    "http",
			Aliases: []string{"hp"},
			Usage:   "启动 HTTP(s) 代理",
			Action: func(c *cli.Context) {
				listen := c.String("listen")
				backend := c.String("backend")
				if backend == "" {
					startHttpProxy(listen)
				} else {
					startHttpProxyByBackend(listen, backend)
				}
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "backend, b",
					Value:  "",
					Usage:  "匹配后端 socks v5 代理",
					EnvVar: "SOCKS_PROXY",
				},
				listenFlag,
			},
		},
		{
			Name:    "socks",
			Aliases: []string{"s"},
			Usage:   "启动 socks(v5) 代理",
			Action: func(c *cli.Context) {
				listen := c.String("listen")
				startSocksV5(listen)
			},
			Flags: []cli.Flag{listenFlag},
		},
	}

	app.Run(os.Args)
}
