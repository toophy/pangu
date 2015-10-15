package thread

import (
	"github.com/toophy/pangu/help"
)

const (
	Atype_null = iota
	Atype_player
	Atype_monster
	Atype_npc
	Atype_pupet
	Atype_bus
	Atype_item
	Atype_last
)

const (
	Amdl_ExAtr = iota
	Amdl_Last
)

// 演员 : 地图上所有对象
// usage :
// a := Actor{}
type Actor struct {
	Mdls      map[int32]interface{} // 模块
	Id        int64                 // 角色编号
	Type      int32                 // 角色类型
	Name      string                // 名称
	CurScreen *Screen               // 当前所在场景

	MoveNode     help.DListNode // Move : 计算"移动"的节点
	CurPos       help.Vec3      // Move : 当前位置
	SrcPos       help.Vec3      // Move : 移动开始位置
	DstPos       help.Vec3      // Move : 移动目标位置
	MoveRate     help.Vec3      // Move : 移动系数 (x3-x1)/s, (y3-y1)/s, (z3-z1)/s
	MoveSpeed    float32        // Move : 移动速度
	MoveMinLen   float32        // Move : 最小容忍距离(小于这个值会立即移动)
	MoveLen      float32        // Move : 移动总长度
	LastMoveTime int64          // Move : 最后移动时间戳(单位: 毫秒)
}

// 演员初始化
// usage :
// a := Actor{}
// a.Init(Atype_player)
func (a *Actor) Init(t int32, id int64, name string, pos help.Vec3, s *Screen) bool {
	if t <= Atype_null || t >= Atype_last {
		return false
	}
	a.Type = t
	a.Id = id
	a.Mdls = make(map[int32]interface{}, 0)
	a.Name = name
	a.CurScreen = s
	a.CurrPos = pos
	a.MoveNode.Init(a)
	a.LastMoveTime = a.CurScreen.Get_thread().GetCurrTime()
	return true
}

// 增加演员功能模块
// usage :
// a := Actor{}
// b := BaseAtr{}
// a.Mdl_add(&b)
//
func (a *Actor) Mdl_add(m interface{}) bool {
	id := int32(Amdl_Last)

	switch m.(type) {
	case *ExAtr:
		id = Amdl_ExAtr
	}

	if id != Amdl_Last {
		if _, ok := a.Mdls[id]; ok == false {
			a.Mdls[id] = m
			return true
		}
	}

	return false
}

// 删除演员功能模块
// usage :
// a := Actor{}
// emB :=  0
// a.Mdl_del(emB)
//
func (a *Actor) Mdl_del(id int32) {
	if _, ok := a.Mdls[id]; ok {
		delete(a.Mdls, id)
	}
}

// 检查演员功能模块是否存在, 返回结构指针
// usage :
// a := Actor{}
// emB :=  0
// v := a.Mdl_check(emB)
// if v != nil {
//   调用模块
//   v.(*BaseAtr).Name = name
// }
func (a *Actor) Mdl_check(id int32) interface{} {
	if v, ok := a.Mdls[id]; ok {
		return v
	}
	return nil
}
