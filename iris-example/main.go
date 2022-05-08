package main

import (
	"thinkdecideact/src/config"
	"thinkdecideact/src/routers"
	"thinkdecideact/src/utils"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	config := config.GetConfig()
	DB := utils.GetDB(config.DB)

	app.PartyFunc("/api", routers.ApiRouter(DB))

	app.Listen(":8080")
}
