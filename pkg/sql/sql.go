package sql

import (
	"errors"
	"fmt"
	"github.com/little-bit-shy/go-xgz/pkg/database/sql"
	"reflect"
	"strings"
)

type Conditions map[string]Condition
type Condition []interface{}
type Fields map[string]interface{}
type DefaultFields map[string]string

// BuildConditions build conditions
func BuildConditions(cc ...Condition) Conditions {
	c := Conditions{}
	for _, v := range cc {
		fields := v[0].(string)
		symbol := v[1]
		value := v[2]
		withOr := v[3]
		c[fields] = Condition{symbol, value, withOr}
	}
	return c
}

// Build build where sql
func Build(query string, args []interface{}, data Conditions, condition ...string) (where string, newArgs []interface{}) {
	for k, v := range data {
		withOr := ""
		if len(v) >= 3 && len(where) > 0 {
			withOr = v[2].(string)
		}
		switch v[0] {
		case "=", "<=>", "<>", "!=", "<=", ">=", ">", "<", "IS NULL", "IS NOT NULL":
			where = fmt.Sprintf("%v %v %v %v ?", where, withOr, k, v[0])
			args = append(args, v[1])
			break
		case "in":
			var ids []string
			switch v[1].(type) {
			case string:
				ids = strings.Split(v[1].(string), ",")
				break
			case []string:
				ids = v[1].([]string)
				break
			}
			bind := ""
			for _, _ = range ids {
				if bind == "" {
					bind = "?"
				} else {
					bind = fmt.Sprintf("%v,%v", bind, "?")
				}
			}

			where = fmt.Sprintf("%v %v %v %v(%v)", where, withOr, k, v[0], bind)
			for _, vv := range ids {
				args = append(args, vv)
			}
			break
		}
	}
	where = " (" + strings.Trim(where, " ") + ")"
	if len(condition) > 0 {
		where = condition[0] + " (" + strings.Trim(where, " ") + ")"
	}
	where = strings.Trim(query, " ") + " " + where
	newArgs = args
	return
}

// BuildRecord build rows
func BuildRecord(columns []string, rows *sql.Rows) (record []map[string]interface{}) {
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}
	record = make([]map[string]interface{}, 0)
	index := 0
	for rows.Next() {
		//将行数据保存到record字典
		rows.Scan(scanArgs...)
		recordIndex := make(map[string]interface{})
		for i, col := range values {
			recordIndex[columns[i]] = col
		}
		record = append(record, recordIndex)
		index++
	}
	return
}

// BuildGU build create or update
func BuildGU(d interface{}, how string, fileds Fields, defaluts DefaultFields) (
	keys string, binds string, values []interface{}, err error) {
	rt := reflect.TypeOf(d)
	rv := reflect.ValueOf(d)
	out := map[string]interface{}{}
	data := map[string]interface{}{}
	for i := 0; i < rt.NumField(); i++ {
		out[rt.Field(i).Name] = rv.Field(i).Interface()
	}

	for k, v := range out {
		if f, ok := fileds[k]; ok {
			switch f.(type) {
			case string:
				data[f.(string)] = v
				break
			case func() (string, interface{}):
				var filed string
				filed, v = f.(func() (string, interface{}))()
				data[filed] = v
				break
			default:
				err = errors.New("the filed type have error")
				return
			}
		}
	}

	for k, v := range defaluts {
		data[k] = v
	}

	switch how {
	case "create":
		for k, v := range data {
			if len(keys) == 0 {
				keys = fmt.Sprintf("%v", k)
			} else {
				keys = fmt.Sprintf("%v,%v", keys, k)
			}
			if len(binds) == 0 {
				binds = fmt.Sprintf("%v", "?")
			} else {
				binds = fmt.Sprintf("%v,%v", binds, "?")
			}
			values = append(values, v)
		}
		break
	case "update":
		for k, v := range data {
			if len(binds) == 0 {
				binds = fmt.Sprintf("%v=%v", k, "?")
			} else {
				binds = fmt.Sprintf("%v,%v=%v", binds, k, "?")
			}
			values = append(values, v)
		}
		break
	}

	return
}
