package main

import (
	"log"
	"sysWeb/routers"
)

func main() {
	r := routers.SetupRouter()
	if err := r.Run("localhost:8080"); err != nil {
		log.Fatalf("service start failed, err:%v\n", err)
	}
}
