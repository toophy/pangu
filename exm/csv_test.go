package exm

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"
)

func TestReadCsv(t *testing.T) {
	fmt.Println("open csv.txt")
	f, err := os.Open("csv.txt")
	if err != nil {
		println(err.Error())
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = '\t'

	d, err2 := r.ReadAll()
	if err2 != nil {
		println(err2.Error())
	} else {
		fmt.Println(d)
	}
}
