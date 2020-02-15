package main

import (
	"blog/conf"
	"blog/router"
	"fmt"
	"log"
)

func main() {
	r := router.NewRouter()
	if err := r.Run(fmt.Sprintf(":%d",conf.ServerConfig.Port));err != nil{
		log.Fatal(err)
	}
}
