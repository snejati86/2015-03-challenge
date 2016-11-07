package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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
			if board[i][j]{
				fmt.Print("X")
			}else{
				fmt.Print(".")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println("===================================")
}


func ( b *board) PlaceShip(ship *ship) error {

	visited:=make(map[*Point]bool)
	original:=b._board
	points:=[]Point{}
	for {
		if  len(visited) == 16*16{
			return errors.New("Can not place the ship.")
		}
		rem:=ship.Length
		//placed:=false
		//Pick a point
		rand.Seed(time.Now().UTC().UnixNano())
		x:=rand.Intn(16)
		y:=rand.Intn(16)
		visited[&Point{x,y}] = true
		cx:=x
		cy:=y
		found:=false


		for cy >=0  && !b._board[cx][cy] {
			b._board[cx][cy]= true
			if rem == 0 {
				found = true
				break
			}
			cy--
			rem--
			points=append(points,Point{cx,cy})


		}
		if found {
			break
		}else{
			b._board=original
			cy=y
			rem=ship.Length
			points=[]Point{}
		}

		for cy <16  && !b._board[cx][cy] {
			b._board[cx][cy]= true
			if rem == 0 {
				found = true
				break
			}
			cy++
			rem--

			points=append(points,Point{cx,cy})

		}
		if found {
			break
		}else{
			b._board=original
			cy=y
			cx=x
			rem=ship.Length
			points=[]Point{}
		}


		for cx >= 0  && !b._board[cx][cy] {
			b._board[cx][cy]= true
			if rem == 0 {
				found = true
				break
			}

			cx--
			rem--
			points=append(points,Point{cx,cy})


		}
		if found {
			break
		}else{
			b._board=original
			cx=x
			rem=ship.Length
			points=[]Point{}
		}

		for cx <16  && !b._board[cx][cy] {
			b._board[cx][cy]= true
			if rem == 0 {
				found = true
				break
			}


			cx++
			rem--
			points=append(points,Point{cx,cy})

		}
		if found {
			break
		}else{
			points=[]Point{}
			b._board=original

		}



	}
	ship.AssignPoints(points)
	return nil
}


func NewBoard() *board{
	b:=&board{}
	b._board=[16][16]bool{}
	return b

}
