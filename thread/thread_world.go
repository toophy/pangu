package thread

import (
	"errors"
	lua "github.com/toophy/gopher-lua"
	"sync"
	"time"
)

const (
	LogBuffSize = 10 * 1024 * 1024
	LogDir      = "../log"
	ProfFile    = "pangu_prof.log"
	LogFileName = LogDir + "/pangu.log"
)

// 主线程
type WorldThread struct {
	Thread
	threadLock  sync.RWMutex
	threadCount int32
	threadIds   map[int32]IThread // 线程池
	luaState    *lua.LState       // Lua实体
	luaNilTable lua.LTable        // Lua空的Table, 供默认参数使用
}

var myWorldThread *WorldThread = nil

// 获取主线程
func GetWorld() *WorldThread {
	if myWorldThread == nil {
		myWorldThread = &WorldThread{}
		err := myWorldThread.Init_master_thread(myWorldThread, "主线程", 100, Evt_lay1_time)
		if err != nil {
			panic(err.Error())
		}
		myWorldThread.Run_thread()
	}
	return myWorldThread
}

// 初始化主线程
func (this *WorldThread) Init_master_thread(self IThread, name string, heart_time int64, lay1_time uint64) error {
	err := this.Init_thread(self, Tid_world, name, heart_time, lay1_time)
	if err == nil {
		this.threadCount = 0
		this.threadIds = make(map[int32]IThread, 0)

		return nil
	}
	return err
}

// 检查线程是否已经存在
func (this *WorldThread) Can_create_thread(tid int32) bool {
	this.threadLock.Lock()
	defer this.threadLock.Unlock()

	if _, ok := this.threadIds[tid]; ok == false {
		return true
	}

	return false
}

// 增加运行的线程
func (this *WorldThread) Add_run_thread(a IThread) {
	this.threadLock.Lock()
	defer this.threadLock.Unlock()

	if _, ok := this.threadIds[a.Get_thread_id()]; ok == false {
		this.threadCount++
		this.threadIds[a.Get_thread_id()] = a
	}
}

// 释放运行的线程
func (this *WorldThread) Release_run_thread(a IThread) {
	this.threadLock.Lock()
	defer this.threadLock.Unlock()

	if _, ok := this.threadIds[a.Get_thread_id()]; ok == true {
		this.threadCount--
		delete(this.threadIds, a.Get_thread_id())
	}
}

// 等待所有线程结束
func (this *WorldThread) Wait_thread_over() {
	for {
		time.Sleep(10 * time.Second)

		this.threadLock.Lock()

		if this.threadCount <= 0 {
			this.threadLock.Unlock()
			time.Sleep(2 * time.Second)
			return
		} else if this.threadCount == 1 {
			evt := &Event_close_thread{}
			evt.Init("", 2000)
			evt.Master = this
			this.PostEvent(evt)
		}

		this.threadLock.Unlock()
	}
}

// 首次运行
func (this *WorldThread) on_first_run() {

	errInit := this.ReloadLuaState()
	if errInit != nil {
		panic(errInit.Error())
	}

	this.Tolua_CommanFunction("main", "OnWorldBegin", nil)
}

// 响应线程退出
func (this *WorldThread) on_end() {
	this.Tolua_CommanFunction("main", "OnWorldEnd", nil)
	if this.luaState != nil {
		this.luaState.Close()
		this.luaState = nil
	}
}

// 响应线程运行
func (this *WorldThread) on_run() {
}

// 初始化LuaState, 可以用来 Reload LuaState
func (this *WorldThread) ReloadLuaState() error {

	if this.luaState != nil {
		this.luaState.Close()
		this.luaState = nil
	}

	this.luaState = lua.NewState()
	if this.luaState == nil {
		return errors.New("[E] 主线程初始化Lua失败")
	}

	RegLua_all_thread_world(this.luaState)

	// 注册公告变量-->本线程
	this.luaState.SetGlobal("ts", this.GetLUserData("WorldThread", this))

	// 执行初始化脚本
	this.luaState.Require("data/thread_init")

	// 加载所有 screens 文件夹里面的 *.lua 文件
	this.luaState.RequireDir("data/world")

	return nil
}

// 新建场景线程
func (this *WorldThread) CreateScreenThread(id int32, name string, heart_time int64, lay1_time uint64, close_time uint64) *ScreenThread {

	if this.Can_create_thread(id) {

		s, err := New_screen_thread(id, name, heart_time, lay1_time)
		if err == nil && s != nil {
			s.Run_thread()
		} else {
			if err != nil {
				this.LogError("新建" + name + "失败:" + err.Error())
			} else {
				this.LogError("新建" + name + "失败:")
			}
			return s
		}

		if close_time > 0 {
			evt1 := &Event_close_thread{}
			evt1.Init("", close_time)
			evt1.Master = s
			this.PostEvent(evt1)
		}

		return s
	}

	return nil
}
