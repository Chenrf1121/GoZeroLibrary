// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	booksFieldNames          = builder.RawFieldNames(&Books{})
	booksRows                = strings.Join(booksFieldNames, ",")
	booksRowsExpectAutoSet   = strings.Join(stringx.Remove(booksFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	booksRowsWithPlaceHolder = strings.Join(stringx.Remove(booksFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheBooksIdPrefix   = "cache:books:id:"
	cacheBooksNamePrefix = "cache:books:name:"
)

type (
	booksModel interface {
		Insert(ctx context.Context, data *Books) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Books, error)
		FindOneByName(ctx context.Context, name string) (*Books, error)
		Update(ctx context.Context, data *Books) error
		Delete(ctx context.Context, id int64) error
	}

	defaultBooksModel struct {
		sqlc.CachedConn
		table string
	}

	Books struct {
		Id       int64  `db:"id"`
		Count    int64  `db:"count"`     // 书本数量
		CountNow int64  `db:"count_now"` // 当前书本数量
		Name     string `db:"name"`      // 书本名称
	}
)

func newBooksModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultBooksModel {
	return &defaultBooksModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`books`",
	}
}

func (m *defaultBooksModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	booksIdKey := fmt.Sprintf("%s%v", cacheBooksIdPrefix, id)
	booksNameKey := fmt.Sprintf("%s%v", cacheBooksNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, booksIdKey, booksNameKey)
	return err
}

func (m *defaultBooksModel) FindOne(ctx context.Context, id int64) (*Books, error) {
	booksIdKey := fmt.Sprintf("%s%v", cacheBooksIdPrefix, id)
	var resp Books
	err := m.QueryRowCtx(ctx, &resp, booksIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", booksRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBooksModel) FindOneByName(ctx context.Context, name string) (*Books, error) {
	booksNameKey := fmt.Sprintf("%s%v", cacheBooksNamePrefix, name)
	var resp Books
	err := m.QueryRowIndexCtx(ctx, &resp, booksNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", booksRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, name); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBooksModel) Insert(ctx context.Context, data *Books) (sql.Result, error) {
	booksIdKey := fmt.Sprintf("%s%v", cacheBooksIdPrefix, data.Id)
	booksNameKey := fmt.Sprintf("%s%v", cacheBooksNamePrefix, data.Name)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, booksRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Count, data.CountNow, data.Name)
	}, booksIdKey, booksNameKey)
	return ret, err
}

func (m *defaultBooksModel) Update(ctx context.Context, newData *Books) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	booksIdKey := fmt.Sprintf("%s%v", cacheBooksIdPrefix, data.Id)
	booksNameKey := fmt.Sprintf("%s%v", cacheBooksNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, booksRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Count, newData.CountNow, newData.Name, newData.Id)
	}, booksIdKey, booksNameKey)
	return err
}

func (m *defaultBooksModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBooksIdPrefix, primary)
}

func (m *defaultBooksModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", booksRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultBooksModel) tableName() string {
	return m.table
}
