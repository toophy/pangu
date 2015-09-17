package thread

import (
	lua "github.com/toophy/gopher-lua"
)

// 获取用Lua类型封装结构指针  *LUserData
func (this *WorldThread) GetLUserData(n string, a interface{}) *lua.LUserData {

	ud := this.luaState.NewUserData()
	ud.Value = a
	this.luaState.SetMetatable(ud, this.luaState.GetTypeMetatable(n))

	return ud
}

// 调用Lua函数 : OnInitWorld
func (this *WorldThread) Tolua_OnInitWorld() (ret int) {
	// 捕捉异常
	defer func() {
		if r := recover(); r != nil {
			ret = -1
			this.LogFatal("Tolua_OnInitWorld : " + r.(error).Error())
		}
	}()

	// 调用Lua脚本函数
	if err := this.luaState.CallByParam(lua.P{
		Fn:      this.luaState.GetFunction("main", "OnInitWorld"), // 调用的Lua函数
		NRet:    1,                                                // 返回值的数量
		Protect: true,                                             // 保护?
	}); err != nil {
		panic(err)
	}

	// 处理Lua脚本函数返回值
	ret_lua := this.luaState.Get(-1)
	ret = int(ret_lua.(lua.LNumber))
	this.luaState.Pop(1)

	return
}

// 调用Lua函数 : OnInitScreen
func (this *WorldThread) Tolua_CommanFunction(m string, f string, t *lua.LTable) (ret *lua.LTable) {
	// 捕捉异常
	defer func() {
		if r := recover(); r != nil {
			ret = nil
			this.LogFatal("WorldThread:Tolua_CommanFunction (" + m + "," + f + ") : " + r.(error).Error())
		}
	}()

	// 调用Lua脚本函数
	if err := this.luaState.CallByParam(lua.P{
		Fn:      this.luaState.GetFunction(m, f), // 调用的Lua函数
		NRet:    1,                               // 返回值的数量
		Protect: true,                            // 保护?
	}, t); err != nil {
		panic(err)
	}

	// 处理Lua脚本函数返回值
	ret_lua := this.luaState.Get(-1)
	ret = ret_lua.(*lua.LTable)
	this.luaState.Pop(1)

	return
}
