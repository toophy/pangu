package config

import ()

// 场景配置
type ScreenConfig struct {
	Id      int32  `csv:"编号"`
	Name    string `csv:"场景名称"`
	ModName string `csv:"Lua模块名"`
	Width   int32  `csv:"宽"`
	Height  int32  `csv:"高"`
}

// 怪物配置
type MonsterConfig struct {
	Id    int32  `csv:"编号"`
	Name  string `csv:"名称"`
	Speed int32  `csv:"速度"`
}
