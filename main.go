package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net"
)

var (
	ip           string
	port         string
	staticFolder string
	debugFlag    bool

	err error
)

func main() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("read config error...")
	}

	pflag.String("bind_ip", "0.0.0.0", "server ip")
	pflag.String("port", "12155", "server port")
	pflag.String("static", "./static", "static folder path")
	pflag.Bool("debug", false, "debug mode")
	pflag.Parse()
	err = viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		fmt.Println("read params from command line error...")
	}
	ip = viper.GetString("bind_ip")
	port = viper.GetString("port")
	staticFolder = viper.GetString("static")
	debugFlag = viper.GetBool("debug")

	if debugFlag {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Static("/", staticFolder)
	ipAdd := net.JoinHostPort(ip, port)

	fmt.Printf("start server at %s \n", ipAdd)
	err = r.Run(ipAdd)
	if err != nil {
		fmt.Println(err)
		return
	}
}
