package etransaction

import "context"

// Transaction 定义通用的事务接口，能够支持本地事务或分布式事务
type Transaction interface {
	ExecTx(ctx context.Context, fn func(ctx context.Context) error) error
}
