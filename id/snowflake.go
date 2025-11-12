package eid

import (
	"fmt"
	"sync"
	"time"
)

type ID string

const (
	workerBits  uint8 = 10                      // 机器ID的位数
	numberBits  uint8 = 12                      // 序列号的位数
	workerMax   int64 = -1 ^ (-1 << workerBits) // 机器ID的最大值
	numberMax   int64 = -1 ^ (-1 << numberBits) // 序列号的最大值
	timeShift   uint8 = workerBits + numberBits // 时间戳向左的偏移量
	workerShift uint8 = numberBits              // 机器ID向左的偏移量
	epoch       int64 = 1288834974657           // 起始时间戳 (2010-11-04 09:42:54.657)
)

type Snowflake struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64 // 机器ID
	number    int64 // 序列号
}

var (
	sf   *Snowflake
	once sync.Once
)

// NewSnowflake 创建一个雪花ID生成器
// workerId 机器ID, 具体作用是区分不同的机器, 在分布式系统中, 不同的机器需要不同的workerId, 否则会生成相同的ID
func NewSnowflake(workerId int64) (*Snowflake, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, fmt.Errorf("worker ID must be between 0 and %d", workerMax)
	}
	return &Snowflake{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

// GetSnowflake 获取雪花ID生成器
func GetSnowflake() *Snowflake {
	once.Do(func() {
		var err error
		sf, err = NewSnowflake(1) // 默认使用workerId=1
		if err != nil {
			panic(err)
		}
	})
	return sf
}

// Gen 生成一个雪花ID
func (s *Snowflake) Gen() ID {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixMilli()
	if s.timestamp == now {
		s.number++
		if s.number > numberMax {
			for now <= s.timestamp {
				now = time.Now().UnixMilli()
			}
			s.number = 0
		}
	} else {
		s.number = 0
	}
	s.timestamp = now

	id := ((now - epoch) << timeShift) |
		(s.workerId << workerShift) |
		s.number

	return ID(fmt.Sprintf("%d", id))
}
