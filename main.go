package main

import (
	"gin-api-example/routers"
)

func main() {
	router := routers.InitRouter()
	router.SetTrustedProxies(nil)
	router.Run(":3000")
}
