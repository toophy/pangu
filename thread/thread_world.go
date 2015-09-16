package thread

import (
	"bytes"
	"errors"
	lua "github.com/toophy/gopher-lua"
	"github.com/toophy/pangu/help"
	"os"
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
	threadIds   map[int32]IThread
	luaState    *lua.LState
	buffs       bytes.Buffer // 日志总缓冲
	logFile     *os.File
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
	err := this.Init_thread(self, Tid_master, name, heart_time, lay1_time)
	if err == nil {
		this.threadCount = 0
		this.threadIds = make(map[int32]IThread, 0)

		this.buffs.Grow(LogBuffSize)

		if !help.IsExist(LogFileName) {
			os.Create(LogFileName)
		}
		file, err := os.OpenFile(LogFileName, os.O_RDWR, os.ModePerm)
		if err != nil {
			return err
		}
		this.logFile = file
		this.logFile.Seek(0, 2)
		return nil
	}
	return err
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

	// 处理文件
	evt := &Event_flush_log{}
	evt.Init("", 300)
	this.PostEvent(evt)

	sc1, err := New_screen_thread(Tid_screen_1, "场景线程1", 100, Evt_lay1_time)
	if err == nil && sc1 != nil {
		sc1.Run_thread()
	} else {
		if err != nil {
			this.LogError("新建场景线程1失败:" + err.Error())
		} else {
			this.LogError("新建场景线程1失败:")
		}
	}

	evt1 := &Event_close_thread{}
	evt1.Init("", 10000)
	evt1.Master = sc1
	this.PostEvent(evt1)

	sc2, err := New_screen_thread(Tid_screen_2, "场景线程2", 100, Evt_lay1_time)
	if err == nil && sc2 != nil {
		sc2.Run_thread()
	} else {
		if err != nil {
			this.LogError("新建场景线程2失败:" + err.Error())
		} else {
			this.LogError("新建场景线程2失败:")
		}
	}

	evt2 := &Event_close_thread{}
	evt2.Init("", 10000)
	evt2.Master = sc2
	this.PostEvent(evt2)
}

// 响应线程退出
func (this *WorldThread) on_end() {
	if this.luaState != nil {
		this.luaState.Close()
		this.luaState = nil
		this.logFile.Close()
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
	this.luaState.DoFile("data/thread_init.lua")

	// 加载所有 screens 文件夹里面的 *.lua 文件
	this.luaState.RequireDir("data/world")

	return nil
}

func (this *WorldThread) Add_log(d string) {
	this.buffs.WriteString(d)
}

func (this *WorldThread) Flush_log() {
	if this.buffs.Len() > 0 {
		this.logFile.Write(this.buffs.Bytes())
		this.buffs.Reset()
	}
}
