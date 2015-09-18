package thread

import (
	lua "github.com/toophy/gopher-lua"
	"github.com/toophy/pangu/help"
)

// 事件 : 场景增/删
type Event_open_screen struct {
	help.Evt_base
	Screen_oid_    int32
	Screen_name_   string
	Screen_thread_ *ScreenThread
	Open           bool
}

// 事件执行
func (this *Event_open_screen) Exec(home interface{}) bool {
	if this.Open {
		if this.Screen_thread_.Add_screen(this.Screen_name_, this.Screen_oid_) {
			this.Screen_thread_.LogError("打开场景成功")
			return true
		}
		this.Screen_thread_.LogError("打开场景失败")
		return true
	}

	if this.Screen_thread_.Del_screen(this.Screen_oid_) {
		this.Screen_thread_.LogError("关闭场景成功")
		return true
	}
	this.Screen_thread_.LogError("关闭场景失败")
	return true
}

// 事件 : lua使用的通用事件
type Event_from_lua struct {
	help.Evt_base
	module   string     // lua模块名
	function string     // lua函数名
	param    lua.LValue // 参数(table)
}

// 事件执行
func (this *Event_from_lua) Exec(home interface{}) bool {
	// 当前线程调用-> 执行这个事件
	switch home.(type) {
	case *ScreenThread:
		home.(*ScreenThread).Tolua_Common_Param(this.module, this.function, this.param)
	case *WorldThread:
		home.(*WorldThread).Tolua_Common_Param(this.module, this.function, this.param)
	}
	return true
}
