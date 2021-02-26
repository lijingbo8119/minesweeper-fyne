package internal

type SquareStatus int

const (
	SquareStatusClosed SquareStatus = iota
	SquareStatusMouseDown
	SquareStatusMarkedFlag
	SquareStatusOpened
	SquareStatusExploded
)

type SquareType int

const (
	SquareTypeNormal SquareType = iota
	SquareTypeMine
)

