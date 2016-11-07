package main

import (
	"testing"
)

func TestBoardPlacement(*testing.T){
	p:=NewPlayer("Sina")

	for i:=0;i<6;i++{
		p.pboard.PlaceShip(GetRandomShip())
	}

/*	ship:=GetRandomShip()
	board.PlaceShip(ship)
	fmt.Println(ship.Points)
	PrintBoard(board._board)*/
}

