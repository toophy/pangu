package thread

import (
	"errors"
	lua "github.com/toophy/gopher-lua"
	"github.com/toophy/pangu/help"
)

// 网络线程
type NetThread struct {
	Thread
}

// 新建网络线程
func New_net_thread(id int32, name string, heart_time int64, lay1_time uint64) (*NetThread, error) {
	a := new(NetThread)
	err := a.Init_net_thread(id, name, heart_time, lay1_time)
	if err == nil {
		return a, nil
	}
	return nil, err
}

// 初始化网络线程
func (this *NetThread) Init_net_thread(id int32, name string, heart_time int64, lay1_time uint64) error {
	if id < Tid_screen_1 || id > Tid_screen_9 {
		return errors.New("线程ID超出范围 [Tid_screen_1,Tid_screen_9]")
	}
	err := this.Init_thread(this, id, name, heart_time, lay1_time)
	if err == nil {
		this.screens = make(ScreenMap, 0)
		this.lastScreenId = (id - 1) * 1000000

		// 移动中角色, 节点初始化
		this.move_Actors.Init(nil)
		this.move_Actors.SrcTid = this.id
		return nil
	}
	return err
}

// 响应线程首次运行
func (this *NetThread) on_first_run() {
}

// 响应线程最先运行
func (this *NetThread) on_pre_run() {
}

// 响应线程运行
func (this *NetThread) on_run() {
}

// 响应线程退出
func (this *NetThread) on_end() {
}
