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
)

func main() {
	ip = *flag.String("bind_ip", "0.0.0.0", "server ip")
	port = *flag.String("port", "12155", "server port")
	staticFolder = *flag.String("static", "./static/", "static folder path")
	flag.Parse()

	r := gin.Default()
    r.Static("/", staticFolder)
	ipAdd := net.JoinHostPort(ip, port)
	err := r.Run(ipAdd)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("启动成功...")
}
