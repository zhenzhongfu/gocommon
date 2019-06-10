package uuid

import (
	"hash/crc32"

	"github.com/sony/sonyflake"
)

type UuidGenerator struct {
	*sonyflake.Sonyflake
}

func NewUuidGenerator(tag string) *UuidGenerator {
	var st sonyflake.Settings
	// TODO 从某个时间点算起
	//st.StartTime = time.Now()
	st.MachineID = func() (uint16, error) {
		return uint16(HashSum(tag)), nil
	}
	return &UuidGenerator{
		sonyflake.NewSonyflake(st),
	}
}

func (u *UuidGenerator) NexID() (uint64, error) {
	return u.NextID()
}

func HashSum(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	return 0
}
