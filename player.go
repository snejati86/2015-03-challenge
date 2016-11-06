package main

type PlayerInterface interface {
	ProposePoint(Point) bool
}
type Player struct {
	score int
	turns_played int
	ships []Ship
	name string
	pboard board
}

func ( player *Player ) ProposePoint(point Point) bool{
	var res bool
	for _,ship:=range player.ships {
		if ship.ProposePoint(point){
			res= true
			break
		}
	}
	return res
}