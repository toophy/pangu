package thread

import (
	"github.com/toophy/pangu/help"
	"sync"
)

var G_thread_msg_pool ThreadMsgPool

func init() {
	G_thread_msg_pool.Init()
}

// 线程间消息存放处
type ThreadMsgPool struct {
	lock   [Tid_last]sync.RWMutex    // 每个线程的消息池有一个独立的读写锁
	header [Tid_last]*help.DListNode // 每个线程的消息池
}

// 初始化
func (this *ThreadMsgPool) Init() {
	for i := 0; i < Tid_last; i++ {
		this.header[i] = new(help.DListNode)
		this.header[i].Init(nil)
	}
}

// 投递线程间消息
func (this *ThreadMsgPool) PostMsg(tid int32, e *help.DListNode) bool {
	if tid >= Tid_master && tid < Tid_last {
		this.lock[tid].Lock()
		defer this.lock[tid].Unlock()

		header := this.header[tid]

		e_pre := e.Pre
		e_next := e.Next

		e.Init(nil)

		header.Pre.Next = e_pre
		e_pre.Pre = header.Pre

		header.Pre = e_next
		e_next.Next = header

		return true
	}
	return false
}

// 获取线程间消息
func (this *ThreadMsgPool) GetMsg(tid int32, e *help.DListNode) bool {
	if tid >= Tid_master && tid < Tid_last {
		this.lock[tid].Lock()
		defer this.lock[tid].Unlock()

		header := this.header[tid]

		header_pre := header.Pre
		header_next := header.Next

		header.Init(nil)

		e.Pre.Next = header_pre
		header_pre.Pre = e.Pre

		e.Pre = header_next
		header_next.Next = e

		return true
	}
	return false
}
