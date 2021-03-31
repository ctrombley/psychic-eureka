package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ctrombley/psychic-eureka/internal/service"
)

func main() {
	r := gin.Default()
	service.initRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
