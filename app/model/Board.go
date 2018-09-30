package model

import (
	"fmt"
)

type Board struct {
	Cells  [9][9]*Cell
	rows   [9]Row
	column [9]Column
	block  [9]Block
}

func (b *Board) Init() {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			b.Cells[r][c] = new(Cell)
			b.Cells[r][c].Init(1)
		}
	}
}

func (b *Board) Output() {
	for r := 0; r < 9; r++ {
		fmt.Println("+-----+-----+-----+-----+-----+-----+-----+-----+-----+")
		fmt.Println("|     |     |     |     |     |     |     |     |     |")
		for c := 0; c < 9; c++ {
			fmt.Printf("|%3v  ", b.Cells[r][c].GetValue())
		}
		fmt.Println("|")
		fmt.Println("|     |     |     |     |     |     |     |     |     |")
	}
	fmt.Println("+-----+-----+-----+-----+-----+-----+-----+-----+-----+")

}
