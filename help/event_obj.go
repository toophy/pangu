package help

type EventObj struct {
	NodeObj DListNode
}

func (this *EventObj) InitEventHeader() {
	this.NodeObj.Init(nil)
}

func (this *EventObj) GetEventHeader() *DListNode {
	return &this.NodeObj
}

func (this *EventObj) AddEvent(e IEvent) bool {
	n := &DListNode{}
	n.Init(e)

	if !e.AddNode(n) {
		return false
	}

	n.Pre = this.NodeObj.Pre
	this.NodeObj.Pre.Next = n
	this.NodeObj.Pre = n
	n.Next = &this.NodeObj

	return true
}
