package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Game struct {
    Id int
    Bag map[rune]int
}

func NewGame(id int) Game {
    bag := make(map[rune]int, 3)
    bag['r'] = 0
    bag['g'] = 0
    bag['b'] = 0
    return Game{ id, bag }
}

func main() {
    document := read("./input.txt")
    games := parse(document)
    ids := 0
    for _, game := range games {
        if game.Bag['r'] <= 12 && game.Bag['g'] <= 13 && game.Bag['b'] <= 14 {
            ids += game.Id
        }
    }
    fmt.Printf("1. Result: %d\n", ids)
}

func read(filename string) []string {
    file, err := os.Open(filename)
    if err != nil {
        panic("Could not open file")
    }
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    index := 0
    result := make([]string, 1024)
    for scanner.Scan() {
        result[index] = scanner.Text()
        index++
    }
    return result
}

func parse(s []string) []Game {
    games := make([]Game, len(s))
    for i, line := range s {
        if len(line) == 0 {
            break
        }

        var id int
        fmt.Sscanf(line, "Game %d:", &id)
        bag := strings.Split(line, ":")[1]
        games[i] = NewGame(id)
        for _, sets := range strings.Split(bag, ";") {
            for _, set := range strings.Split(sets, ",") {
                var count int
                var color rune
                fmt.Sscanf(set, "%d %c", &count, &color)
                if games[i].Bag[color] < count {
                    games[i].Bag[color] = count
                }
            }
        }
    }
    return games
}
