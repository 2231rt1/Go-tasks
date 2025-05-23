package main

import (
	"slices"
	"strings"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func (p Player) calculateRating() float64 {
	if p.Misses == 0 {
		return float64(p.Goals) + float64(p.Assists)/2
	}
	return (float64(p.Goals) + float64(p.Assists)/2) / float64(p.Misses)
}

func NewPlayer(name string, goals, misses, assists int) Player {
	player := Player{
		Name:    name,
		Goals:   goals,
		Misses:  misses,
		Assists: assists,
	}
	player.Rating = player.calculateRating()
	return player
}

func goalsSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		if a.Goals > b.Goals {
			return -1
		}
		if a.Goals < b.Goals {
			return 1
		}
		return strings.Compare(a.Name, b.Name)
	})
	return players
}

func ratingSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		if a.Rating > b.Rating {
			return -1
		}
		if a.Rating < b.Rating {
			return 1
		}
		return strings.Compare(a.Name, b.Name)
	})
	return players
}

func gmSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		var ratioA, ratioB float64
		if a.Misses == 0 {
			ratioA = float64(a.Goals)
		} else {
			ratioA = float64(a.Goals) / float64(a.Misses)
		}
		if b.Misses == 0 {
			ratioB = float64(b.Goals)
		} else {
			ratioB = float64(b.Goals) / float64(b.Misses)
		}
		if ratioA > ratioB {
			return -1
		}
		if ratioA < ratioB {
			return 1
		}
		return strings.Compare(a.Name, b.Name)
	})
	return players
}
