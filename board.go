package main

import (
	"errors"
	"fmt"
	"strconv"
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
func ( b *board) PlaceShip(ship *Ship) error {

	visited:=make(map[*Point]bool)
	original:=b._board
	points:=[]Point{}
	for {
		fmt.Println("Getting the next point")
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
		fmt.Println("search at point "+strconv.Itoa(cx)+","+strconv.Itoa(cy))
		fmt.Println("remiaing "+strconv.Itoa(rem))


		for cy >=0  && !b._board[cx][cy] {
			b._board[cx][cy]= true
			if rem == 0 {
				found = true
				break
			}
			cy--
			rem--
			points=append(points,Point{cx,cy})
			fmt.Println("Place at point "+strconv.Itoa(cx)+","+strconv.Itoa(cy))


		}
		if found {
			break
		}else{
			b._board=original
			cy=y
			rem=ship.Length
			points=[]Point{}
		}
		fmt.Println("search at point "+strconv.Itoa(cx)+","+strconv.Itoa(cy))
		fmt.Println("remiaing "+strconv.Itoa(rem))

		for cy <16  && !b._board[cx][cy] {
			b._board[cx][cy]= true
			if rem == 0 {
				found = true
				break
			}
			cy++
			rem--
			fmt.Println("Place at point "+strconv.Itoa(cx)+","+strconv.Itoa(cy))

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

		fmt.Println("search at point "+strconv.Itoa(cx)+","+strconv.Itoa(cy))
		fmt.Println("remiaing "+strconv.Itoa(rem))

		for cx >= 0  && !b._board[cx][cy] {
			b._board[cx][cy]= true
			if rem == 0 {
				found = true
				break
			}

			cx--
			rem--
			fmt.Println("Place at point "+strconv.Itoa(cx)+","+strconv.Itoa(cy))
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
		fmt.Println("search at point "+strconv.Itoa(cx)+","+strconv.Itoa(cy))
		fmt.Println("remiaing "+strconv.Itoa(rem))

		for cx <16  && !b._board[cx][cy] {
			b._board[cx][cy]= true
			if rem == 0 {
				found = true
				break
			}


			cx++
			rem--
			fmt.Println("Place at point "+strconv.Itoa(cx)+","+strconv.Itoa(cy))
			points=append(points,Point{cx,cy})

		}
		if found {
			break
		}else{
			points=[]Point{}
			b._board=original

		}



	}
	fmt.Println("We are here")
	ship.AssignPoints(points)
	return nil
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

