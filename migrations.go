package main

import (
	"github.com/jinzhu/gorm"
	"github.com/leeyjeen/monblog/domain/model"
	"gopkg.in/gormigrate.v1"
)

var Migrations = []*gormigrate.Migration{
	{
		ID: "202007191545",
		Migrate: func(tx *gorm.DB) error {
			user := new(model.User)
			return tx.AutoMigrate(user).Error
		},
		Rollback: func(tx *gorm.DB) error {
			user := new(model.User)
			return tx.DropTableIfExists(user).Error
		},
	},
}
