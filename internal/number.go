package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/gogf/gf/util/gconv"
	"strings"
)

type number struct {
	widget.Icon
}

func (this *number) SetNumber(n int) *number {
	this.ExtendBaseWidget(this)
	if n > 9 {
		panic("SetNumber error")
	}
	switch n {
	case 0:
		this.SetResource(resources.num0)
	case 1:
		this.SetResource(resources.num1)
	case 2:
		this.SetResource(resources.num2)
	case 3:
		this.SetResource(resources.num3)
	case 4:
		this.SetResource(resources.num4)
	case 5:
		this.SetResource(resources.num5)
	case 6:
		this.SetResource(resources.num6)
	case 7:
		this.SetResource(resources.num7)
	case 8:
		this.SetResource(resources.num8)
	case 9:
		this.SetResource(resources.num9)
	}
	this.Refresh()
	return this
}

type Number struct {
	Container *fyne.Container
	numbers   []*number
}

func (this *Number) SetNumber(n int) {
	if n < 0 {
		panic("SetNumber error")
	}
	_n := strings.Split(gconv.String(n), "")
	switch len(_n) {
	case 1:
		this.numbers[0].SetNumber(0)
		this.numbers[1].SetNumber(0)
		this.numbers[2].SetNumber(gconv.Int(_n[0]))
	case 2:
		this.numbers[0].SetNumber(0)
		this.numbers[1].SetNumber(gconv.Int(_n[0]))
		this.numbers[2].SetNumber(gconv.Int(_n[1]))
	default:
		this.numbers[0].SetNumber(gconv.Int(_n[len(_n)-3]))
		this.numbers[1].SetNumber(gconv.Int(_n[len(_n)-2]))
		this.numbers[2].SetNumber(gconv.Int(_n[len(_n)-1]))
	}
}

func NewNumberContainer() Number {
	instance := Number{}
	instance.numbers = []*number{
		new(number).SetNumber(0),
		new(number).SetNumber(0),
		new(number).SetNumber(0),
	}
	instance.Container = container.New(layout.NewHBoxLayout(), instance.numbers[0], instance.numbers[1], instance.numbers[2])
	return instance
}
