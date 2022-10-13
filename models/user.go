package models

import (
	"time"

	"gorm.io/gorm"
)

// User type
type UserType struct {
	gorm.Model
	Name        string
	LevelAccess int
}

// User model
type User struct {
	gorm.Model
	Name       string
	Email      *string
	Birthday   *time.Time
	Github     string
	UserTypeID int
	UserType   UserType
}

// Field
type Field struct {
	gorm.Model
	Name string
}

// Tags
type Tags struct {
	gorm.Model
	Name    string
	FieldID int
	Field   Field
}

// Post
type Post struct {
	gorm.Model
	UserId  string
	Title   string
	Content string
	Tags    Tags `gorm:"many2many:post_tags"`
}
