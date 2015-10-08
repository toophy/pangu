package thread

import ()

// 帐号,可以存在于地图中,指挥多个Actor
type Account struct {
	Name      string  // 帐号名
	CurScreen *Screen // 当前场景
	CurActor  *Actor  // 当前角色
}
