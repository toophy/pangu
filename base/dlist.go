package base

import ()

type DListNode struct {
	Pre  *DListNode  // 前一个
	Next *DListNode  // 后一个
	Data interface{} // 事件对象
}

func (this *DListNode) Init(d interface{}) {
	this.Pre = this
	this.Next = this
	this.Data = d
}

func (this *DListNode) Clear() {
	this.Pre = nil
	this.Next = nil
	this.Data = nil
}

func (this *DListNode) IsEmpty() bool {
	return this.Pre == this && this.Next == this
}
