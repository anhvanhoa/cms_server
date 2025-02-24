package bootstrap

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Env struct {
	MODE_ENV string

	PORT_DB string
	NAME_DB string
	USER_DB string
	PASS_DB string
	HOST_DB string
	URL_DB  string

	NAME_APP string
	PORT_APP string
}

func NewEnv() *Env {
	env := &Env{}
	absPath, err := filepath.Abs("./")
	if err != nil {
		log.Fatal("Error getting the absolute path:", err)
	}

	mode := os.Getenv("ENV_MODE")
	viper.SetConfigType("yaml")
	if mode == "production" {
		viper.SetConfigName("prod.config")
	} else {
		viper.SetConfigName("dev.config")
	}
	viper.AddConfigPath(absPath)
	err = viper.ReadInConfig()
	if err != nil {
		panic("Error reading config file, " + err.Error())
	}

	err = viper.UnmarshalExact(env)
	if err != nil {
		panic("Error unmarshalling config file, " + err.Error())
	}

	return env
}

func (env *Env) IsProduction() bool {
	return strings.ToLower(env.MODE_ENV) == "production"
}
