package istorage

type IDBConfig interface {
	GetDriver() string
	GetConnect() string
	GetShowSQL() bool
}

type IRedisConfig interface {
	GetURL() string
	GetPassword() string
	GetDBIndex() int
}
