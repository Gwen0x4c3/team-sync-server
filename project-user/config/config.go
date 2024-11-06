package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	Zap    ZapConfig
	Redis  RedisConfig
	Grpc   GrpcConfig
}

type ServerConfig struct {
	Name string
	Addr string
}

type ZapConfig struct {
	DebugFileName string
	InfoFileName  string
	WarnFileName  string
	MaxSize       int
	MaxAge        int
	MaxBackups    int
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type GrpcConfig struct {
	Name string
	Addr string
}

var Cfg Config

func init() {
	log.Println("Init project-user config")
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./project-user/config")
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	log.Println("config loaded")
}
