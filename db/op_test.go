package db

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	memoryDB       = ":memory:"
	createTableSQL = "CREATE TABLE test (id INTEGER PRIMARY KEY)"
)

func TestOpen(t *testing.T) {
	db := open(memoryDB)
	defer db.Close()

	err := db.Ping()
	assert.Nil(t, err)

	initSchema(db)

	checkTableExists(t, db, "config")
	checkTableExists(t, db, "scan")
	checkTableExists(t, db, "host")
}

func TestExec(t *testing.T) {
	db := open(memoryDB)
	defer db.Close()

	err := exec(db, createTableSQL)
	assert.Nil(t, err)
	checkTableExists(t, db, "test")

	err = exec(db, createTableSQL)
	assert.NotNil(t, err)
	checkTableExists(t, db, "test")
}

func checkTableExists(t *testing.T, db *sql.DB, table string) {
	r, err := db.Query(
		"SELECT name FROM sqlite_master WHERE type='table' AND name=?",
		table,
	)
	defer r.Close()

	assert.Nil(t, err)
	assert.True(t, r.Next())

	c, err := r.Columns()
	assert.Nil(t, err)
	assert.Equal(t, "name", c[0])
}
