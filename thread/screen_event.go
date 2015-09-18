package thread

import (
	_ "bytes"
	"fmt"
	lua "github.com/toophy/gopher-lua"
	"github.com/toophy/pangu/help"
)

// 事件 : lua使用的通用事件
type Event_from_lua_screen struct {
	help.Evt_base
	sid      int32      // 场景ID
	module   string     // lua模块名
	function string     // lua函数名
	param    lua.LValue // 参数(table)
}

// 事件执行
func (this *Event_from_lua_screen) Exec(home interface{}) bool {
	// 当前线程调用-> 执行这个事件
	switch home.(type) {
	case *ScreenThread:
		home.(*ScreenThread).Tolua_Common_Screen_Param(this.module, this.function, this.param, this.sid)
	}
	return true
}

// func (this *Event_heart_beat) SayHello(home interface{}) {
// 	evt := &Event_thread_hello{SrcThread: this.Screen_.Get_thread().Get_thread_id(), Chat: /*help.RandStr(5)*/ "nimei", Replay: false}
// 	if this.Screen_.Get_thread().Get_thread_id() == 1 {
// 		evt.DstThread = 2
// 	} else if this.Screen_.Get_thread().Get_thread_id() == 2 {
// 		evt.DstThread = 1
// 	}
// 	evt.Init("", 100)
// 	home.(*ScreenThread).PostThreadMsg(evt.DstThread, evt)
// }

// 事件 : 线程问好
type Event_thread_hello struct {
	help.Evt_base
	SrcThread int32
	DstThread int32
	Chat      string
	Replay    bool
}

// 事件执行
func (this *Event_thread_hello) Exec(home interface{}) bool {

	// home.(*ScreenThread).LogDebug("%s", this.Chat)

	if !this.Replay {
		evt := &Event_thread_hello{SrcThread: this.DstThread, DstThread: this.SrcThread, Chat: /*"r " +*/ this.Chat, Replay: true}
		evt.Init("", 100)
		home.(*ScreenThread).PostThreadMsg(evt.DstThread, evt)
	}

	return true
}

func (this *Event_thread_hello) PrintSelf() {
	fmt.Printf("   {E} %d->%d %s\n", this.SrcThread, this.DstThread, this.Chat)
}
