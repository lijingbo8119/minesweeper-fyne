package internal

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/gogf/gf/util/gconv"
	"time"
)

type _state struct {
	window fyne.Window

	matrix *Matrix

	mouseState *MouseState

	startTime *time.Time
	endTime   *time.Time
}

func (this *_state) SetWindow(w fyne.Window) {
	this.window = w
}

func (this *_state) SetStartTime(t ...*time.Time) {
	if len(t) > 0 {
		this.startTime = t[0]
	} else {
		_t := time.Now()
		this.startTime = &_t
	}
}

func (this *_state) SetEndTime(t ...*time.Time) {
	if len(t) > 0 {
		this.endTime = t[0]
	} else {
		_t := time.Now()
		this.endTime = &_t
	}
}

func (this *_state) doCheckFinish() {
	minesCount := 0
	markedMinesCount := 0
	markedNormalsCount := 0
	explodedCount := 0

	this.matrix.
		findSquares(func(s *Square) bool { return true }).
		each(func(s *Square) {
			if s.SquareType == SquareTypeMine {
				minesCount++
			}
			if s.SquareStatus == SquareStatusExploded {
				explodedCount++
			}
			if s.SquareStatus != SquareStatusMarkedFlag {
				return
			}
			if s.SquareType == SquareTypeMine {
				markedMinesCount++
			} else {
				markedNormalsCount++
			}
		})

	if explodedCount > 0 {
		this.SetEndTime()
		dialog.ShowError(errors.New("Exploded!"), this.window)
		return
	}

	if minesCount == markedMinesCount && markedNormalsCount == 0 {
		this.matrix.
			findSquares(func(s *Square) bool { return s.SquareStatus == SquareStatusClosed && s.SquareType == SquareTypeNormal }).
			each(func(s *Square) { s.open(true) })
		this.SetEndTime()
		dialog.ShowInformation("congratulations!", "haha~", this.window)
	}
}

func (this *_state) SetMatrixParamAndRender(rowsLength int, colsLength int, minesCount int) {
	this.matrix = NewMatrix(rowsLength, colsLength, minesCount)
	this.SetStartTime(nil)
	this.SetEndTime(nil)
	this.mouseState = NewMouseState().
		RegisterLeftMouseDownHandler(this.leftMouseDownHandler).RegisterLeftMouseUpHandler(this.leftMouseUpHandler).
		RegisterRightMouseDownHandler(this.rightMouseDownHandler).RegisterRightMouseUpHandler(this.rightMouseUpHandler).
		RegisterLeftRightMouseDownHandler(this.leftRightMouseDownHandler).RegisterLeftRightMouseUpHandler(this.leftRightMouseUpHandler).
		RegisterResetHandler(this.resetHandler)

	c0 := NewNumberContainer()
	go func() {
		for {
			time.Sleep(1 * time.Second)
			if this.startTime == nil {
				continue
			}
			end := time.Now()
			c0.SetNumber(gconv.Int(end.Sub(*this.startTime).Seconds()))
			if this.endTime != nil {
				break
			}
		}
	}()

	this.window.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("Game",
			fyne.NewMenuItem("Easy", func() { this.SetMatrixParamAndRender(9, 9, 10) }),
			fyne.NewMenuItem("Intermediate", func() { this.SetMatrixParamAndRender(16, 16, 40) }),
			fyne.NewMenuItem("Expert", func() { this.SetMatrixParamAndRender(16, 30, 99) }),
		),
	))

	b1 := widget.NewButton("Restart", func() {
		this.SetMatrixParamAndRender(rowsLength, colsLength, minesCount)
	})

	c1 := container.New(layout.NewVBoxLayout(), c0.Container, b1)

	c2 := container.New(layout.NewGridLayoutWithColumns(len((*this.matrix)[0])))
	for _, row := range *this.matrix {
		for _, s := range row {
			c2.Add(s)
		}
	}

	this.window.Resize(fyne.Size{Height: 1, Width: 1})
	this.window.SetContent(container.NewVBox(c1, c2))
}

func (this *_state) resetHandler(c Coordinate) {
	this.matrix.
		findSquares(func(square *Square) bool { return square.SquareStatus == SquareStatusMouseDown }).
		each(func(s *Square) { s.setStatus(SquareStatusClosed) })
}

func (this *_state) leftMouseDownHandler(c Coordinate) {
	if this.endTime != nil {
		return
	}
	this.matrix.
		findSquares(func(square *Square) bool { return square.SquareCoordinate.equal(c) && square.SquareStatus == SquareStatusClosed }).
		each(func(s *Square) { s.setStatus(SquareStatusMouseDown) })

	this.doCheckFinish()
}

func (this *_state) leftMouseUpHandler(c Coordinate) {
	if this.endTime != nil {
		return
	}
	if this.startTime == nil {
		this.SetStartTime()
	}

	this.matrix.
		findSquares(func(square *Square) bool { return square.SquareCoordinate.equal(c) && square.SquareStatus == SquareStatusMouseDown }).
		each(func(s *Square) { s.open(true) })

	this.doCheckFinish()
}

func (this *_state) rightMouseDownHandler(c Coordinate) {
	if this.endTime != nil {
		return
	}
}

func (this *_state) rightMouseUpHandler(c Coordinate) {
	if this.endTime != nil {
		return
	}
	if this.startTime == nil {
		this.SetStartTime()
	}

	this.matrix.
		findSquares(func(square *Square) bool { return square.SquareCoordinate.equal(c) }).
		each(func(s *Square) { s.mark() })

	this.doCheckFinish()
}

func (this *_state) leftRightMouseDownHandler(c Coordinate) {
	square := this.matrix.findSquare(func(square *Square) bool { return square.SquareCoordinate.equal(c) })

	switch square.SquareStatus {
	case SquareStatusMouseDown:
		break
	case SquareStatusOpened:
		break
	case SquareStatusClosed:
		return
	default:
		return
	}

	square.AroundSquares.
		filter(func(s *Square) bool {
			return s.SquareStatus == SquareStatusClosed
		}).
		each(func(s *Square) {
			s.setStatus(SquareStatusMouseDown)
		})
}

func (this *_state) leftRightMouseUpHandler(c Coordinate) {
	square := this.matrix.findSquare(func(square *Square) bool { return square.SquareCoordinate.equal(c) })

	switch square.SquareStatus {
	case SquareStatusOpened:
		square.openAroundSquares()
		this.doCheckFinish()
		break
	default:
		return
	}
}

var State = new(_state)
