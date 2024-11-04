package main

import (
	"context"
	db "gin-api-example/database"
	"gin-api-example/routers"
)

func main() {
	conn := db.InitDbConn()
	defer conn.Close(context.Background())
	router := routers.InitRouter()
	router.SetTrustedProxies(nil)
	router.Run(":3000")
}
