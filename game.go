package main

import (
	"fmt"
	"strconv"
)



type Game struct {
	players []*Player
	parser *InputParser
	MyBoard *board
}

func (game *Game) InitGame(){
	player1 := NewPlayer("Me")
	player2 := NewPlayer("Mo")

	game.players = append(game.players,player1)
	game.players = append(game.players,player2)
	game.parser = &InputParser{}
	for i:=0;i<6;i++{
		ship1:=GetRandomShip()
		ship2:=GetRandomShip()
		player1.pboard.PlaceShip(ship1)
		player1.ships=append(player1.ships,*ship1)
		player2.pboard.PlaceShip(ship2)
		player2.ships=append(player2.ships,*ship2)
	}
}

func ( game *Game) BeginGame(){
	gameEnd := false
	for {
		if ( !gameEnd ){
			for i,player:= range game.players{
				fmt.Println("Player "+strconv.Itoa(i)+" turn : ")
				fmt.Println("===============================================")

				PrintBoard(player.fired)
				fmt.Println("===============================================")

				x,y:=game.parser.ParsePoint()
				// UGLY !
				var otherPlayer *Player
				if (i ==0 ){
					otherPlayer=game.players[1]
				}else{
					otherPlayer=game.players[0]
				}
				PrintBoard(otherPlayer.pboard._board)
				fmt.Println("Firing at position "+strconv.Itoa(x)+","+strconv.Itoa(y))
				if h,s:=otherPlayer.ProposePoint(Point{x,y});h==true{
					player.score++
					player.fired[x][y]=true
					fmt.Println("Player "+strconv.Itoa(i)+" confirmed hit!")
					fmt.Println("Your current score is "+strconv.Itoa(player.score))
					if s == true{
						fmt.Println("BOOM XXXXXXX BOOM")
						fmt.Println("===============================================")
						fmt.Println("SHIP SANK!")
					}
				}else{
					fmt.Println("Player "+strconv.Itoa(i)+" missed")

				}
				player.turns_played++
				if ( player.turns_played == 6 ){
					fmt.Println("Game ending!")
					gameEnd = true
				}
				fmt.Println("===============================================")
			}

		}else{
			break
		}


	}
}
