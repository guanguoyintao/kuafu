package eid

import (
	"testing"
	"time"
)

func TestNewSnowflake(t *testing.T) {
	// 测试正常情况
	sf, err := NewSnowflake(1)
	if err != nil {
		t.Errorf("NewSnowflake failed: %v", err)
	}
	if sf == nil {
		t.Error("NewSnowflake returned nil")
	}

	// 测试workerId超出范围
	_, err = NewSnowflake(-1)
	if err == nil {
		t.Error("Expected error for negative workerId")
	}

	_, err = NewSnowflake(workerMax + 1)
	if err == nil {
		t.Error("Expected error for workerId exceeding max")
	}
}

func TestGetSnowflake(t *testing.T) {
	// 测试单例模式
	sf1 := GetSnowflake()
	sf2 := GetSnowflake()
	t.Log(sf1, sf2)
	if sf1 != sf2 {
		t.Error("GetSnowflake should return the same instance")
	}
}

func TestSnowflake_Gen(t *testing.T) {
	sf := GetSnowflake()

	// 测试生成多个ID
	ids := make(map[ID]bool)
	for i := 0; i < 1000; i++ {
		id := sf.Gen()
		if ids[id] {
			t.Errorf("Duplicate ID generated: %s", id)
		}
		ids[id] = true
	}

	// 测试时间回拨情况
	sf.timestamp = time.Now().UnixMilli() + 1000 // 模拟时间回拨
	id := sf.Gen()
	if id == "" {
		t.Error("Failed to generate ID after time rollback")
	}
}

func TestSnowflake_Concurrent(t *testing.T) {
	sf := GetSnowflake()
	ch := make(chan ID, 1000)

	// 并发测试
	for i := 0; i < 1000; i++ {
		go func() {
			ch <- sf.Gen()
		}()
	}

	// 收集生成的ID
	ids := make(map[ID]bool)
	for i := 0; i < 1000; i++ {
		id := <-ch
		if ids[id] {
			t.Errorf("Duplicate ID generated in concurrent test: %s", id)
		}
		ids[id] = true
	}
}
