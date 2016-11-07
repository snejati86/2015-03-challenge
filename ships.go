package main

import (
	"time"
	"math/rand"
)

type ship struct {
	Name string
	Length int
	Score int
	Points map[Point]bool
	Hits int
}

func (ship *ship) AssignPoints(points []Point){
	ship.Points=make(map[Point]bool)
	for _,v:=range points {
		ship.Points[v] = true
	}
}

func ( ship *ship) ProposePoint (point Point) (bool,bool) {
	for p,_:=range ship.Points{
		if ( p == point){
			ship.Hits++
			return true,ship.Hits==ship.Length
		}
	}
	return false,false
}

func GetRandomShip() *ship{
	rand.Seed(time.Now().UTC().UnixNano())
	x:=rand.Intn(len(MODELS))
	return &MODELS[x]
}


var (
	MODELS = []ship{{Name:"Air Carrier",Score:20,Length:5},{Name:"BattleShip",Score:12,Length:4},{Name:"Submarine",Score:6,Length:3},{Name:"Destroyer",Score:6,Length:3},{Name:"Cruiser",Score:6,Length:3},{Name:"Patrol Boat",Score:6,Length:3}}
)