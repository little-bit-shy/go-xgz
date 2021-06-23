package helper

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
)

// Panic throw panic
func Panic(err interface{}) {
	if err != nil {
		panic(err)
	}
}

// O2md5 md5
func O2md5(str string) (md5str string) {
	h := md5.New()
	h.Write([]byte(str))
	md5str = hex.EncodeToString(h.Sum(nil))
	return
}

// Get64 get int64
func GetInt64(i interface{}) (value int64) {
	if v, ok := i.(string); ok {
		// 去除小数点后
		nonFractionalPart := strings.Split(v, ".")
		atoi, _ := strconv.Atoi(nonFractionalPart[0])
		value = int64(atoi)
	} else if v, ok := i.(float32); ok {
		value = int64(v)
	} else if v, ok := i.(float64); ok {
		value = int64(v)
	} else if v, ok := i.(int); ok {
		value = int64(v)
	} else if v, ok := i.(int32); ok {
		value = int64(v)
	} else if v, ok := i.(int64); ok {
		value = v
	} else if v, ok := i.(decimal.Decimal); ok {
		value = v.IntPart()
	}
	return
}

// Get32 get32
func GetInt32(i interface{}) (value int32) {
	if v, ok := i.(string); ok {
		// 去除小数点后
		nonFractionalPart := strings.Split(v, ".")
		atoi, _ := strconv.Atoi(nonFractionalPart[0])
		value = int32(atoi)
	} else if v, ok := i.(float32); ok {
		value = int32(v)
	} else if v, ok := i.(float64); ok {
		value = int32(v)
	} else if v, ok := i.(int); ok {
		value = int32(v)
	} else if v, ok := i.(int64); ok {
		value = int32(v)
	} else if v, ok := i.(int32); ok {
		value = v
	} else if v, ok := i.(decimal.Decimal); ok {
		value = int32(v.IntPart())
	}
	return
}

// Get32 get32
func GetInt(i interface{}) (value int) {
	if v, ok := i.(string); ok {
		// 去除小数点后
		nonFractionalPart := strings.Split(v, ".")
		atoi, _ := strconv.Atoi(nonFractionalPart[0])
		value = int(atoi)
	} else if v, ok := i.(float32); ok {
		value = int(v)
	} else if v, ok := i.(float64); ok {
		value = int(v)
	} else if v, ok := i.(int); ok {
		value = v
	} else if v, ok := i.(int64); ok {
		value = int(v)
	} else if v, ok := i.(int32); ok {
		value = int(v)
	} else if v, ok := i.(decimal.Decimal); ok {
		value = int(v.IntPart())
	}
	return
}

// LimitOffset db limit offset
func LimitOffset(page int64, pageCount int64) (limit int64, offset int64) {
	limitMax := int64(50)
	limit = pageCount
	offset = 0
	if page <= 1 {
		offset = 0
	}
	if page > 1 {
		offset = (page - 1) * pageCount
	}
	if limit >= limitMax {
		limit = limitMax
	}
	return
}
