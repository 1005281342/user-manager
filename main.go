package main

import (
	"github.com/1005281342/user-manager/auth"
	"github.com/gin-gonic/gin"

	"github.com/1005281342/user-manager/api/routes"
	"github.com/1005281342/user-manager/db"
)

func main() {
	err := db.Connect()
	if err != nil {
		panic(err)
	}

	//err = cache.Connect()
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = search.Connect()
	//if err != nil {
	//	panic(err)
	//}

	r := gin.Default()

	jwtAuth := auth.NewJWTAuth("secret-key")
	routes.SetupUserRoutes(r, jwtAuth)
	routes.SetupAuthRoutes(r)

	r.Run()
}
