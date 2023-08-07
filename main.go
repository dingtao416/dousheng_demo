package main

import (
	"fmt"
	"log"

	"github.com/abuziming/dousheng_demo/config"
	"github.com/abuziming/dousheng_demo/service"
)

func main() {
	go service.RunMessageServer()

	r := initRouter()

	err := r.Run(fmt.Sprintf(":%d", config.Global.Server.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Panicln(err)
	}
}
