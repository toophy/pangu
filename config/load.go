package config

import (
	"github.com/toophy/csv"
	"io/ioutil"
	"os"
)

func LoadCsv2Struct(name string, data interface{}) bool {

	f, err := os.Open(name)
	if err != nil {
		println(err.Error())
		return false
	}
	defer f.Close()

	d, err2 := ioutil.ReadAll(f)
	if err2 != nil {
		println(err2.Error())
	} else {
		csv.Unmarshal('\t', true, d, data)
	}

	return true
}
