package main

import (
	"github.com/ctrombley/psychic-eureka/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	service.InitRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
