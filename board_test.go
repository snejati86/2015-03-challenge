package main

import (
	"testing"
)

func TestBoardPlacement(*testing.T){
	ship:=&Ship{}
	ship.Length=14
	ship.Name="Dora"
	board:=&board{}
	board._board=[16][16]bool{}
	board.PlaceShip(ship)

	PrintBoard(board._board)
}

