package irisapptest

import (
	"testing"

	"github.com/kataras/iris/v12"
)

func TestDemoCtrlDemo(t *testing.T) {
	e := NewHttpExcept(t)
	e.GET("/demo").Expect().Status(iris.StatusOK).Body().Equal("hello world")
}
