package tests

import (
	"fmt"
	config "github.com/ubrabbit/go-public/config"
	"testing"
)

const filepath = "config_test.conf"

func TestConfig(t *testing.T) {
	fmt.Printf("\n\n=====================  TestConfig  =====================\n")
	config.InitConfig(filepath)
	config.InitMongoDBConfig()
	config.InitMysqlConfig()
	config.InitRabbitMQConfig()
	config.InitRedisConfig()
}

func TestConfig2(t *testing.T) {
	fmt.Println("GetPath(run) == ", config.GetPath("run"))
	fmt.Println("GetPath(log) == ", config.GetPath("log"))
	fmt.Println("GetPath(data) == ", config.GetPath("data"))
	fmt.Println("GetPath(cache) == ", config.GetPath("cache"))
}
