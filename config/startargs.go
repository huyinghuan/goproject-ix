package config

import (
	"errors"
	"fmt"
)

var zapLogLevelMap map[string]int = map[string]int{
	"debug": -1,
	"info":  0,
	"warn":  1,
	"error": 2,
}

type StartArgs struct {
	Env      string
	Port     int
	LogLevel string
}

func (s *StartArgs) GetPort() string {
	return fmt.Sprintf(":%d", s.Port)
}

func (s *StartArgs) GetZapLogLevel() int {
	if level, ok := zapLogLevelMap[s.LogLevel]; ok {
		return level
	}
	return 0
}

func (s *StartArgs) Vaild() error {
	if s.Port < 1024 || s.Port > 65535 {
		return errors.New("port range is 1024 to 65535")
	}
	return nil
}
