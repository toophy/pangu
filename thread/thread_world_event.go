package thread

import (
	"github.com/toophy/pangu/help"
)

// 事件 : 线程投递的日志
type Event_thread_log struct {
	help.Evt_base
	Data string
}

// 事件执行
func (this *Event_thread_log) Exec(home interface{}) bool {
	home.(*WorldThread).Add_log(this.Data)
	return true
}
