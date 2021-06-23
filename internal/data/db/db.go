package db

import (
	"context"
	sql2 "database/sql"
	"github.com/little-bit-shy/go-xgz/internal/dao"
	"github.com/little-bit-shy/go-xgz/pkg/dao/db"
	sql3 "github.com/little-bit-shy/go-xgz/pkg/database/sql"
	"github.com/little-bit-shy/go-xgz/pkg/sql"
)

type Book struct {
	Id   int64  `mapstructure:"id"`
	Name string `mapstructure:"name"`
}

// GetBook get book data
func GetBook(d *dao.Dao, tx *sql3.Tx, ctx context.Context) (result *Book, metadata *db.Metatada, err error) {
	var (
		args []interface{}
	)

	var querySql string
	querySql, args = sql.Build(
		"select * from book where ",
		args,
		sql.BuildConditions(sql.Condition{"id", "=", 1, "AND"}),
	)
	var results db.Results
	var rows *sql3.Rows
	if results, rows, err = d.Db.QueryTx(ctx, tx, querySql, args...); err != nil {
		return
	}
	result = new(Book)
	if metadata, err = d.Db.One(results, rows, &result); err != nil {
		return
	}
	return
}

// UpdateBook update book data
func UpdateBook(d *dao.Dao, tx *sql3.Tx, ctx context.Context, name string) (result sql2.Result, err error) {
	var (
		args []interface{}
	)

	var updateSql string
	updateSql, args = sql.Build(
		"update book set name=? where ",
		args,
		sql.BuildConditions(sql.Condition{"id", "=", 1, "AND"}),
	)
	args = append([]interface{}{name}, args...)
	result, err = d.Db.ExecTx(ctx, tx, updateSql, args...)
	return
}

// CreateBook create book data
func CreateBook(d *dao.Dao, tx *sql3.Tx, ctx context.Context, name string) (result sql2.Result, err error) {
	result, err = d.Db.ExecTx(ctx, tx, "insert book(id,name) value(?,?)", 1, name)
	return
}
