package main

import (
	"fmt"
	"math"
	"sort"
)

type Location struct {
	X int
	Y int
}

type Player struct {
	Location
	Name string
}

func (p Player) distance(x2, y2 int) float64 {
	x := math.Pow(float64(x2-p.X), 2)
	y := math.Pow(float64(y2-p.Y), 2)
	return math.Sqrt(x + y)
}

type PlayerSort struct {
	Players  []Player
	SortFunc func(p1, p2 Player) bool
}

func (ps PlayerSort) Len() int {
	return len(ps.Players)
}

func (ps PlayerSort) Swap(i, j int) {
	ps.Players[i], ps.Players[j] = ps.Players[j], ps.Players[i]
}

func (ps PlayerSort) Less(i, j int) bool {
	return ps.SortFunc(ps.Players[i], ps.Players[j])
}

func sortByDistance(players []Player, x, y int) {
	ps := PlayerSort{Players: players,
		SortFunc: func(p1 Player, p2 Player) bool {
			return p1.distance(x, y) < p2.distance(x, y)
		}}
	sort.Sort(ps)
}

func main() {
	ps := []Player{
		{Name: "1", Location: Location{10, 60}},
		{Name: "2", Location: Location{20, 70}},
		{Name: "3", Location: Location{30, 80}},
		{Name: "4", Location: Location{40, 90}},
	}
	fmt.Printf("%+v\n", ps)
	sortByDistance(ps, 100, 200)
	fmt.Printf("%+v\n", ps)
}
