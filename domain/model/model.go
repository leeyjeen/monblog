package model

/* entities layer */
import "time"

// User this is a model for user
type User struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	Name         string    `gorm:"type:varchar(15); not null" json:"name"`
	Email        string    `gorm:"type:varchar(50); not null; unique" json:"email"`
	PasswordHash string    `gorm:"type:varchar(32); not null; default:''"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Post this is a model for post
type Post struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	author    User      `gorm:"ForeignKey:AuthorID"`
	AuthorID  uint      `gorm:"not null" json:"author_id"`
	Title     string    `gorm:"type:varchar(30); not null" json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Content   string    `gorm:"type:text; not null" json:"content"`
}

// PostComment this is a model for post comment
type PostComment struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	post      Post      `gorm:"ForeignKey:PostID"`
	PostID    uint      `gorm:"not null" json:"post_id"`
	ParentID  uint      `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Content   string    `gorm:"type:text; not null" json:"content"`
}

// Category this is a model for category
type Category struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	ParentID uint   `json:"parent_id"`
	Name     string `gorm:"type:varchar(15); not null" json:"name"`
}

// PostCategory this is a model for post category
type PostCategory struct {
	post       Post     `gorm:"ForeignKey:PostID"`
	PostID     uint     `json:"post_id"`
	category   Category `gorm:"ForeignKey:CategoryID"`
	CategoryID uint     `json:"category_id"`
}
