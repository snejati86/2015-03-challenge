package main

type Ship struct {
	Name string
	Length int
	Score int
	Points map[Point]bool
	Hits int
}

func (ship *Ship) AssignPoints(points []Point){
	ship.Points=make(map[Point]bool)
	for _,v:=range points {
		ship.Points[v] = true
	}
}

func ( ship *Ship) ProposePoint (point Point) bool {
	for p,_:=range ship.Points{
		if ( p == point){
			ship.Hits++
			return true
		}
	}
	return false
}