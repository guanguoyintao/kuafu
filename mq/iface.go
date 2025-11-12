package emq

import (
	"context"
	"encoding/json"
)

// Message 定义通用消息结构
type Message struct {
	Header map[string]string `json:"header"`
	Data   json.RawMessage   `json:"data"`
}

// MQClient 定义通用 MQ 接口
type MQClient interface {
	Publish(ctx context.Context, topic string, eventType []string, msg *Message) error
	Subscribe(ctx context.Context, groupID, topic string, eventType []string, handler func(ctx context.Context, msg *Message) error) error
	Close(context.Context) error
}
