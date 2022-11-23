package main

import (
	"database/sql"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// gorm.Model definition
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User2 struct {
	Name string `gorm:"<-:create"` // allow read and create
	// Name string `gorm:"<-:update"`          // allow read and update
	// Name string `gorm:"<-"`                 // allow read and write (create and update)
	// Name string `gorm:"<-:false"`           // allow read, disable write permission
	// Name string `gorm:"->"`                 // readonly (disable write permission unless it configured)
	// Name string `gorm:"->;<-:create"`       // allow read and create
	// Name string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
	// Name string `gorm:"-"`                  // ignore this field when write and read with struct
	// Name string `gorm:"-:all"`              // ignore this field when write, read and migrate with struct
	// Name string `gorm:"-:migration"`        // ignore this field when migrate with struct
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)

}
