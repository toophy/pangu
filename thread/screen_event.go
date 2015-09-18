package thread

import (
	_ "bytes"
	"fmt"
	"github.com/toophy/pangu/help"
)

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
