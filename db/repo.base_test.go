package db

import (
	"go-network-scanner/api"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFailed(t *testing.T) {
	db := open(memoryDB)
	defer db.Close()
	// Should fail as tables not created
	r := NewConfigRepo(db, &api.Config{Name: "tst"})
	config, err := r.Get("name")
	assert.Equal(t, api.EmptyConfig(), config)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "SQL logic error: no such table: config")
}

func TestAllFailed(t *testing.T) {
	db := open(memoryDB)
	defer db.Close()
	// Should fail as tables not created
	r := NewHostRepo(db, &api.Config{Name: "tst"})
	hosts, err := r.All()
	assert.Nil(t, hosts)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to retrieve all rows")
}

func TestUpsertFailed(t *testing.T) {
	db := open(memoryDB)
	defer db.Close()
	// Should fail as tables not created
	r := NewHostRepo(db, &api.Config{Name: "tst"})
	rowsAffected, err := r.Upsert(&api.Host{})
	assert.Zero(t, rowsAffected)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to upsert")
}
