package crud

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDsn(t *testing.T) {
	dsn := "host=localhost user=postgres password=changeme dbname=test_db port=5432 sslmode=disable TimeZone=UTC"
	newDsn := NewDsn("localhost", "5432", "postgres", "changeme", "test_db", "disable", "UTC")
	assert.Equal(t, newDsn, dsn)
}
