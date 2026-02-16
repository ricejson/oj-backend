package main

import (
	"context"
	"fmt"

	"github.com/ricejson/oj-backend/service/judge/sandbox"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func InitViper() {
	// 读取命令行参数
	s := pflag.String("config", "", "config file path")
	pflag.Parse()
	// 设置文件路径
	viper.SetConfigFile(*s)
	// 读取配置信息
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	InitViper()
	instance := sandbox.NewInstance(viper.GetString("sandbox.type"))
	response, err := instance.ExecuteCode(context.Background(), &sandbox.ExecuteCodeRequest{
		Code:         "main",
		InputSamples: []string{"1 2", "3, 4"},
	})
	fmt.Println(response, err)
}
