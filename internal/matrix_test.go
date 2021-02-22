package internal_test

import (
	"fynetest/internal"
	"github.com/gogf/gf/frame/g"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	matrix := internal.NewMatrix(10, 10, 10)
	g.Dump(matrix)
}
