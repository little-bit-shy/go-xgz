package db

import (
	"context"
	sql2 "database/sql"
	"github.com/little-bit-shy/go-xgz/pkg/database/sql"
	sql3 "github.com/little-bit-shy/go-xgz/pkg/sql"
	"github.com/mitchellh/mapstructure"
	"io"
)

type Results []map[string]interface{}

type Record interface{}

type Records []Record

type Db struct {
	Cfg     *sql.Config
	Connect *sql.DB
	io.Closer
}

type Metatada struct {
	Exist bool
}

// New new db connect
func New(cfg *sql.Config) *Db {
	db := sql.NewMySQL(cfg)
	return &Db{
		Cfg:     cfg,
		Connect: db,
	}
}

// One get one data
func (d *Db) One(record []map[string]interface{}, rows *sql.Rows, r Record) (metadata *Metatada, err error) {
	metadata = &Metatada{}
	if len(record) > 0 {
		metadata.Exist = true
		if err = mapstructure.WeakDecode(record[0], r); err != nil {
			return
		}
	}
	return
}

// All get all data
func (d *Db) All(record []map[string]interface{}, rows *sql.Rows, r Record) (metadata *Metatada, err error) {
	metadata = &Metatada{}
	if len(record) > 0 {
		metadata.Exist = true
		if err = mapstructure.WeakDecode(record, r); err != nil {
			return
		}
	}
	return
}

// Query query sql
func (d *Db) Query(ctx context.Context, query string, args ...interface{}) (record Results, rows *sql.Rows, err error) {
	db := d.Connect
	if rows, err = db.Query(ctx, query, args...); err != nil {
		return
	}
	defer rows.Close()
	columns, _ := rows.Columns()
	// build data
	record = sql3.BuildRecord(columns, rows)
	return
}

// QueryTx query tx sql
func (d *Db) QueryTx(ctx context.Context, tx *sql.Tx, query string, args ...interface{}) (record Results, rows *sql.Rows, err error) {
	db := tx
	if rows, err = db.Query(query, args...); err != nil {
		return
	}
	columns, _ := rows.Columns()
	// build data
	record = sql3.BuildRecord(columns, rows)

	if err = rows.Err(); err != nil {
		return
	}
	return
}

// Exec exec sql
func (d *Db) Exec(ctx context.Context, query string, args ...interface{}) (result sql2.Result, err error) {
	db := d.Connect
	result, err = db.Exec(ctx, query, args...)
	return
}

// ExecTx exec tx sql
func (d *Db) ExecTx(ctx context.Context, tx *sql.Tx, query string, args ...interface{}) (result sql2.Result, err error) {
	db := tx
	result, err = db.Exec(query, args...)
	return
}

func (d *Db) Close() error {
	_ = d.Connect.Close()
	return nil
}
