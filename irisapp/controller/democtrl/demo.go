package democtrl

import "github.com/kataras/iris/v12"

func Demo(ctx iris.Context) {
	ctx.WriteString("hello world")
}
