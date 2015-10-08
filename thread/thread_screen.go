package thread

import (
	"errors"
	lua "github.com/toophy/gopher-lua"
	"github.com/toophy/pangu/help"
)

// 场景容器
type ScreenMap map[int32]*Screen

// 场景线程
type ScreenThread struct {
	Thread

	lastScreenId int32          // 最后一个场景id
	screens      ScreenMap      // screen 列表
	luaState     *lua.LState    // Lua实体
	luaNilTable  lua.LTable     // Lua空的Table, 供默认参数使用
	move_Actors  help.DListNode // 移动中的角色
}

// 新建场景线程
func New_screen_thread(id int32, name string, heart_time int64, lay1_time uint64) (*ScreenThread, error) {
	a := new(ScreenThread)
	err := a.Init_screen_thread(id, name, heart_time, lay1_time)
	if err == nil {
		return a, nil
	}
	return nil, err
}

// 初始化场景线程
func (this *ScreenThread) Init_screen_thread(id int32, name string, heart_time int64, lay1_time uint64) error {
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

// 增加场景
func (this *ScreenThread) Add_screen(name string, oid int32) bool {
	a := new(Screen)
	if !a.Load(name, this.lastScreenId, oid, this) {
		return false
	}

	this.screens[this.lastScreenId] = a

	this.lastScreenId++

	return true
}

// 删除场景
func (this *ScreenThread) Del_screen(id int32) bool {
	if _, ok := this.screens[id]; ok {
		this.screens[id].Unload()
		delete(this.screens, id)
		return true
	}
	return false
}

// 获取场景
func (this *ScreenThread) Get_screen(id int32) *Screen {
	if _, ok := this.screens[id]; ok {
		return this.screens[id]
	}
	return nil
}

// 响应线程首次运行
func (this *ScreenThread) on_first_run() {

	errInit := this.reloadLuaState()
	if errInit != nil {
		this.LogError(errInit.Error())
		return
	}

	this.Tolua_Common("main", "OnScreenThreadBegin")
}

// 响应线程退出
func (this *ScreenThread) on_end() {
	this.Tolua_Common("main", "OnScreenThreadEnd")
	if this.luaState != nil {
		this.luaState.Close()
		this.luaState = nil
	}
}

// 响应线程运行
func (this *ScreenThread) on_run() {
}

// 初始化LuaState, 可以用来 Reload LuaState
func (this *ScreenThread) reloadLuaState() error {

	if this.luaState != nil {
		this.luaState.Close()
		this.luaState = nil
	}

	this.luaState = lua.NewState()
	if this.luaState == nil {
		return errors.New("场景线程初始化Lua失败")
	}

	RegLua_all_thread_screen(this.luaState)
	RegLua_all_screen(this.luaState)

	// 注册公告变量-->本线程
	this.luaState.SetGlobal("ts", this.GetLUserData("ScreenThread", this))

	// 执行初始化脚本
	this.luaState.Require("data/thread_init")

	// 加载所有 screens 文件夹里面的 *.lua 文件
	this.luaState.RequireDir("data/screens")

	return nil
}

// !!!只能获取, 不准许保存指针, 获取LState
func (this *ScreenThread) GetLuaState() *lua.LState {
	return this.luaState
}

// lua投递事件
func (this *ScreenThread) PostEventFromLua(m string, f string, t uint64, p lua.LValue) bool {
	evt := &Event_from_lua{}
	evt.Init("", t)
	evt.module = m
	evt.function = f
	evt.param = p
	return this.PostEvent(evt)
}

// 新增移动中的角色
func (this *ScreenThread) AddMoveActor(n *help.DListNode) {
	old_pre := this.move_Actors.Pre

	this.move_Actors.Pre = n
	n.Next = &this.move_Actors
	n.Pre = old_pre
	old_pre.Next = n
}
