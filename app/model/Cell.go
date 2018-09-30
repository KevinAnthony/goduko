package model

import "github.com/fatih/color"

type Cell struct {
	set    uint8
	pen    uint8
	pencil [9]uint8
}

func (c Cell) GetValue() interface{} {
	if c.set != 0 {
		return color.New(color.FgGreen).SprintFunc()(c.set)
	}
	if c.pen != 0 {
		return color.New(color.FgRed).SprintFunc()(c.pen)
	}
	return color.New(color.FgWhite).SprintfFunc()(" ")
}
func (c Cell) Init(value uint8) {
	c.set = value
}

func (c Cell) Write(value uint8) {
	c.pen = value
}
