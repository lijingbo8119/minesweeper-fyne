package core

type RefEvent func(val interface{})

type Ref struct {
	val    interface{}
	events []RefEvent
}

func (this *Ref) AddListener(e RefEvent) *Ref {
	this.events = append(this.events, e)
	return this
}

func (this *Ref) Get() interface{} {
	return this.val
}

func (this *Ref) Set(val interface{}) *Ref {
	this.val = val
	for _, e := range this.events {
		e(val)
	}
	return this
}
