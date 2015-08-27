package help

type IEventObj interface {
	AddEvent(n *DListNode)
}

type EventObj struct {
	NodeObj DListNode
}

func (this *EventObj) InitEventHeader() {
	this.NodeObj.Init(nil)
}

func (this *EventObj) GetEventHeader() *DListNode {
	return &this.NodeObj
}

func (this *EventObj) AddEvent(n *DListNode) {
}
