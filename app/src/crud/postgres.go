package crud

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"reflect"
	"sync"
	"time"
)

type postgresConn struct {
	conn *gorm.DB
}

var instance *postgresConn
var once sync.Once

func GetPostgresConn() *postgresConn {
	once.Do(func() {
		session, err := createSession()
		if err != nil {
			log.Panicf("Cannot create a connection to postgres")
		}
		instance = &postgresConn{
			conn: session,
		}
	})
	return instance
}

type User struct {
	ID uint

	Name    string
	Address string
	Age     uint

	CreatedAt time.Time
	UpdatedAt time.Time
}

func PostgresPoc() {
	//db, err := CreateSession()
	//if err != nil {
	//	log.Println(err)
	//}
	db := GetPostgresConn().conn

	_ = Migrate(db, User{})
	res := Find(db, User{}, "name = ?", "Anurag")
	tx := Create(db, User{
		ID:      3,
		Name:    "Anurag",
		Address: "825 Geneva Ave",
		Age:     33555,
	})
	log.Println(tx)
	res = Find(db, User{}, "name = ?", "Anurag")
	log.Println(res)
	tx = Update(db, res, User{Name: "Anurag", Age: 12345})
	log.Println(tx)
	res = Find(db, User{}, "name = ?", "Anurag")
	log.Println(res)
	//db.Delete(res, res.(User).ID)
	//Delete(db, res, res.(User).ID)
	res = Find(db, User{}, "name = ?", "Anurag")
}

func createSession() (*gorm.DB, error) {
	dsn := "host=postgres user=postgres password=changeme dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("err:", err)
	}
	return db, err
}

func Migrate(db *gorm.DB, obj interface{}) error {
	// Migrate the schema
	user, ok := obj.(User)
	if ok {
		fmt.Printf("Hello %s!\n", user.Name)
	}
	err := db.AutoMigrate(user)
	if err != nil {
		fmt.Println("err:", err)
	}
	return err
}

func Create(db *gorm.DB, obj interface{}) *gorm.DB {
	// Create
	user, ok := obj.(User)
	if ok {
		fmt.Printf("Hello %s!\n", user.Name)
	}
	tx := db.Create(&user)
	return tx
}

func Find(db *gorm.DB, obj interface{}, conds ...interface{}) interface{} {
	// Read
	//var user User
	//db.Find(&user, "name = ?", "Anurag")
	user, ok := obj.(User)
	if ok {
		fmt.Printf("Hello %s!\n", user.Name)
	}
	db.Find(&user, conds...)
	//db.Find(&user, "name = ?", "Anurag")
	fmt.Println(user)
	return user
}

func Update(db *gorm.DB, obj interface{}, updateValues interface{}) *gorm.DB {
	// Update
	//db.Model(&user).Update("age", user.Age-1)
	//db.Model(&user).Updates(User{Name: "Anu", Age: 25}) // non-zero fields
	if reflect.TypeOf(obj) != reflect.TypeOf(updateValues) {
		return nil
	}
	//db.Model(&user).Updates(map[string]interface{}{"Name": "Anu", "Age": 25})
	//db.Find(&user, "name = ?", "Anurag")
	tx := db.Model(obj).Updates(updateValues)
	fmt.Println(tx)
	return tx
}

func Delete(db *gorm.DB, obj interface{}, conds interface{}) *gorm.DB {
	// Delete
	tx := db.Delete(obj, conds)
	return tx
}
