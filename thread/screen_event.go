package thread

import (
	"fmt"
	"github.com/toophy/pangu/help"
)

// 事件 : 场景心跳
type Event_heart_beat struct {
	help.Evt_base
	Screen_ *Screen
}

// 事件执行
func (this *Event_heart_beat) Exec(home interface{}) bool {

	this.Screen_.Tolua_heart_beat()

	for i := 0; i < 20000; i++ {
		this.SayHello(home)
	}

	evt := &Event_heart_beat{Screen_: this.Screen_}
	evt.Init("", 100)
	this.Screen_.Get_thread().PostEvent(evt)

	return true
}

func (this *Event_heart_beat) SayHello(home interface{}) {
	evt := &Event_thread_hello{SrcThread: this.Screen_.Get_thread().Get_thread_id(), Chat: /*help.RandStr(5)*/ "nimei", Replay: false}
	if this.Screen_.Get_thread().Get_thread_id() == 1 {
		evt.DstThread = 2
	} else if this.Screen_.Get_thread().Get_thread_id() == 2 {
		evt.DstThread = 1
	}
	evt.Init("", 100)
	home.(*ScreenThread).PostThreadMsg(evt.DstThread, evt)
}

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

	//home.(*ScreenThread).LogDebug("%s", this.Chat)

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
