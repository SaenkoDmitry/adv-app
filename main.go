package main

import (
	"github.com/SaenkoDmitry/advertisement-app/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/adv", routers.AdvCreate)
	r.GET("/adv/:id", routers.AdvGet)
	r.GET("/adv", routers.AdvGetAll)
	r.Run(":9100")
}
