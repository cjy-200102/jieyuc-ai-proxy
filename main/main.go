package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"jieyuc-ai-proxy/routes"
	"log"
)

// main program entrance
func main() {
	r := gin.Default()
	viper.SetConfigFile("./bootstrap.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	port := viper.GetInt("server.port")
	routes.InitHttpRoutes(r)
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Println("启动失败……")
		panic(err)
	}
}
