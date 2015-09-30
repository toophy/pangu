package exm

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"
)

func TestAll(t *testing.T) {
	// ReadCsv(t)
	// ReflectSome(t)
	RelectInterface(t)
}

//
type Humen struct {
	Id    int    `csv:"编号"`
	Name  string `csv:"名称"`
	Class string `csv:"种类"`
	Count int    `csv:"数量"`
}

// 测试interface
func RelectInterface(t *testing.T) {
	y := make([]Humen, 3)
	CallInterface(y)

}

func CallInterface(v interface{}) {

	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Slice {
		vx := reflect.ValueOf(v)
		l := vx.Len()
		for i := 0; i < l; i++ {
			vs := vx.Index(i).Pointer()
			vx.Index(i).

			// ts := reflect.TypeOf(vs)
			// rl := ts.NumField()
			// for k := 0; k < rl; k++ {
			// 	x := ts.Field(k)
			// 	fmt.Printf("%-v\n", x)
			// }

			if reflect.ValueOf(vs).Kind() == reflect.Ptr {
				fmt.Println("ptr")
				reflect.ValueOf(vs).Elem().FieldByName("Name").SetString("gogo")
			} else {
				fmt.Println(reflect.ValueOf(vs).Kind())
			}

			//fmt.Println(reflect.ValueOf(vs).Elem().CanSet())
		}
	} else {
		rl := t.NumField()
		for i := 0; i < rl; i++ {
			x := t.Field(i)
			fmt.Printf("%-v\n", x)
		}
	}

}

// reflect
func ReflectSome(t *testing.T) {

	var y Humen
	tx := reflect.TypeOf(y)
	vx := make([]Humen, 100)

	csv_len := ReadMisc("csv.txt", tx, vx)
	fmt.Println(csv_len)
	fmt.Printf("%-v", vx)
}

func ReadMisc(file_name string, t reflect.Type, v interface{}) (count int) {

	if reflect.TypeOf(v).Kind() != reflect.Slice {
		return
	}

	count = 0

	f, err := os.Open(file_name)
	if err != nil {
		println(err.Error())
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = '\t'

	d, err2 := r.ReadAll()
	if err2 != nil {
		println(err2.Error())
		return count
	}

	rl := t.NumField()

	vx := reflect.ValueOf(v)
	fmt.Println(vx.Len())

	for row, _ := range d {
		// 跳过第一行
		if row == 0 {
			continue
		}
		// 从第二行开始这样处理
		for col, _ := range d[row] {

			for x := 0; x < rl; x++ {
				p := t.Field(x)
				if p.Tag.Get("csv") == d[0][col] {
					switch p.Type.Name() {
					case "string":
						vs := vx.Index(count).Interface()

						ts := reflect.ValueOf(vs)
						println(ts.FieldByName(p.Name).Type().String())

						reflect.ValueOf(vs).Elem().FieldByName(p.Name).SetString(d[row][col])
					case "int":
						vs := vx.Index(count).Interface()
						n, _ := strconv.Atoi(d[row][col])
						reflect.ValueOf(vs).Elem().FieldByName(p.Name).SetInt(int64(n))
					}
					break
				}
			}
		}
		// 处理完成一行
		count++
	}

	return count
}

// 读csv文件
func ReadCsv(t *testing.T) {
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
