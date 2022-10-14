package config

import (
	"fmt"
	"os"

	"github.com/huyinghuan/app/bundle"
	"gopkg.in/yaml.v2"
)

type DBConfig struct {
	Driver   string `yaml:"driver"`
	Connect  string `yaml:"connect"`
	ShowSQL  bool   `yaml:"show_sql"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Hosts    string `yaml:"hosts"`
	DBName   string `yaml:"db_name"`
}

func (c *DBConfig) GetDriver() string {
	return c.Driver
}

func (c *DBConfig) GetConnect() string {
	if c.Connect == "" {
		return fmt.Sprintf("%s:%s@tcp(%s)/%s", c.User, c.Password, c.Hosts, c.DBName)
	}
	return c.Connect
}

func (c *DBConfig) GetShowSQL() bool {
	return c.ShowSQL
}

type ExtraDb struct {
	MainConnect DBConfig `yaml:"main"`
}

type RedisConfig struct {
	URL      string `yaml:"url"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func (r *RedisConfig) GetURL() string {
	return r.URL
}

func (r *RedisConfig) GetPassword() string {
	return r.Password
}

func (r *RedisConfig) GetDBIndex() int {
	return r.DB
}

type Config struct {
	CorsAllowedOrigins []string    `yaml:"corsAllowedOrigins"`
	Redis              RedisConfig `yaml:"redis"`
	Db                 DBConfig    `yaml:"db"`
}

var config *Config

// init 读取配置文件
func ReadFromFile(env string) string {
	runmode := env
	if runmode == "" {
		runmode = os.Getenv("RUN_MODE")
	}
	configPath := fmt.Sprintf("asserts/config.%s.yml", runmode)
	body := bundle.ReadFile(configPath)
	err := yaml.Unmarshal(body, &config)
	if err != nil {
		panic(err)
	}
	return configPath
}

// Get 获取配置文件
func Get() *Config {
	return config
}
