package main

import (
	"github.com/SaenkoDmitry/advertisement-app/db"
	"github.com/SaenkoDmitry/advertisement-app/routers"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/SaenkoDmitry/advertisement-app/docs"
)

// @title Swagger Advertisement store API
// @version 1.0
// @description This is a server for advertisement store.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://github.com/SaenkoDmitry/advertisement-app

// @host localhost:9100
func main() {
	db.Connect()
	defer db.GetDB().Close()

	r := gin.Default()
	r.POST("/adv", routers.AdvCreate)
	r.GET("/adv/:id", routers.AdvGet)
	r.GET("/adv", routers.AdvGetAll)
	r.Static("/swaggerui/", "swaggerui")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":9100")
}
