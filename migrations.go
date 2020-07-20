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
	{
		ID: "202007202350",
		Migrate: func(tx *gorm.DB) error {
			user := new(model.User)
			post := new(model.Post)
			postComment := new(model.PostComment)
			category := new(model.Category)
			postCategory := new(model.PostCategory)
			if err := tx.Model(user).DropColumn("age").Error; err != nil {
				return err
			}
			return tx.AutoMigrate(user, post, postComment, category, postCategory).Error
		},
		Rollback: func(tx *gorm.DB) error {
			user := new(model.User)
			post := new(model.Post)
			postComment := new(model.PostComment)
			category := new(model.Category)
			postCategory := new(model.PostCategory)
			return tx.DropTableIfExists(user, post, postComment, category, postCategory).Error
		},
	},
	{
		ID: "202007210010",
		Migrate: func(tx *gorm.DB) error {
			user := new(model.User)
			post := new(model.Post)
			postComment := new(model.PostComment)
			postCategory := new(model.PostCategory)

			if err := tx.Model(user).AddUniqueIndex("email", "email").Error; err != nil {
				return err
			}
			if err := tx.Model(post).AddForeignKey("author_id", "users(id)", "CASCADE", "CASCADE").Error; err != nil {
				return err
			}
			if err := tx.Model(postComment).AddForeignKey("post_id", "posts(id)", "CASCADE", "CASCADE").Error; err != nil {
				return err
			}
			if err := tx.Model(postCategory).AddForeignKey("post_id", "posts(id)", "CASCADE", "CASCADE").Error; err != nil {
				return err
			}
			if err := tx.Model(postCategory).AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE").Error; err != nil {
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			user := new(model.User)
			post := new(model.Post)
			postComment := new(model.PostComment)
			postCategory := new(model.PostCategory)
			if err := tx.Model(user).RemoveIndex("email").Error; err != nil {
				return err
			}
			if err := tx.Model(post).RemoveForeignKey("author_id", "users(id)").Error; err != nil {
				return err
			}
			if err := tx.Model(postComment).RemoveForeignKey("post_id", "posts(id)").Error; err != nil {
				return err
			}
			if err := tx.Model(postCategory).RemoveForeignKey("post_id", "posts(id)").Error; err != nil {
				return err
			}
			if err := tx.Model(postCategory).RemoveForeignKey("category_id", "categories(id)").Error; err != nil {
				return err
			}
			return nil
		},
	},
}
