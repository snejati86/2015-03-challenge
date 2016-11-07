package main


type PlayerInterface interface {
	ProposePoint(Point) bool
}
type Player struct {
	score int
	turns_played int
	ships []ship
	name string
	pboard *board
	fired [16][16]bool
}

func ( player *Player ) ProposePoint(point Point) (bool,bool){
	for _,ship:=range player.ships {
		if h,s:=ship.ProposePoint(point);h==true{
			return h,s
		}
	}
	return false,false
}

func NewPlayer(name string) *Player{
	p:=&Player{}
	p.pboard=NewBoard()
	p.name=name
	p.fired=[16][16]bool{}
	return p
}