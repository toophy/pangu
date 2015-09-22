package config

import (
	"encoding/csv"
	"os"
	"strconv"
)

// 设计目标 :
// 			1. 所有游戏配置都在这里加载
// 			2. 对应同类型的表格文件, 使用通用加载方式, 使用reflect, 结构后面的flag(csv标题名), 进行关联并自动导入二维表
// 			3. 增加一个专用的配置总表, 包含着结构名对应的表格文件名, 因为是静态表, 而且基于程序, 所以bin增加一个config目录, 用于保存 (main.cfg)
//			4. 使用reflect可以做到配置完main.cfg和增加一个表元素对应的结构(flag是标题)以及一个 map表, 外加get文件, 可以顺利使用配置, load函数做到通用(reflect技术)

type ScreenConfig struct {
	Name    string
	ModName string
	Width   int32
	Height  int32
}

type ScreensConfig struct {
	config map[int32]*ScreenConfig
}

var GScreens ScreensConfig

func init() {
	GScreens.LoadScreenConfig("./data/screens/screen_list.txt")
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
		this.config = make(map[int32]*ScreenConfig)
		for i, _ := range d {
			if i > 0 {
				id, _ := strconv.Atoi(d[i][0])
				w, _ := strconv.Atoi(d[i][3])
				h, _ := strconv.Atoi(d[i][4])
				this.config[int32(id)] = &ScreenConfig{Name: d[i][1], ModName: d[i][2], Width: int32(w), Height: int32(h)}
			}
		}
	}
	return true
}

func (this *ScreensConfig) GetScreenConfig(id int32) *ScreenConfig {
	if v, ok := this.config[id]; ok {
		return v
	}
	return nil
}
