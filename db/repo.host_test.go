package db

import (
	"fmt"
	"go-network-scanner/api"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	mockTime, _ = time.Parse(time.DateOnly, "2021-01-01")
)

func TestAll(t *testing.T) {
	db := open(memoryDB)
	defer db.Close()
	initSchema(db)

	r := NewHostRepo(db, &api.Config{Name: "tst"})
	r.Upsert(createHost(1))
	r.Upsert(createHost(2))
	r.Upsert(createHost(3))

	hosts, err := r.All()
	assert.Nil(t, err)
	assert.Len(t, hosts, 3)
	h1 := hosts[0]
	h2 := hosts[1]
	h3 := hosts[2]
	assert.Equal(t, "ip1", h1.IP)
	assert.Equal(t, "ip2", h2.IP)
	assert.Equal(t, "ip3", h3.IP)
	assert.Equal(t, "os1", h1.OsType)
	assert.Equal(t, "os2", h2.OsType)
	assert.Equal(t, "os3", h3.OsType)
	assert.Equal(t, "hl1", h1.HostnameLabel)
	assert.Equal(t, "hl2", h2.HostnameLabel)
	assert.Equal(t, "hl3", h3.HostnameLabel)
	assert.Equal(t, "cl1", h1.CustomLabel)
	assert.Equal(t, "cl2", h2.CustomLabel)
	assert.Equal(t, "cl3", h3.CustomLabel)
	// FirstDiscovered and LastSeen are set at creation, not from supplied values
	assert.NotEqual(t, mockTime, h1.FirstDiscovered)
	assert.NotEqual(t, mockTime, h2.FirstDiscovered)
	assert.NotEqual(t, mockTime, h3.FirstDiscovered)
	assert.NotEqual(t, mockTime, h1.LastSeen)
	assert.NotEqual(t, mockTime, h2.LastSeen)
	assert.NotEqual(t, mockTime, h3.LastSeen)
	// at creation, FirstDiscovered and LastSeen are the same
	assert.Equal(t, h1.FirstDiscovered, h1.LastSeen)
	assert.Equal(t, h2.FirstDiscovered, h2.LastSeen)
	assert.Equal(t, h3.FirstDiscovered, h3.LastSeen)

	time.Sleep(time.Second)
	h1ToUpdate := createHost(1)
	h1ToUpdate.CustomLabel = "cl1-updated"
	r.Upsert(h1ToUpdate)
	hosts, err = r.All()
	assert.Nil(t, err)
	assert.Len(t, hosts, 3)

	h1Updated := hosts[0]
	assert.Equal(t, "cl1-updated", h1Updated.CustomLabel)
	assert.Equal(t, h1.FirstDiscovered, h1Updated.FirstDiscovered)
	assert.NotEqual(t, h1.LastSeen, h1Updated.LastSeen)
}

func createHost(id int) *api.Host {
	return &api.Host{
		IP:              "ip" + fmt.Sprint(id),
		FirstDiscovered: mockTime,
		LastSeen:        mockTime,
		OsType:          "os" + fmt.Sprint(id),
		HostnameLabel:   "hl" + fmt.Sprint(id),
		CustomLabel:     "cl" + fmt.Sprint(id),
	}
}
