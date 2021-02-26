package internal

import (
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
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
		State.mouseState.LeftMouseDown(this.SquareCoordinate)
	case desktop.MouseButtonSecondary:
		State.mouseState.RightMouseDown(this.SquareCoordinate)
	}
}

func (this *Square) MouseUp(m *desktop.MouseEvent) {
	switch m.Button {
	case desktop.MouseButtonPrimary:
		State.mouseState.LeftMouseUp(this.SquareCoordinate)
	case desktop.MouseButtonSecondary:
		State.mouseState.RightMouseUp(this.SquareCoordinate)
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
	case SquareStatusMarkedWrong:
		this.SetResource(resources.mine1)
	case SquareStatusOpened:
		switch this.SquareStatus {
		case SquareStatusClosed:
			switch this.SquareType {
			case SquareTypeNormal:
				switch len(this.AroundSquares.filter(func(square *Square) bool { return square.SquareType == SquareTypeMine })) {
				case 0:
					this.SetResource(resources.mineNum0)
				case 1:
					this.SetResource(resources.mineNum1)
				case 2:
					this.SetResource(resources.mineNum2)
				case 3:
					this.SetResource(resources.mineNum3)
				case 4:
					this.SetResource(resources.mineNum4)
				case 5:
					this.SetResource(resources.mineNum5)
				case 6:
					this.SetResource(resources.mineNum6)
				case 7:
					this.SetResource(resources.mineNum7)
				case 8:
					this.SetResource(resources.mineNum8)
				}
			case SquareTypeMine:
				this.SetResource(resources.mine0)
			}
		case SquareStatusMouseDown:
			switch this.SquareType {
			case SquareTypeNormal:
				switch len(this.AroundSquares.filter(func(square *Square) bool { return square.SquareType == SquareTypeMine })) {
				case 0:
					this.SetResource(resources.mineNum0)
				case 1:
					this.SetResource(resources.mineNum1)
				case 2:
					this.SetResource(resources.mineNum2)
				case 3:
					this.SetResource(resources.mineNum3)
				case 4:
					this.SetResource(resources.mineNum4)
				case 5:
					this.SetResource(resources.mineNum5)
				case 6:
					this.SetResource(resources.mineNum6)
				case 7:
					this.SetResource(resources.mineNum7)
				case 8:
					this.SetResource(resources.mineNum8)
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
		State.SetEndTime()
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
	markedWrongMines := this.AroundSquares.
		filter(func(s *Square) bool { return s.SquareType == SquareTypeNormal && s.SquareStatus == SquareStatusMarkedFlag })
	if len(markedWrongMines) > 0 {
		markedWrongMines.each(func(s *Square) { s.setStatus(SquareStatusMarkedWrong) })
		this.AroundSquares.
			filter(func(s *Square) bool { return s.SquareType == SquareTypeMine && (s.SquareStatus == SquareStatusClosed || s.SquareStatus == SquareStatusMouseDown) }).
			each(func(s *Square) { s.setStatus(SquareStatusOpened) })
		State.SetEndTime()
		return
	}

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
