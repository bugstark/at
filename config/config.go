package config

import (
	"flag"
	"log"
	"os"

	"github.com/spf13/viper"
)

var c *viper.Viper

func init() {
	var config string
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
		if configEnv := os.Getenv("ATCONFIG"); configEnv == "" {
			config = "config.yaml"
			log.Printf("正在使用config的默认值\n")
		} else {
			config = configEnv
			log.Printf("正在使用ATCONFIG环境变量,config的路径为%v\n", config)
		}
	} else {
		log.Printf("正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	v.WatchConfig()
	c = v
}

func Get() *viper.Viper {
	return c
}
