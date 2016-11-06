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
	player1 := &Player{}
	player2 := &Player{}
	game.players = append(game.players,player1)
	game.players = append(game.players,player2)
	game.parser = &InputParser{}
}

func ( game *Game) BeginGame(){
	gameEnd := false
	for {
		if ( !gameEnd ){
			for i,player:= range game.players{
				fmt.Println("Player "+strconv.Itoa(i)+" turn : ")
				x,y:=game.parser.ParsePoint()
				var otherPlayer *Player
				// UGLY !
				if (i ==0 ){
					otherPlayer=game.players[1]
				}else{
					otherPlayer=game.players[0]
				}
				fmt.Println("Firing at position "+strconv.Itoa(x)+","+strconv.Itoa(y))
				if otherPlayer.ProposePoint(Point{x,y}){
					player.score++
					fmt.Println("Player "+strconv.Itoa(i)+" confirmed hit!")
					fmt.Println("Your current score is "+strconv.Itoa(player.score))
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
