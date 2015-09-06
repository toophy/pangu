package thread

import (
	"github.com/toophy/pangu/help"
)

// 事件 : 线程关闭
type Event_close_thread struct {
	help.Evt_base
	Master IThread
}

// 事件执行
func (this *Event_close_thread) Exec(home interface{}) bool {
	if this.Master != nil {
		this.Master.pre_close_thread()
		return true
	}

	println("没找到线程")
	return true
}

// 事件 : 释放节点
type Event_pre_release_dlinknode struct {
	help.Evt_base
}

// 事件执行
func (this *Event_pre_release_dlinknode) Exec(home interface{}) bool {
	// 从线程中释放这个节点
	home.(IThread).sendThreadFreeNode()

	//
	evtReleaseNode := &Event_pre_release_dlinknode{}
	evtReleaseNode.Init("", 3000)
	home.(IThread).PostEvent(evtReleaseNode)

	return true
}

// 事件 : 释放节点
type Event_release_dlinknode struct {
	help.Evt_base
	Header help.DListNode
}

// 事件执行
func (this *Event_release_dlinknode) Exec(home interface{}) bool {
	// 从线程中释放这个节点
	home.(IThread).releaseDlinkNode(&this.Header)

	return true
}
