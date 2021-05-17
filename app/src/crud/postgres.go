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
var postgresConnOnce sync.Once

func NewPostgresConn(dsn string) (*postgresConn, error) { // Only for testing
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
	postgresConnOnce.Do(func() {
		// TODO: create dsn string from env variables
		dsn := NewDsn("postgres", "5432", "postgres", "changeme", "postgres", "disable", "UTC")
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

func NewDsn(host string, port string, user string, password string, dbname string, sslmode string, timezone string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone)
}
