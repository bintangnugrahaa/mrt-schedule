package main

import (
	"github.com/bintangnugrahaa/mrt-schedule/modules/station"
	"github.com/gin-gonic/gin"
)

func main() {
	InitiateRoute()
}

func InitiateRoute() {
	var (
		route = gin.Default()
		api   = route.Group("/v1/api")
	)
	station.Initiate(api)

	route.Run(":8080")
}
