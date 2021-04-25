package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID uint

	Name    string
	Address string
	Age     uint

	CreatedAt time.Time
	UpdatedAt time.Time
}

func PostgresPoc() {

	dsn := "host=localhost user=postgres password=changeme dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("err:", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("err:", err)
	}

	// Create
	db.Create(&User{ID: 1, Name: "Anurag", Address: "sf", Age: 32})

	// Read
	var user User
	db.Find(&user, "name = ?", "Anurag")
	fmt.Println(user)

	// Update
	db.Model(&user).Update("age", user.Age-1)
	//db.Model(&user).Updates(User{Name: "Anu", Age: 25}) // non-zero fields
	//db.Model(&user).Updates(map[string]interface{}{"Name": "Anu", "Age": 25})
	db.Find(&user, "name = ?", "Anurag")
	fmt.Println(user)

	// Delete
	//db.Delete(&user, user.ID)
}
