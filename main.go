package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/leeyjeen/monblog/infrastructure/datastore"
	r "github.com/leeyjeen/monblog/infrastructure/router"
	"github.com/leeyjeen/monblog/registry"
)

var router *gin.Engine

func main() {
	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	regi := registry.NewRegistry(db)

	router = r.NewRouter(regi.NewAppController())

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("API Port is not set..")
	}

	router.Run(":" + port)
}
