package thread

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Name    string
	ModName string
	Width   int32
	Height  int32
}

type ScreensConfig struct {
	config map[int32]*Config
}

var screen_config ScreensConfig

func init() {
	screen_config.LoadScreenConfig("./data/screens/screen_list.txt")
}

func (this *ScreensConfig) LoadScreenConfig(name string) bool {

	f, err := os.Open(name)
	if err != nil {
		println(err.Error())
		return false
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = '\t'
	r.TrimLeadingSpace = true

	d, err2 := r.ReadAll()
	if err2 != nil {
		println(err2.Error())
	} else {
		this.config = make(map[int32]*Config)
		for i, _ := range d {
			if i > 0 {
				id, _ := strconv.Atoi(d[i][0])
				w, _ := strconv.Atoi(d[i][3])
				h, _ := strconv.Atoi(d[i][4])
				this.config[int32(id)] = &Config{Name: d[i][1], ModName: d[i][2], Width: int32(w), Height: int32(h)}
			}
		}
	}
	return true
}

func (this *ScreensConfig) GetScreenConfig(id int32) *Config {
	if v, ok := this.config[id]; ok {
		return v
	}
	return nil
}
