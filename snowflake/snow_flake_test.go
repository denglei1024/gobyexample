package snowflake

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Snowflake 结构表示雪花算法的参数
type Snowflake struct {
	mu          sync.Mutex
	startTime   int64
	machineID   int64
	sequence    int64
	lastStamp   int64
	sequenceBit int64
	machineBit  int64
	timeBit     int64
}

// 初始化雪花算法
func NewSnowflake(startTime int64, machineID int64, sequenceBit int64, machineBit int64, timeBit int64) *Snowflake {
	return &Snowflake{
		startTime:   startTime,
		machineID:   machineID,
		sequenceBit: sequenceBit,
		machineBit:  machineBit,
		timeBit:     timeBit,
		sequence:    0,
		lastStamp:   -1,
	}
}

// GenerateID 生成唯一ID
func (sf *Snowflake) GenerateID() int64 {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	currentStamp := time.Now().UnixNano() / 1e6

	if currentStamp == sf.lastStamp {
		sf.sequence = (sf.sequence + 1) & ((1 << sf.sequenceBit) - 1)
		if sf.sequence == 0 {
			for currentStamp <= sf.lastStamp {
				currentStamp = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		sf.sequence = 0
	}

	if currentStamp < sf.lastStamp {
		fmt.Println("lock moved backwards. Refusing to generate id for", sf.lastStamp-currentStamp, "milliseconds")
		return 0
	}
	sf.lastStamp = currentStamp

	ID := ((currentStamp - sf.startTime) << (sf.machineBit + sf.sequenceBit)) |
		(sf.machineID << sf.sequenceBit) |
		sf.sequence
	return ID
}

func TestGenerateID(t *testing.T) {
	startTime := int64(1609459200000) // 2021-01-01 00:00:00
	machineID := int64(1)
	sequenceBit := int64(12)
	machineBit := int64(5)
	timeBit := int64(41)

	sf := NewSnowflake(startTime, machineID, sequenceBit, machineBit, timeBit)

	// 生成10个唯一ID
	for i := 0; i < 1000; i++ {
		id := sf.GenerateID()
		t.Logf("Generating ID %d", id)
	}
}
