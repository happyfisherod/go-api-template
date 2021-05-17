package crud

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

type postgresConn struct {
	conn *gorm.DB
}

var postgresInstance *postgresConn
var postgresOnce sync.Once

func NewPostgresConn(dsn string) (*postgresConn, error) {
	session, err := createSession(dsn)
	if err != nil {
		log.Println("Cannot create a connection to postgres", err)
	}
	postgresInstance = &postgresConn{
		conn: session,
	}
	return postgresInstance, err
}

func GetPostgresConn() *postgresConn {
	postgresOnce.Do(func() {
		dsn := "host=postgres user=postgres password=changeme dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		// TODO: create dsn string from env variables
		session, err := createSession(dsn)
		if err != nil {
			log.Fatal("Cannot create a connection to postgres", err)
		}
		postgresInstance = &postgresConn{
			conn: session,
		}
	})
	return postgresInstance
}

func createSession(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("err:", err)
	}
	return db, err
}

///////////////////

func PostgresPoc() {
	//db := GetPostgresConn().conn
	//
	//_ = Migrate(db, User{})
	//res := Find(db, User{}, "name = ?", "Anurag")
	//tx := Create(db, User{
	//	ID:      3,
	//	Name:    "Anurag",
	//	Address: "825 Geneva Ave",
	//	Age:     33555,
	//})
	//log.Println(tx)
	//res = Find(db, User{}, "name = ?", "Anurag")
	//log.Println(res)
	//tx = Update(db, res, User{Name: "Anurag", Age: 12345})
	//log.Println(tx)
	//res = Find(db, User{}, "name = ?", "Anurag")
	//log.Println(res)
	////db.Delete(res, res.(User).ID)
	////Delete(db, res, res.(User).ID)
	//res = Find(db, User{}, "name = ?", "Anurag")

	err := GetBlockRawModel().Migrate()
	if err != nil {
		log.Println(err)
	}

}

//func Migrate(db *gorm.DB, obj interface{}) error {
//	// Migrate the schema
//	user, ok := obj.(User)
//	if ok {
//		fmt.Printf("Hello %s!\n", user.Name)
//	}
//	err := db.AutoMigrate(user)
//	if err != nil {
//		fmt.Println("err:", err)
//	}
//	return err
//}
//
//func Create(db *gorm.DB, obj interface{}) *gorm.DB {
//	// Create
//	user, ok := obj.(User)
//	if ok {
//		fmt.Printf("Hello %s!\n", user.Name)
//	}
//	tx := db.Create(&user)
//	return tx
//}
//
//func Find(db *gorm.DB, obj interface{}, conds ...interface{}) interface{} {
//	// Read
//	//var user User
//	//db.Find(&user, "name = ?", "Anurag")
//	user, ok := obj.(User)
//	if ok {
//		fmt.Printf("Hello %s!\n", user.Name)
//	}
//	db.Find(&user, conds...)
//	//db.Find(&user, "name = ?", "Anurag")
//	fmt.Println(user)
//	return user
//}
//
//func Update(db *gorm.DB, obj interface{}, updateValues interface{}) *gorm.DB {
//	// Update
//	//db.Model(&user).Update("age", user.Age-1)
//	//db.Model(&user).Updates(User{Name: "Anu", Age: 25}) // non-zero fields
//	if reflect.TypeOf(obj) != reflect.TypeOf(updateValues) {
//		return nil
//	}
//	//db.Model(&user).Updates(map[string]interface{}{"Name": "Anu", "Age": 25})
//	//db.Find(&user, "name = ?", "Anurag")
//	tx := db.Model(obj).Updates(updateValues)
//	fmt.Println(tx)
//	return tx
//}
//
//func Delete(db *gorm.DB, obj interface{}, conds interface{}) *gorm.DB {
//	// Delete
//	tx := db.Delete(obj, conds)
//	return tx
//}
