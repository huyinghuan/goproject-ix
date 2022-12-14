package irisapp

import (
	"github.com/huyinghuan/app/irisapp/controller/democtrl"
	"github.com/kataras/iris/v12"
)

func NewApp(ver string) *iris.Application {
	app := iris.New()
	app.Get("/version", func(ctx iris.Context) {
		ctx.WriteString(ver)
	})
	app.Get("/demo", democtrl.Demo)
	return app
}
