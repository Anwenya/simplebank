package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error)
}

// 执行数据查询并处理事务
type SqlStore struct {
	*Queries
	connPool *pgxpool.Pool
}

// 创建一个Store
func NewStore(connPool *pgxpool.Pool) Store {
	return &SqlStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
