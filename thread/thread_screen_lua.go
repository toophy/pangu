package thread

import (
	lua "github.com/toophy/gopher-lua"
)

// 注册本包所有Lua接口结构
func RegLua_all_thread_screen(L *lua.LState) error {

	type regLuaFunc func(string, *lua.LState) error

	regLuaStructs := map[string]regLuaFunc{
		"ScreenThread": regLua_screen_thread,
	}

	for k, _ := range regLuaStructs {
		regLuaStructs[k](k, L)
	}

	return nil
}

// 向Lua注册结构 : ScreenThread
func regLua_screen_thread(struct_name string, L *lua.LState) error {

	mt := L.NewTypeMetatable(struct_name)
	L.SetGlobal(struct_name, mt)

	// 检查Lua首个参数是不是对象指针
	check := func(L *lua.LState) *ScreenThread {
		ud := L.CheckUserData(1)
		if v, ok := ud.Value.(*ScreenThread); ok {
			return v
		}
		L.ArgError(1, struct_name+" expected")

		return nil
	}

	// 成员函数
	// L.SetField(mt, "new", L.NewFunction(newScreenThread))

	// 成员变量
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(),

		map[string]lua.LGFunction{

			// 增加场景
			"Add_screen": func(L *lua.LState) int {
				p := check(L)
				name := L.CheckString(2)
				oid := int32(L.CheckInt(3))
				ret := p.Add_screen(name, oid)
				L.Push(lua.LBool(ret))
				return 1
			},

			// 删除场景
			"Del_screen": func(L *lua.LState) int {
				p := check(L)
				id := int32(L.CheckInt(2))
				ret := p.Del_screen(id)
				L.Push(lua.LBool(ret))
				return 1
			},

			// 获取场景
			"Get_screen": func(L *lua.LState) int {
				p := check(L)
				id := int32(L.CheckInt(2))
				ret := p.Get_screen(id)
				L.Push(p.GetLUserData("Screen", ret))
				return 1
			},

			// 获取线程号
			"Get_thread_id": func(L *lua.LState) int {
				p := check(L)
				ret := p.Get_thread_id()
				L.Push(lua.LNumber(ret))
				return 1
			},

			// lua投递事件
			"PostEventFromLua": func(L *lua.LState) int {
				p := check(L)
				m := L.CheckString(2)
				f := L.CheckString(3)
				t := uint64(L.CheckInt64(4))
				param := L.CheckAny(5)
				ret := p.PostEventFromLua(m, f, t, param)
				L.Push(lua.LBool(ret))
				return 1
			},

			// LogDebug
			"LogDebug": func(L *lua.LState) int {
				p := check(L)
				data := L.CheckString(2)
				p.LogDebug(data)
				return 1
			},

			// LogInfo
			"LogInfo": func(L *lua.LState) int {
				p := check(L)
				data := L.CheckString(2)
				p.LogInfo(data)
				return 1
			},

			// LogWarn
			"LogWarn": func(L *lua.LState) int {
				p := check(L)
				data := L.CheckString(2)
				p.LogWarn(data)
				return 1
			},

			// LogError
			"LogError": func(L *lua.LState) int {
				p := check(L)
				data := L.CheckString(2)
				p.LogError(data)
				return 1
			},

			// LogFatal
			"LogFatal": func(L *lua.LState) int {
				p := check(L)
				data := L.CheckString(2)
				p.LogFatal(data)
				return 1
			},
		}))

	return nil
}
