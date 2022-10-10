package gotest

import (
	"testing"

	utils "github.com/liuhuo23/utils"
)

func TestAdd(t *testing.T) {
	a := utils.Add(1, 2)
	if a != 3 {
		t.Error("失败！")
	}
}
