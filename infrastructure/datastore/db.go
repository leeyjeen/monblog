package datastore

/* Frameworks and Drives Layer */
import (
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

type dbManager struct {
}

type DBManager interface {
	Connect() *gorm.DB
	Close(db *gorm.DB) error
	Migrate(db *gorm.DB, migrations []*gormigrate.Migration) error
}

func NewDBManager() (DBManager, error) {
	var m DBManager = &dbManager{}
	val, ok := m.(DBManager)
	if !ok {
		return nil, errors.New("NewDBManager() : invalid structure")
	}
	return val, nil
}

func (m *dbManager) Connect() *gorm.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PW")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8,utf8mb4&parseTime=true", user, password, host, port, dbName)
	db, err := gorm.Open("mysql", dbString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (m *dbManager) Close(db *gorm.DB) error {
	return db.Close()
}

func (m *dbManager) Migrate(db *gorm.DB, migrations []*gormigrate.Migration) error {
	db.LogMode(true)
	gm := gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	return gm.Migrate()
}
