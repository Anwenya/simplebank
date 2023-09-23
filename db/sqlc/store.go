package db

import (
	"context"
	"database/sql"
	"fmt"
)

// 执行数据查询并处理事务
type Store struct {
	*Queries
	db *sql.DB
}

// 创建一个Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// 在一个事务中执行查询语句
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	// 开启事务
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// 通过回调执行具体操作
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	// 整体操作没有问题再提交事务
	return tx.Commit()
}

// 交易参数
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// 交易结果
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// 账户之间的转账
// 创建交易记录 添加账户 更新账户余额
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// 1.创建交易记录
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}
		// 2.记录账户余额变动
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// 3.更新账户余额
		result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			Amount: -arg.Amount,
			ID:     arg.FromAccountID,
		})
		if err != nil {
			return err
		}

		result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			Amount: arg.Amount,
			ID:     arg.ToAccountID,
		})

		return nil
	})
	return result, err
}
