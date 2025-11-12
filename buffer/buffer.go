package ebuffer

import (
	"bufio"
	"bytes"
	"io"
	"sync"
)

const (
	KB = 1 << 10 // 1KB
	MB = 1 << 20 // 1MB
	GB = 1 << 30 // 1GB
)

// 定义一个全局缓存池管理
var (
	writerPool32KB  = sync.Pool{New: func() interface{} { return bufio.NewWriterSize(nil, 32<<10) }}
	writerPool512KB = sync.Pool{New: func() interface{} { return bufio.NewWriterSize(nil, 512<<10) }}
	writerPool2MB   = sync.Pool{New: func() interface{} { return bufio.NewWriterSize(nil, 2<<20) }}
	writerPool4MB   = sync.Pool{New: func() interface{} { return bufio.NewWriterSize(nil, 4<<20) }}
	writerPool8MB   = sync.Pool{New: func() interface{} { return bufio.NewWriterSize(nil, 8<<20) }}
)

// 全局 buff池
var (
	bufferPool32KB  = &sync.Pool{New: func() interface{} { return bytes.NewBuffer(make([]byte, 0, 32<<10)) }}
	bufferPool512KB = &sync.Pool{New: func() interface{} { return bytes.NewBuffer(make([]byte, 0, 512<<10)) }}
	bufferPool2MB   = &sync.Pool{New: func() interface{} { return bytes.NewBuffer(make([]byte, 0, 2<<20)) }}
	bufferPool4MB   = &sync.Pool{New: func() interface{} { return bytes.NewBuffer(make([]byte, 0, 4<<20)) }}
	bufferPool8MB   = &sync.Pool{New: func() interface{} { return bytes.NewBuffer(make([]byte, 0, 8<<20)) }}
)

// PoolBuffer 是一个基于对象池的内存 buffer
type PoolBuffer struct {
	*bytes.Buffer
	pool *sync.Pool
}

// NewPoolBuffer 创建一个新的 PoolBuffer
func NewPoolBuffer(size int64) *PoolBuffer {
	var pool *sync.Pool
	switch {
	case size <= 5<<20:
		pool = bufferPool32KB
	case size <= 10<<20:
		pool = bufferPool512KB
	case size <= 100<<20:
		pool = bufferPool2MB
	case size <= 5<<30:
		pool = bufferPool4MB
	default:
		pool = bufferPool8MB
	}
	buf := pool.Get().(*bytes.Buffer)
	buf.Reset()
	return &PoolBuffer{Buffer: buf, pool: pool}
}

// Close 将 PoolBuffer 放回对象池
func (p *PoolBuffer) Close() error {
	if p.Buffer != nil {
		p.Reset()
		p.pool.Put(p.Buffer)
		p.Buffer = nil
	}
	return nil
}

func (p *PoolBuffer) Len() int {
	return p.Buffer.Len()
}

// String 返回缓冲区的字符串表示
func (p *PoolBuffer) String() string {
	return p.Buffer.String()
}

// GetBufferIOWriter 根据大小获取 bufio.Writer
func GetBufferIOWriter(w io.Writer, size int64) (*bufio.Writer, func()) {
	var writer *bufio.Writer
	var release func()
	switch {
	case size <= 5<<20:
		writer = writerPool32KB.Get().(*bufio.Writer)
		release = func() {
			writer.Flush()
			writer.Reset(nil)
			writerPool32KB.Put(writer)
		}
	case size <= 10<<20:
		writer = writerPool512KB.Get().(*bufio.Writer)
		release = func() {
			writer.Flush()
			writer.Reset(nil)
			writerPool512KB.Put(writer)
		}
	case size <= 100<<20:
		writer = writerPool2MB.Get().(*bufio.Writer)
		release = func() {
			writer.Flush()
			writer.Reset(nil)
			writerPool2MB.Put(writer)
		}
	case size <= 5<<30:
		writer = writerPool4MB.Get().(*bufio.Writer)
		release = func() {
			writer.Flush()
			writer.Reset(nil)
			writerPool4MB.Put(writer)
		}
	default:
		writer = writerPool8MB.Get().(*bufio.Writer)
		release = func() {
			writer.Flush()
			writer.Reset(nil)
			writerPool8MB.Put(writer)
		}
	}
	writer.Reset(w)
	return writer, release
}
