package thread

import ()

// 获取Id
func (a *Actor) Base_getId() int64 {
	return a.Id
}

// 获取类型
func (a *Actor) Base_getType() int32 {
	return a.Type
}

func (a *Actor) Base_getName() string {
	return a.Name
}

func (a *Actor) Base_setName(name string) {
	a.Name = name
}

func (a *Actor) Base_getPos() help.Vec3 {
	return a.CurrPos
}

func (a *Actor) Base_setPos(p help.Vec3) {
	a.CurrPos = p
}

func (a *Actor) Base_isMoving() bool {
	return !a.MoveNode.IsEmpty()
}

func (a *Actor) Base_moveTo(p help.Vec3) bool {
	if a.MoveNode.IsEmpty() {
		// 移动新增加
		a.TargetPos = p

		a.CurScreen.Get_thread().AddMoveActor(&a.MoveNode)
	} else {
		a.TargetPos = p
	}
	return true
}

func (a *Actor) Base_onMove() {
	// 根据过去的时间戳, 向目标移动一点
}
