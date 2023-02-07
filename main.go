package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
)

var (
	ip           string
	port         string
	staticFolder string
	debugFlag    bool
)

func main() {
	flag.StringVar(&ip, "bind_ip", "0.0.0.0", "server ip")
	flag.StringVar(&port, "port", "12155", "server port")
	flag.StringVar(&staticFolder, "static", "./static/", "static folder path")
	flag.BoolVar(&debugFlag, "debug", false, "debug mode")
	flag.Parse()

	if debugFlag {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Static("/", staticFolder)
	ipAdd := net.JoinHostPort(ip, port)
	fmt.Printf("start server at %s \n", ipAdd)
	err := r.Run(ipAdd)
	if err != nil {
		fmt.Println(err)
		return
	}
}
