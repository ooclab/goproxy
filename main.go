package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd 是主命令对象
var rootCmd = &cobra.Command{
	Use:   "goproxy SUBCOMMAND ARGS",
	Short: "非常容易使用的代理工具",
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	cobra.OnInitialize(initRootConfig)
}

func initRootConfig() {
	viper.SetEnvPrefix("SUPERVISOR") // will be uppercased automatically
	viper.BindEnv("DEBUG")
	viper.BindPFlag("debug", rootCmd.Flags().Lookup("verbose"))

	verbose := viper.GetBool("debug")
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func main() {

	// For http
	var httpCmd = &cobra.Command{
		Use:   "http ARGS",
		Short: "启动 HTTP(s) 代理",
		Run: func(cmd *cobra.Command, args []string) {
			listen := viper.GetString("http_listen")
			backend := viper.GetString("backend")
			if backend == "" {
				startHttpProxy(listen)
			} else {
				startHttpProxyByBackend(listen, backend)
			}
		},
	}

	httpCmd.Flags().String("backend", "", "匹配后端 socks v5 代理")
	viper.BindPFlag("SOCKS_PROXY", httpCmd.Flags().Lookup("backend"))

	httpCmd.Flags().String("listen", "127.0.0.1:9000", "代理监听地址")
	viper.BindPFlag("HTTP_LISTEN", httpCmd.Flags().Lookup("listen"))

	// For socks
	var socksCmd = &cobra.Command{
		Use:   "socks ARGS",
		Short: "启动 socks(v5) 代理",
		Run: func(cmd *cobra.Command, args []string) {
			listen := viper.GetString("socks_listen")
			startSocksV5(listen)
		},
	}
	socksCmd.Flags().String("listen", "127.0.0.1:8000", "代理监听地址")
	viper.BindPFlag("SOCKS_LISTEN", socksCmd.Flags().Lookup("listen"))

	rootCmd.AddCommand(httpCmd)
	rootCmd.AddCommand(socksCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
