package database

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/huyinghuan/app/istorage"
	"gopkg.in/JX3BOX/gologger.v2"
	logXorm "gopkg.in/JX3BOX/gologger.v2/xorm"
	"xorm.io/xorm"
)

var enginPool = make(map[string]*xorm.Engine)

// var mainEngine *xorm.Engine
var cacheDB *redis.Client

func InitMysql(conf istorage.IDBConfig, names ...string) {
	name := "default"
	if len(names) > 0 {
		name = names[0]
	}
	if _, ok := enginPool[name]; ok {
		return
	}
	conn, connectErr := xorm.NewEngine(conf.GetDriver(), conf.GetConnect())
	if connectErr != nil {
		panic(connectErr)
	}
	conn.ShowSQL(conf.GetShowSQL())
	conn.SetLogger(logXorm.XormLogger(gologger.GetSingleInstance()))
	conn.SetLogLevel(2) // warning
	enginPool[name] = conn
}

func InitRedis(redisConfig istorage.IRedisConfig) {
	cacheDB = redis.NewClient(&redis.Options{
		Addr:     redisConfig.GetURL(),
		Password: redisConfig.GetPassword(),
		// Database to be selected after connecting to the server.
		DB: redisConfig.GetDBIndex(),
	})
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := cacheDB.Ping(ctx).Err(); err != nil {
		panic(err)
	} else {
		gologger.Info("Redis connected:" + redisConfig.GetURL())
	}
}

func GetDriver(names ...string) *xorm.Engine {
	name := "default"
	if len(names) > 0 {
		name = names[0]
	}
	return enginPool[name]
}

func GetRedis() *redis.Client {
	return cacheDB
}
