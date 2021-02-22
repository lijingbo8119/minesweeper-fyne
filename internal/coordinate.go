package internal

import (
	"math"
)

type Coordinate struct {
	I int
	J int
}

func (this Coordinate) equal(c Coordinate) bool {
	return !this.isEmpty() && this.I == c.I && this.J == c.J
}

func (this Coordinate) near(c Coordinate) bool {
	if this.equal(c) {
		return false
	}
	return int(math.Abs(float64(this.I-c.I))) <= 1 && int(math.Abs(float64(this.J-c.J))) <= 1
}

func (this Coordinate) isEmpty() bool {
	return this.I == -1 && this.J == -1
}

func newCoordinate(i int, j int) Coordinate {
	return Coordinate{i, j}
}

func newEmptyCoordinate() Coordinate {
	return Coordinate{-1, -1}
}
