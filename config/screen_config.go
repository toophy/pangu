package config

import (
	"fmt"
	"github.com/toophy/csv"
	"io/ioutil"
	"os"
)

// 设计目标 :
// 			1. 所有游戏配置都在这里加载
// 			2. 对应同类型的表格文件, 使用通用加载方式, 使用reflect, 结构后面的flag(csv标题名), 进行关联并自动导入二维表
// 			3. 增加一个专用的配置总表, 包含着结构名对应的表格文件名, 因为是静态表, 而且基于程序, 所以bin增加一个config目录, 用于保存 (main.cfg)
//			4. 使用reflect可以做到配置完main.cfg和增加一个表元素对应的结构(flag是标题)以及一个 map表, 外加get文件, 可以顺利使用配置, load函数做到通用(reflect技术)

type ScreenConfig struct {
	Id      int32  `csv:"编号"`
	Name    string `csv:"场景名称"`
	ModName string `csv:"Lua模块名"`
	Width   int32  `csv:"宽"`
	Height  int32  `csv:"高"`
}

type ScreensConfig struct {
	config map[int32]*ScreenConfig
}

var GScreens ScreensConfig

func init() {
	GScreens.LoadScreenConfig("./data/config/screen_list.txt")
}

func (this *ScreensConfig) LoadScreenConfig(name string) bool {

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
		this.config = make(map[int32]*ScreenConfig)

		pp := make([]ScreenConfig, 0)
		csv.Unmarshal('\t', true, d, &pp)

		fmt.Println(pp)
	}
	return true
}

func (this *ScreensConfig) GetScreenConfig(id int32) *ScreenConfig {
	if v, ok := this.config[id]; ok {
		return v
	}
	return nil
}
