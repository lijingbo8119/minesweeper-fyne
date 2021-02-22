package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/gogf/gf/util/gconv"
	"time"
)

type _state struct {
	window fyne.Window

	matrix *Matrix

	leftMouseDownCoordinate  Coordinate
	leftMouseUpCoordinate    Coordinate
	rightMouseDownCoordinate Coordinate
	rightMouseUpCoordinate   Coordinate

	startTime *time.Time
	endTime   *time.Time

	durationSeconds binding.String
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
	if this.endTime != nil {
		this.durationSeconds.Set(gconv.String(gconv.Int64(this.endTime.Sub(*this.startTime).Seconds())))
	}
}

func (this *_state) doCheck() {
	minesCount := 0
	markedMinesCount := 0
	markedNormalsCount := 0

	this.matrix.
		findSquares(func(s *Square) bool { return true }).
		each(func(s *Square) {
			if s.SquareType == SquareTypeMine {
				minesCount++
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

	if minesCount == markedMinesCount && markedNormalsCount == 0 {
		this.SetEndTime()
	}
}

func (this *_state) SetMatrixParamAndRender(rowsLength int, colsLength int, minesCount int) {
	this.matrix = NewMatrix(rowsLength, colsLength, minesCount)
	this.SetStartTime(nil)
	this.SetEndTime(nil)
	this.resetCoordinates()

	this.durationSeconds = binding.NewString()
	label := widget.NewLabelWithData(this.durationSeconds)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			if this.startTime == nil {
				continue
			}
			end := time.Now()
			this.durationSeconds.Set(gconv.String(gconv.Int64(end.Sub(*this.startTime).Seconds())))
			if this.endTime != nil {
				break
			}
		}
	}()
	b1 := widget.NewButton("Restart", func() {
		this.SetMatrixParamAndRender(rowsLength, colsLength, minesCount)
	})
	b2 := widget.NewButton("Easy", func() {
		this.SetMatrixParamAndRender(9, 9, 10)
	})
	b3 := widget.NewButton("Intermediate", func() {
		this.SetMatrixParamAndRender(16, 16, 40)
	})
	b4 := widget.NewButton("Expert", func() {
		this.SetMatrixParamAndRender(16, 30, 99)
	})
	c1 := container.New(layout.NewVBoxLayout(), label, b1, b2, b3, b4)

	c2 := container.New(layout.NewGridLayoutWithColumns(len((*this.matrix)[0])))
	for _, row := range *this.matrix {
		for _, s := range row {
			c2.Add(s)
		}
	}
	this.window.Resize(fyne.Size{Height: 1, Width: 1})
	this.window.SetContent(container.NewVBox(c1, c2))
}

func (this *_state) resetCoordinates() {
	this.leftMouseDownCoordinate = newEmptyCoordinate()
	this.leftMouseUpCoordinate = newEmptyCoordinate()
	this.rightMouseDownCoordinate = newEmptyCoordinate()
	this.rightMouseUpCoordinate = newEmptyCoordinate()
	mouseDownSquares := this.matrix.findSquares(func(square *Square) bool { return square.SquareStatus == SquareStatusMouseDown })
	for _, s := range mouseDownSquares {
		s.setStatus(SquareStatusClosed)
	}
}

func (this *_state) leftMouseDown(c Coordinate) {
	if this.endTime != nil {
		return
	}
	this.leftMouseDownCoordinate = c
	square := this.matrix.findSquare(func(square *Square) bool { return square.SquareCoordinate.equal(c) })
	if square.SquareStatus == SquareStatusClosed {
		square.setStatus(SquareStatusMouseDown)
	}
}

func (this *_state) leftMouseUp(c Coordinate) {
	if this.endTime != nil {
		return
	}
	this.leftMouseUpCoordinate = c
	if !this.leftMouseUpCoordinate.equal(this.leftMouseDownCoordinate) {
		this.resetCoordinates()
		return
	}
	if !this.rightMouseDownCoordinate.isEmpty() {
		if !this.rightMouseDownCoordinate.equal(this.leftMouseDownCoordinate) {
			this.resetCoordinates()
			return
		}
		if this.rightMouseUpCoordinate.equal(this.leftMouseUpCoordinate) {
			this.leftRightClick(c)
			this.resetCoordinates()
			return
		}
		if this.rightMouseDownCoordinate.equal(this.leftMouseUpCoordinate) {
			return
		}
	}

	this.leftClick(c)
	this.resetCoordinates()
}

func (this *_state) rightMouseDown(c Coordinate) {
	if this.endTime != nil {
		return
	}
	this.rightMouseDownCoordinate = c
}

func (this *_state) rightMouseUp(c Coordinate) {
	if this.endTime != nil {
		return
	}
	this.rightMouseUpCoordinate = c
	if !this.rightMouseUpCoordinate.equal(this.rightMouseDownCoordinate) {
		this.resetCoordinates()
		return
	}
	if !this.leftMouseDownCoordinate.isEmpty() {
		if !this.leftMouseDownCoordinate.equal(this.leftMouseDownCoordinate) {
			this.resetCoordinates()
			return
		}
		if this.leftMouseUpCoordinate.equal(this.rightMouseUpCoordinate) {
			this.leftRightClick(c)
			this.resetCoordinates()
			return
		}
		if this.leftMouseDownCoordinate.equal(this.rightMouseUpCoordinate) {
			return
		}
	}
	this.rightClick(c)
	this.resetCoordinates()
}

func (this *_state) leftClick(c Coordinate) {
	if this.startTime == nil {
		this.SetStartTime()
	}
	this.matrix.findSquare(func(square *Square) bool { return square.SquareCoordinate.equal(c) }).open(true)
	this.doCheck()
}

func (this *_state) rightClick(c Coordinate) {
	this.matrix.findSquare(func(square *Square) bool { return square.SquareCoordinate.equal(c) }).mark()
	if this.startTime == nil {
		t := time.Now()
		this.startTime = &t
	}
	this.doCheck()
}

func (this *_state) leftRightClick(c Coordinate) {
	square := this.matrix.findSquare(func(square *Square) bool { return square.SquareCoordinate.equal(c) })
	if square.SquareStatus != SquareStatusOpened {
		return
	}
	square.openAroundSquares()
	this.doCheck()
}

var State = new(_state)
