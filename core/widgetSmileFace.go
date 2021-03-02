package core

import (
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type SmileFace struct {
	widget.Icon
	event func()
}

func (this *SmileFace) MouseDown(m *desktop.MouseEvent) {
	this.SetResource(resources.face1)
}

func (this *SmileFace) MouseUp(m *desktop.MouseEvent) {
	this.SetResource(resources.face0)
	this.event()
}

func (this *SmileFace) AddEvent(event func()) {
	this.event = event
}

func newSmileFace(event func()) *SmileFace {
	s := &SmileFace{event: event}
	s.SetResource(resources.face0)
	s.ExtendBaseWidget(s)
	return s
}
