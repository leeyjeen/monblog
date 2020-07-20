package model

/* entities layer */
import "time"

// User this is a model for user
type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Post this is a model for post
type Post struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	AuthorID  uint      `json:"author_id"`
	ParentID  uint      `json:"parent_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Content   string    `json:"content"`
}

// PostComment this is a model for post comment
type PostComment struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	PostID    uint      `json:"post_id"`
	ParentID  uint      `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Content   string    `json:"content"`
}

// Category this is a model for category
type Category struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	ParentID uint   `json:"parent_id"`
	Name     string `json:"name"`
}

// PostCategory this is a model for post category
type PostCategory struct {
	PostID     uint `json:"post_id"`
	CategoryID uint `json:"category_id"`
}
