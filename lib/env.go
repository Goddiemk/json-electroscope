package lib

import (
	"fmt"
	"sync"

	"github.com/joho/godotenv"
	envconfig "github.com/kelseyhightower/envconfig"
)

type Envs struct {
	AppName  string
	MYSQLURI string
	Port     string
}

var envInstance Envs
var envOnce sync.Once
var environ = GetEnvironment()

func GetEnvironment() Envs {
	envOnce.Do(func() {
		var appname = "elect"
		if err := godotenv.Load("system.env"); err != nil {
			fmt.Println(err.Error())
		}

		if err := envconfig.Process(appname, &envInstance); err != nil {
			fmt.Println(err)
		}

		envInstance.AppName = appname
	})
	return envInstance
}
