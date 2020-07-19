package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/leeyjeen/monblog/infrastructure/datastore"
	r "github.com/leeyjeen/monblog/infrastructure/router"
	"github.com/leeyjeen/monblog/registry"
)

var (
	router *gin.Engine
	db     *gorm.DB
)

func main() {
	dbManager, err := datastore.NewDBManager()
	if err != nil {
		log.Fatalf("%v", err)
	}
	db = dbManager.Connect()
	defer dbManager.Close(db)
	if err := dbManager.Migrate(db, Migrations); err != nil {
		log.Fatalf("%v", err)
	}

	regi := registry.NewRegistry(db)

	router = r.NewRouter(regi.NewAppController())

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("API Port is not set..")
	}

	router.Run(":" + port)
}
