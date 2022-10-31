package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BorrowModel = (*customBorrowModel)(nil)

type (
	// BorrowModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBorrowModel.
	BorrowModel interface {
		borrowModel
	}

	customBorrowModel struct {
		*defaultBorrowModel
	}
)

// NewBorrowModel returns a model for the database table.
func NewBorrowModel(conn sqlx.SqlConn, c cache.CacheConf) BorrowModel {
	return &customBorrowModel{
		defaultBorrowModel: newBorrowModel(conn, c),
	}
}
