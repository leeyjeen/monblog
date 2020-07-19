package datastore

/* Frameworks and Drives Layer */
import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDB() *gorm.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PW")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8,utf8mb4&parseTime=true", user, password, host, port, dbName)
	fmt.Println("dbString ", dbString)
	db, err := gorm.Open("mysql", dbString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
