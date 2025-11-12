package etransaction

import (
	"context"
	"gorm.io/gorm"
)

// LocalTransaction 本地事务的具体实现
type LocalTransaction struct {
	db *gorm.DB
}

// NewLocalTransaction 创建本地事务实例
func NewLocalTransaction(db *gorm.DB) Transaction {
	return &LocalTransaction{db: db}
}

// ExecTx 使用 GORM 执行本地事务
func (t *LocalTransaction) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 将当前事务注入到上下文中，使后续操作能够访问
		txCtx := context.WithValue(ctx, "tx", tx)
		return fn(txCtx)
	})
}
