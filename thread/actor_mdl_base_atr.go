package thread

import (
	"github.com/toophy/pangu/help"
)

type BaseAtr struct {
	Name string    // 名称
	Pos  help.Vec3 // 当前位置
}

func (a *Actor) BaseAtr_getName() string {
	v := a.Mdl_check(Amdl_BaseAtr)

	if v != nil {
		return v.(*BaseAtr).Name
	}

	return ""
}

func (a *Actor) BaseAtr_setName(name string) {
	v := a.Mdl_check(Amdl_BaseAtr)

	if v != nil {
		v.(*BaseAtr).Name = name
	}
}

func (a *Actor) BaseAtr_getPos() help.Vec3 {
	v := a.Mdl_check(Amdl_BaseAtr)

	if v != nil {
		return v.(*BaseAtr).Pos
	}

	return help.Vec3{}
}

func (a *Actor) BaseAtr_setPos(p help.Vec3) {
	v := a.Mdl_check(Amdl_BaseAtr)

	if v != nil {
		v.(*BaseAtr).Pos = p
	}
}

func (a *Actor) BaseAtr_isMoving() bool {
	return !a.MoveNode.IsEmpty()
}

func (a *Actor) BaseAtr_moveTo(v help.Vec3) {
	if a.MoveNode.IsEmpty() {
		// 移动新增加
	}

}
