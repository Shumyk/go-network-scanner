package db

import (
	"fmt"
	"go-network-scanner/api"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigAll(t *testing.T) {
	db := open(memoryDB)
	defer db.Close()
	initSchema(db)

	r := NewConfigRepo(db, &api.Config{Name: "tst"})
	r.Upsert(createConfig(1))
	r.Upsert(createConfig(2))
	r.Upsert(createConfig(3))

	configs, err := r.All()
	assert.Nil(t, err)
	assert.Len(t, configs, 3)

	c1 := configs[0]
	c2 := configs[1]
	c3 := configs[2]
	assert.Equal(t, "name1", c1.Name)
	assert.Equal(t, "name2", c2.Name)
	assert.Equal(t, "name3", c3.Name)
	assert.Equal(t, false, c1.EnableProxy)
	assert.Equal(t, true, c2.EnableProxy)
	assert.Equal(t, false, c3.EnableProxy)
	assert.Equal(t, "ip1-1", c1.IPRange1)
	assert.Equal(t, "ip2-1", c1.IPRange2)
	assert.Equal(t, "ip3-1", c1.IPRange3)
	assert.Equal(t, "ip1-2", c2.IPRange1)
	assert.Equal(t, "ip2-2", c2.IPRange2)
	assert.Equal(t, "ip3-2", c2.IPRange3)
	assert.Equal(t, "ip1-3", c3.IPRange1)
	assert.Equal(t, "ip2-3", c3.IPRange2)
	assert.Equal(t, "ip3-3", c3.IPRange3)
	// CreatedAt is set at creation, not from supplied values
	assert.NotEqual(t, mockTime, c1.CreatedAt)
	assert.NotEqual(t, mockTime, c2.CreatedAt)
	assert.NotEqual(t, mockTime, c3.CreatedAt)

	configByName, err := r.Get("name2")
	assert.Nil(t, err)
	assert.Equal(t, c2, configByName)
}

func createConfig(i int) *api.Config {
	return &api.Config{
		Name:        fmt.Sprintf("name%d", i),
		EnableProxy: i%2 == 0,
		IPRange1:    fmt.Sprintf("ip1-%d", i),
		IPRange2:    fmt.Sprintf("ip2-%d", i),
		IPRange3:    fmt.Sprintf("ip3-%d", i),
		CreatedAt:   mockTime,
	}
}
