package main

import (
	"fmt"
	"sort"
	"strings"
)

type player struct {
	name  string
	score int32
}

func main() {
	playerMap := map[string]int32{
		"amy":      100,
		"david":    100,
		"heraldo":  50,
		"aakansha": 75,
		"aleksa":   150,
	}

	players := []player{}
	for k, v := range playerMap {
		players = append(players, player{k, v})
	}

	sort.Slice(players, func(i, j int) bool {
		current := players[j]
		next := players[i]
		if current.score == next.score {
			return strings.Compare(current.name, next.name) > 0
		} else {
			return current.score < next.score
		}
	})
	fmt.Println(players)
}
