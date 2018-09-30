package main

import "github.com/KevinAnthony/godoku/app/model"

func main() {
	board := model.Board{}
	board.Init()
	board.Output()
}
