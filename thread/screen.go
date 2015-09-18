package thread

import (
	"fmt"
	lua "github.com/toophy/gopher-lua"
	"github.com/toophy/pangu/actor"
	"github.com/toophy/pangu/help"
)

type Screen struct {
	help.EventObj
	Name    string
	ModName string
	Id      int32
	Oid     int32
	Actors  map[int64]*actor.Actor
	luaData lua.LValue
	thread  *ScreenThread
}

func (this *Screen) Load(name string, id int32, oid int32, t *ScreenThread) bool {
	if t == nil {
		fmt.Println("场景线程不存在")
		return false
	}
	this.InitEventHeader()
	config := screen_config.GetScreenConfig(oid)
	if config == nil {
		t.LogError("场景%s加载失败: 没有找到场景模板(%d)", name, oid)
		return false
	}

	if len(name) > 0 {
		this.Name = name
	} else {
		this.Name = config.Name
	}

	this.ModName = config.ModName

	this.Id = id
	this.Oid = oid
	this.Actors = make(map[int64]*actor.Actor, 0)
	this.thread = t
	this.luaData = this.thread.GetLuaState().NewTable()
	t.LogInfo("场景%s加载成功", this.Name)
	this.Tolua_screen_init()

	return true
}

// 场景卸载
// 场景关联的定时器, 事件, 统统要卸载掉
// 场景内的精灵呢? 有些定时器, 事件, 也是场景关联的
// 没有场景的精灵怎么进行操作呢?
func (this *Screen) Unload() {
	this.thread.LogInfo("场景%s卸载成功", this.Name)
	//this.thread.RemoveEventList(this.GetEventHeader())
	this.thread = nil
	this.luaData = nil
}

// 获取场景管理luaTable
func (this *Screen) Get_data() lua.LValue {
	return this.luaData
}

// 获取场景名称
func (this *Screen) Get_name() string {
	return this.Name
}

// 获取模块(Lua)名称
func (this *Screen) Get_mod_name() string {
	return this.ModName
}

// 获取场景ID
func (this *Screen) Get_id() int32 {
	return this.Id
}

// 获取场景模板ID
func (this *Screen) Get_oid() int32 {
	return this.Oid
}

// !!! 获取的指针不能保存, 获取场景配置
func (this *Screen) Get_config() *Config {
	return screen_config.GetScreenConfig(this.Oid)
}

// 获取所在线程
func (this *Screen) Get_thread() *ScreenThread {
	return this.thread
}

// 投递事件
func (this *Screen) PostEvent(f string, t uint64, p lua.LValue) bool {
	evt := &Event_from_lua_screen{}
	evt.Init("", t)
	evt.sid = this.Get_id()
	evt.module = this.ModName
	evt.function = f
	evt.param = p
	return this.thread.PostEvent(evt)
}

// 登录地图
func (this *Screen) Actor_enter(a *actor.Actor) bool {
	if _, ok := this.Actors[a.GetId()]; ok {
		this.Actors[a.GetId()] = a
		return true
	}
	return false
}

// 离开地图
func (this *Screen) Actor_leave(id int64) bool {
	if _, ok := this.Actors[id]; ok {
		delete(this.Actors, id)
		return true
	}
	return false
}

// 角色移动
func (this *Screen) Actor_move(id int64, pos help.Vec3, check bool) {
	// 如果 check 为 true
	// 主要是位置检查
	// 1. 边界
	// 2. 障碍检查
	// 3. Actor碰撞检查
	// 否则
	// 1. 边界检查
}

// 角色移动验证
func (this *Screen) Actor_move_check(id int64, pos help.Vec3) {
	// 主要是位置检查
	// 1. 边界
	// 2. 障碍检查
	// 3. Actor碰撞检查
}
