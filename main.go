package main

import (
	"context"
	"github.com/joho/godotenv"
	db "gin-api-example/database"
	"gin-api-example/routers"
)

func main() {
	godotenv.Load()
	conn := db.InitDbConn()
	defer conn.Close(context.Background())
	router := routers.InitRouter()
	router.SetTrustedProxies(nil)
	router.Run(":3000")
}
