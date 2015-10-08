package config

import ()

func (this *Config) GetScreen(id int32) *ScreenConfig {
	if v, ok := this.screens[id]; ok {
		return v
	}
	return nil
}

func (this *Config) GetMonster(id int32) *MonsterConfig {
	if v, ok := this.monsters[id]; ok {
		return v
	}
	return nil
}
