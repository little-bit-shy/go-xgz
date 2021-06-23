package config

import (
	"encoding/json"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/little-bit-shy/go-xgz/pkg/config/toml"
	"github.com/little-bit-shy/go-xgz/pkg/helper"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	PREFIX = "env:"
)

// GetConfig get config with file or env
func Env(ct *paladin.TOML, cf interface{}, key string) (err error) {
	var rawText string
	rawText, err = ct.Get(key).Raw()
	helper.Panic(err)
	_, _ = toml.Decode(rawText, cf)
	raws := Raws2Map(rawText)
	// don't handle this error
	tR := reflect.TypeOf(cf).Elem()
	vR := reflect.ValueOf(cf).Elem()
	for i := 0; i < tR.NumField(); i++ {
		fT := tR.Field(i)
		vT := vR.FieldByName(fT.Name)
		raw := raws[fT.Name]
		if !strings.HasPrefix(raw, PREFIX) {
			continue
		}
		env := os.Getenv(strings.Replace(raw, PREFIX, "", 1))
		if env == "" {
			continue
		}
		switch vT.Type().Name() {
		case "bool":
			var v bool
			v, err = strconv.ParseBool(env)
			helper.Panic(err)
			vT.SetBool(v)
			break
		case "int", "int32", "int64":
			var v int
			v, err = strconv.Atoi(env)
			helper.Panic(err)
			vT.SetInt(int64(v))
			break
		case "float", "float32", "float64":
			var v float64
			v, err = strconv.ParseFloat(env, 64)
			helper.Panic(err)
			vT.SetFloat(v)
			break
		case "string":
			vT.SetString(env)
			break
		case "Duration":
			var v time.Duration
			v, err = time.ParseDuration(env)
			helper.Panic(err)
			vT.SetInt(v.Nanoseconds())
			break
		default:
			sliceData := []string{}
			err = json.Unmarshal([]byte(env), &sliceData)
			helper.Panic(err)
			vT.Set(reflect.ValueOf(sliceData))
			break
		}
	}
	return
}

// Raws2Map raws to map
func Raws2Map(rawText string) (raws map[string]string) {
	raws = map[string]string{}
	row := strings.Split(rawText, "\n")
	for _, v := range row {
		slice := strings.Split(v, "=")
		if len(slice) == 2 {
			raws[strings.Trim(strings.Trim(slice[0], " "), "\"")] =
				strings.Trim(strings.Trim(slice[1], " "), "\"")
		}
	}
	return
}
