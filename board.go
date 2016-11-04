package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Point struct {
	x int
	y int
}

type board struct {
	_board [16][16]bool
}
func PrintBoard(board [16][16]bool ) {
	for i := 0; i < 16; i++ {
		for j:=0;j<16;j++{
			fmt.Print(board[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println("===================================")
}
func ( b *board) PlaceShip(ship *Ship) error {
	rem:=ship.Length
	placed:=false
	for i := 0; i < 16; i++ {
		if ( !placed  ){
			for j:=0 ; j < 16 ; j++{
				if !b._board[i][j]{
					res:= initSearch(&b._board,rem,i,j,ship)
					if ( res ){
						return nil
					}
				}
				if ( rem == 0 ){
					placed = true
					break
				}
			}
		}else{
			break
		}
	}
	return errors.New("Not found.")
}

func initSearch(board *[16][16]bool,togo int,x int,y int,ship *Ship) bool {
	fmt.Println("we have "+strconv.Itoa(togo)+ " to go in "+strconv.Itoa(x)+" and "+strconv.Itoa(y))
	var points []Point
	var left = togo
	//Check left
	for  i:=x ; i < len(board[0]) ; i++{
		if  !board[i][y]{
			left--
			board[i][y] = true
			points=append(points,Point{i,y})
			if  left == 0{
				ship.AssignPoints(points)
				return true
			}
			continue

		}else{
			left = togo
			points=nil
		}
	}

	for  i:=y ; i < len(board) ; i--{
		if  !board[x][i]{
			left--
			board[x][i] = true
			points=append(points,Point{x,i})
			if  left == ship.Length{
				ship.AssignPoints(points)
				return true
			}
			continue

		}else{
			left = togo
			points=nil
		}
	}

	return false


	//Check right



}
