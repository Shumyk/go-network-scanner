package db

import (
	"fmt"
	"go-network-scanner/api"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestScanAll(t *testing.T) {
	db := open(memoryDB)
	defer db.Close()
	initSchema(db)

	r := NewScanRepo(db, &api.Config{Name: "tst"})
	r.Upsert(createScan(1))
	r.Upsert(createScan(2))
	r.Upsert(createScan(3))

	scans, err := r.All()
	assert.Nil(t, err)
	assert.Len(t, scans, 3)

	s1 := scans[0]
	s2 := scans[1]
	s3 := scans[2]
	assert.Equal(t, "ip1", s1.IP)
	assert.Equal(t, "ip2", s2.IP)
	assert.Equal(t, "ip3", s3.IP)
	assert.Equal(t, uint16(1), s1.Port)
	assert.Equal(t, uint16(2), s2.Port)
	assert.Equal(t, uint16(3), s3.Port)
	assert.Equal(t, "tst", s1.Session)
	assert.Equal(t, "tst", s2.Session)
	assert.Equal(t, "tst", s3.Session)
	assert.Equal(t, "banner1", s1.ServiceBanner)
	assert.Equal(t, "banner2", s2.ServiceBanner)
	assert.Equal(t, "banner3", s3.ServiceBanner)
	// FirstDiscovered and LastSeen are set at creation, not from supplied values
	assert.NotEqual(t, mockTime, s1.FirstDiscovered)
	assert.NotEqual(t, mockTime, s2.FirstDiscovered)
	assert.NotEqual(t, mockTime, s3.FirstDiscovered)
	assert.NotEqual(t, mockTime, s1.LastSeen)
	assert.NotEqual(t, mockTime, s2.LastSeen)
	assert.NotEqual(t, mockTime, s3.LastSeen)
	// at creation, FirstDiscovered and LastSeen are the same
	assert.Equal(t, s1.FirstDiscovered, s1.LastSeen)
	assert.Equal(t, s2.FirstDiscovered, s2.LastSeen)
	assert.Equal(t, s3.FirstDiscovered, s3.LastSeen)

	time.Sleep(time.Second)
	s1ToUpdate := createScan(1)
	s1ToUpdate.ServiceBanner = "banner1-updated"
	id, err := r.Upsert(s1ToUpdate)
	assert.Nil(t, err)
	assert.Equal(t, 1, id)

	scans, err = r.All()
	assert.Nil(t, err)
	assert.Len(t, scans, 3)

	s1Updated := scans[0]
	assert.Equal(t, "banner1-updated", s1Updated.ServiceBanner)
	assert.Equal(t, s1.FirstDiscovered, s1Updated.FirstDiscovered)
	assert.NotEqual(t, s1.LastSeen, s1Updated.LastSeen)
}

func createScan(id int) *api.Scan {
	return &api.Scan{
		IP:              "ip" + fmt.Sprint(id),
		Port:            uint16(id),
		Session:         "session" + fmt.Sprint(id),
		ServiceBanner:   "banner" + fmt.Sprint(id),
		FirstDiscovered: mockTime,
		LastSeen:        mockTime,
	}
}
