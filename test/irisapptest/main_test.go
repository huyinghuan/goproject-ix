package irisapptest

import (
	"os"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/huyinghuan/app/config"
	"github.com/huyinghuan/app/irisapp"
	"github.com/huyinghuan/app/irisapp/setup"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"gopkg.in/JX3BOX/gologger.v2"

	_ "github.com/go-sql-driver/mysql"
)

var isOnline = false
var isDev = true

func getDomain() string {
	if isDev {
		return "http://localhost:15186"
	}
	return "https://exmaple.com"
}

var app *iris.Application

func NewHttpExcept(t *testing.T) *httpexpect.Expect {
	if isOnline {
		return httpexpect.New(t, getDomain())
	} else {
		if app == nil {
			app = irisapp.NewApp("Test")
		}
		return httptest.New(t, app)
	}
}

func TestMain(m *testing.M) {
	confFile := config.ReadFromFile("dev")
	gologger.InitLogger(gologger.DebugLevel)
	gologger.Infow("config file", "path", confFile)
	setup.Do()
	os.Exit(m.Run())
}
