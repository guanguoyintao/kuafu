package etransaction

import (
	"context"
	"github.com/dtm-labs/dtmgrpc"
)

// DTMTransaction 分布式事务的实现
type DTMTransaction struct {
	serverAddress string // dtm 服务的地址
}

// NewDTMTransaction 创建分布式事务实例
func NewDTMTransaction(serverAddress string) Transaction {
	return &DTMTransaction{serverAddress: serverAddress}
}

// ExecTx 使用 DTM 执行分布式事务
func (t *DTMTransaction) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	gid := dtmgrpc.MustGenGid(t.serverAddress) // 生成全局事务ID
	// 直接调用 dtmgrpc.TccGlobalTransaction 而不需要 dtmgrpc.Client
	return dtmgrpc.TccGlobalTransaction(t.serverAddress, gid, func(tcc *dtmgrpc.TccGrpc) error {
		txCtx := context.WithValue(ctx, "tcc", tcc) // 将 TCC 事务放入上下文
		return fn(txCtx)
	})
}
