package db

import (
	"context"
	"fmt"
)

// 在一个事务中执行查询语句
func (store *SqlStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	// 开启事务
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return err
	}
	// 通过回调执行具体操作
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	// 整体操作没有问题再提交事务
	return tx.Commit(ctx)
}
