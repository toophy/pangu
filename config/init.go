package config

import ()

// 设计目标 :
// 			1. 所有游戏配置都在这里加载
// 			2. 对应同类型的表格文件, 使用通用加载方式, 使用reflect, 结构后面的flag(csv标题名), 进行关联并自动导入二维表
// 			3. 增加一个专用的配置总表, 包含着结构名对应的表格文件名, 因为是静态表, 而且基于程序, 所以bin增加一个config目录, 用于保存 (main.cfg)
//			4. 使用reflect可以做到配置完main.cfg和增加一个表元素对应的结构(flag是标题)以及一个 map表, 外加get文件, 可以顺利使用配置, load函数做到通用(reflect技术)

type Config struct {
	screens  map[int32]*ScreenConfig
	monsters map[int32]*MonsterConfig
}

var GConfig Config

func (this *Config) Init() {
	this.screens = make(map[int32]*ScreenConfig)
	this.monsters = make(map[int32]*MonsterConfig)
}

func init() {
	GConfig.Init()

	// 加载场景配置
	{
		pp := make([]ScreenConfig, 0)
		LoadCsv2Struct("./data/config/screen_list.txt", &pp)
		for i, _ := range pp {
			GConfig.screens[pp[i].Id] = &pp[i]
		}
	}

	// 加载怪物配置
	{
		pp := make([]MonsterConfig, 0)
		LoadCsv2Struct("./data/config/monster_list.txt", &pp)
		for i, _ := range pp {
			GConfig.monsters[pp[i].Id] = &pp[i]
		}
	}

}
