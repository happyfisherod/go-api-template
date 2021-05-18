package postgres_crud_test

import (
	"github.com/geometry-labs/app/crud/postgres_crud"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Postgres", func() {

	var (
		dsn string

		host     string
		port     string
		user     string
		password string
		dbname   string
		sslmode  string
		timezone string
	)

	BeforeEach(func() {
		dsn = "host=localhost user=postgres password=changeme dbname=test_db port=5432 sslmode=disable TimeZone=UTC"

		host = "localhost"
		port = "5432"
		user = "postgres"
		password = "changeme"
		dbname = "test_db"
		sslmode = "disable"
		timezone = "UTC"
	})

	Describe("Create DSN string", func() {
		Context("new dsn string", func() {
			It("dsn string valid", func() {
				Expect(postgres_crud.NewDsn(host, port, user, password, dbname, sslmode, timezone)).To(Equal(dsn))
			})
		})
	})
})
