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

// 当前不移动的情况下, 才能进行新的移动
func (a *Actor) Base_moveTo(s *help.Vec3, d *help.Vec3, speed float32) bool {
	if a.MoveNode.IsEmpty() {
		l := a.SrcPos.Dist(a.DstPos)
		if l > 0.01 {
			// 移动新增加
			a.MoveLen = l

			a.SrcPos = s
			a.DstPos = d

			a.MoveRate.X = (a.DstPos.X - a.SrcPos.X) / l
			a.MoveRate.Y = (a.DstPos.Y - a.SrcPos.Y) / l
			a.MoveRate.Z = (a.DstPos.Z - a.SrcPos.Z) / l

			a.MoveSpeed = speed * 0.001

			a.MoveMinLen = a.MoveSpeed * 50 // 四舍五入

			a.LastMoveTime = a.CurScreen.Get_thread().GetCurrTime()

			a.CurScreen.Get_thread().AddMoveActor(&a.MoveNode)
			return true
		}
	}

	return false
}

// 移动回调,返回false退出移动列表
func (a *Actor) Base_onMove() bool {
	// 本线段上总共移动的距离
	s := float32((a.CurScreen.Get_thread().GetCurrTime() - a.LastMoveTime) * a.MoveSpeed)
	// 剩余距离小于xx, 立即移动到结束位置上, 并结束移动
	if (s + a.MoveMinLen) > a.MoveLen {
		// 到达目标, 结束移动
		a.CurPos = a.DstPos
		// a.MoveNode.Pop()
		return false
	}

	// 移动继续
	var n help.Vec3
	n.X = a.SrcPos.X + a*a.MoveRate.X
	n.Y = a.SrcPos.Y + a*a.MoveRate.Y
	n.Z = a.SrcPos.Z + a*a.MoveRate.Z

	// 触发移动点问题, 当前点, 目标点
	a.CurPos = n
	return true
}
