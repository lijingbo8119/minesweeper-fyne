package internal

import "errors"

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

var (
	ErrorStatusError = errors.New("状态错误")
)
