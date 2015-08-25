package base

import (
	"fmt"
	"testing"
)

type IEvent interface {
	Exec(home interface{}) bool
}

type EvtPool struct {
	header DListNode
}

func (this *EvtPool) Init() {
	this.header.Init(nil)
}

func (this *EvtPool) Eat(name string) {
	fmt.Printf("吃%s\n", name)
}

func (this *EvtPool) Post(d interface{}) {
	n := &DListNode{}
	n.Init(d)

	n.Pre = this.header.Pre

	this.header.Pre.Next = n
	this.header.Pre = n

	n.Next = &this.header
}

func (this *EvtPool) Run() {
	for {
		if this.header.IsEmpty() {
			break
		}

		n := this.header.Next

		n.Data.(IEvent).Exec(this)

		n.Pre.Next = n.Next
		n.Next.Pre = n.Pre
		n.Data = nil
	}

}

type Evt_eat struct {
	FoodName string
}

func (this *Evt_eat) Exec(home interface{}) bool {

	home.(*EvtPool).Eat(this.FoodName)

	return true
}

func TestDlist(t *testing.T) {

	var g_Pool EvtPool
	g_Pool.Init()

	g_Pool.Post(&Evt_eat{FoodName: "西瓜"})
	g_Pool.Post(&Evt_eat{FoodName: "葡萄"})
	g_Pool.Post(&Evt_eat{FoodName: "黄瓜"})
	g_Pool.Post(&Evt_eat{FoodName: "大蒜"})

	g_Pool.Run()
}
