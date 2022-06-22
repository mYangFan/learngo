package main

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

/**
once.Do(),只执行一次
 */
type Config struct {
	Server string
	Port   int64
}

var (
	once sync.Once
	configg *Config
)
func main() {
	for i := 0; i < 10; i++ {
		go func() {
			_ = ReadConfig()
		}()
	}
	time.Sleep(time.Second)
}

func ReadConfig() *Config {
	once.Do(func() {
		var err error
		configg = &Config{Server: os.Getenv("TT_SERVER_URL")}
		configg.Port, err = strconv.ParseInt(os.Getenv("TT_PORT"), 10,0)
		if err != nil {
			configg.Port = 8080
		}

		log.Println("init config")
	})
	return configg
}


