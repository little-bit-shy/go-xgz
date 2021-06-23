package language

import (
	"errors"
	"flag"
	"github.com/go-kratos/kratos/pkg/ecode"
	"github.com/little-bit-shy/go-xgz/pkg/helper"
	"io/ioutil"
	"strings"
)

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "language", "", "default language path")
}

// raws2Map raws to map
func raws2Map(rawText string) (raws map[int]string) {
	raws = map[int]string{}
	row := strings.Split(rawText, "\n")
	for _, v := range row {
		slice := strings.Split(v, "=")
		if len(slice) == 2 {
			key := helper.GetInt(strings.Trim(strings.Trim(slice[0], " "), "\""))
			value := strings.Trim(strings.Trim(slice[1], " "), "\"")
			raws[key] = value
		}
	}
	return
}

// RegisterCode register code
func RegisterCode() {
	if confPath == "" {
		helper.Panic(errors.New("the language path is a must"))
	}
	var codeBytes []byte
	var err error
	if codeBytes, err = ioutil.ReadFile(confPath + "/code"); err != nil {
		helper.Panic("code language file not exits")
	}
	helper.Panic(err)
	cms := raws2Map(string(codeBytes))
	ecode.Register(cms)
}
