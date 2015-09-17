package thread

import (
	lua "github.com/toophy/gopher-lua"
)

// 获取用Lua类型封装结构指针  *LUserData
func (this *ScreenThread) GetLUserData(n string, a interface{}) *lua.LUserData {

	ud := this.luaState.NewUserData()
	ud.Value = a
	this.luaState.SetMetatable(ud, this.luaState.GetTypeMetatable(n))

	return ud
}

// 调用Lua函数 : 调用Lua函数
func (this *ScreenThread) Tolua_CommanFunction(m string, f string, t *lua.LTable) (ret lua.LValue) {
	// 捕捉异常
	defer func() {
		if r := recover(); r != nil {
			ret = nil
			this.LogFatal("ScreenThread:Tolua_CommanFunction (" + m + "," + f + ") : " + r.(error).Error())
		}
	}()

	if t == nil {
		t = &this.luaNilTable
	}

	// 调用Lua脚本函数
	if err := this.luaState.CallByParam(lua.P{
		Fn:      this.luaState.GetFunction(m, f), // 调用的Lua函数
		NRet:    1,                               // 返回值的数量
		Protect: true,                            // 保护?
	}, t); err != nil {
		panic(err)
	}

	// 处理Lua脚本函数返回值
	ret = this.luaState.Get(-1)
	this.luaState.Pop(1)
	return
}
