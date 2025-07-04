package bootstrap

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type queue struct {
	Addr        string
	DB          int
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
	Network     string
	Concurrency int
	Queues      map[string]int
}

type jwtSecret struct {
	Access  string
	Refresh string
	Verify  string
	Forgot  string
}

type dbCache struct {
	Addr        string
	DB          int
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
	Network     string
}

type Env struct {
	MODE_ENV string

	URL_DB string

	NAME_APP string
	PORT_APP string
	HOST_APP string

	QUEUE    *queue
	DB_CACHE *dbCache

	SECRET_OTP string

	JWT_SECRET *jwtSecret

	FRONTEND_URL string

	TEST_EMAIL string
}

func NewEnv(env any) {
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
}

func (env *Env) IsProduction() bool {
	return strings.ToLower(env.MODE_ENV) == "production"
}
