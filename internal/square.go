package internal

import (
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"time"
)

type Square struct {
	widget.Icon
	SquareStatus     SquareStatus
	SquareType       SquareType
	SquareCoordinate Coordinate
	AroundSquares    Squares
}

func (this *Square) MouseDown(m *desktop.MouseEvent) {
	switch m.Button {
	case desktop.MouseButtonPrimary:
		State.leftMouseDown(this.SquareCoordinate)
	case desktop.MouseButtonSecondary:
		State.rightMouseDown(this.SquareCoordinate)
	}
}

func (this *Square) MouseUp(m *desktop.MouseEvent) {
	switch m.Button {
	case desktop.MouseButtonPrimary:
		State.leftMouseUp(this.SquareCoordinate)
	case desktop.MouseButtonSecondary:
		State.rightMouseUp(this.SquareCoordinate)
	}
}

func (this *Square) setStatus(s SquareStatus) {
	switch s {
	case SquareStatusClosed:
		this.SetResource(resources.blank)
	case SquareStatusMouseDown:
		this.SetResource(resources.mousedown)
	case SquareStatusMarkedFlag:
		this.SetResource(resources.flag)
	case SquareStatusOpened:
		switch this.SquareStatus {
		case SquareStatusClosed:
			switch this.SquareType {
			case SquareTypeNormal:
				switch len(this.AroundSquares.filter(func(square *Square) bool { return square.SquareType == SquareTypeMine })) {
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
				}
			case SquareTypeMine:
				this.SetResource(resources.mine0)
			}
		case SquareStatusMouseDown:
			switch this.SquareType {
			case SquareTypeNormal:
				switch len(this.AroundSquares.filter(func(square *Square) bool { return square.SquareType == SquareTypeMine })) {
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
				}
			case SquareTypeMine:
				this.SetResource(resources.mine0)
			}
		case SquareStatusMarkedFlag:
			switch this.SquareType {
			case SquareTypeNormal:
				this.SetResource(resources.mine1)
			case SquareTypeMine:
			}
		}
	case SquareStatusExploded:
		this.SetResource(resources.mine2)
	}
	this.SquareStatus = s
}

func (this *Square) open(triggeredByClick bool) bool {
	if this.SquareStatus != SquareStatusClosed && this.SquareStatus != SquareStatusMouseDown {
		return true
	}
	if this.SquareType == SquareTypeMine && triggeredByClick {
		this.setStatus(SquareStatusExploded)
		t := time.Now()
		State.endTime = &t
		return false
	}

	unmarkedMines := this.AroundSquares.filter(func(s *Square) bool {
		return s.SquareType == SquareTypeMine && s.SquareStatus != SquareStatusMarkedFlag
	})
	if triggeredByClick {
		if len(unmarkedMines) == 0 {
			this.openAroundSquares()
		}
		this.setStatus(SquareStatusOpened)
	}

	if !triggeredByClick && this.SquareType == SquareTypeNormal {
		this.setStatus(SquareStatusOpened)
	}

	if !triggeredByClick && this.SquareType == SquareTypeNormal && len(unmarkedMines) == 0 {
		this.openAroundSquares()
	}

	return true
}

func (this *Square) mark() {
	if this.SquareStatus != SquareStatusClosed && this.SquareStatus != SquareStatusMouseDown && this.SquareStatus != SquareStatusMarkedFlag {
		return
	}
	if this.SquareStatus == SquareStatusMarkedFlag {
		this.setStatus(SquareStatusClosed)
		return
	}
	this.setStatus(SquareStatusMarkedFlag)
}

func (this *Square) openAroundSquares() {
	if unmarkedMines := this.AroundSquares.filter(func(s *Square) bool {
		return s.SquareType == SquareTypeMine && s.SquareStatus != SquareStatusMarkedFlag
	}); len(unmarkedMines) > 0 {
		return
	}
	for _, s := range this.AroundSquares {
		s.open(false)
	}
}

func newSquare(t SquareType) *Square {
	s := &Square{
		SquareType: t,
	}
	s.ExtendBaseWidget(s)
	s.setStatus(SquareStatusClosed)
	return s
}
