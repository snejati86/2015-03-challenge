package main

import (
	"testing"
	"fmt"
)

func TestBoardPlacement(*testing.T){
	ship:=&Ship{}
	ship.Length=6
	ship.Name="Dora"

	ship2:=&Ship{}
	ship2.Length=15
	ship2.Name="Dorsa"
	board:=&board{}
	board._board=[16][16]bool{}
	board.PlaceShip(ship)
	board.PlaceShip(ship2)
	fmt.Println(ship.Points)
	PrintBoard(board._board)
}

