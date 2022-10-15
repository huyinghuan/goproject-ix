package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/huyinghuan/app/config"
	"github.com/huyinghuan/app/irisapp"
	"gopkg.in/JX3BOX/gologger.v2"
)

var (
	Version   = "Development"
	BuildTime = "Dev"
	Author    = "ec.huyinghuan@gmail.com"
)

func main() {
	var args config.StartArgs
	flag.IntVar(&args.Port, "port", 8080, "端口号")
	flag.StringVar(&args.LogLevel, "log", "info", "日志等级 debug,info, warn, error")
	flag.StringVar(&args.Env, "env", "product", "环境变量")
	flag.Parse()
	if err := args.Vaild(); err != nil {
		panic(err)
	}
	confFile := config.ReadFromFile(args.Env)
	gologger.InitLogger(gologger.Level(args.GetZapLogLevel()))
	gologger.Infow("config file", "path", confFile)

	app := irisapp.NewApp(Version)

	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		<-gracefulStop
		gologger.Warn("服务关闭中...")
		app.Shutdown(context.Background())
		time.Sleep(time.Second)
	}()

	app.Listen(args.GetPort())
}
