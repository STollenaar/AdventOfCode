package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	Number int  `json:"Number"`
	Marked bool `json:"Marked"`
}

type Board struct {
	Grid  [5][5]*Position `json:"Position"`
	Score int             `json:"Score"`
	Bingo bool            `json:"Bingo"`
}

type LineOc struct {
	Occurences int `json:"Occurences"`
}

type Bingo struct {
	Board      *Board `json:"Board"`
	LastNumber int    `json:"LastNumber"`
}

var boards []*Board
var bingos []*Bingo

func hasBingo(board *Board) bool {
	blockY := make([]LineOc, 5)
	for _, row := range board.Grid {
		blockX := LineOc{}
		for x, column := range row {
			if column.Marked {
				blockX.Occurences++
				blockY[x].Occurences++
				if blockY[x].Occurences == 5 {
					return true
				}
			}
		}
		if blockX.Occurences == 5 {
			return true
		}
	}
	return false
}

func checkGrid(board *Board, number int) bool {
	for _, row := range board.Grid {
		for _, pos := range row {
			if pos.Number == number {
				pos.Marked = true
				return true
			}
		}
	}
	return false
}

func getFirstWin(boards []*Board, numbers []int) (*Board, int) {
	for _, called := range numbers {
		for _, board := range boards {
			if checkGrid(board, called) {
				board.Score -= called
				if hasBingo(board) {
					return board, called
				}
			}
		}
	}
	return &Board{}, 0
}

func getLastWin(boards []*Board, numbers []int) (*Board, int) {
	for _, called := range numbers {
		for _, board := range boards {
			if !board.Bingo && checkGrid(board, called) {
				board.Score -= called
				if hasBingo(board) {
					bingos = append(bingos, &Bingo{Board: board, LastNumber: called})
					board.Bingo = true
				}
			}
		}
	}
	lastBingo := bingos[len(bingos)-1]
	return lastBingo.Board, lastBingo.LastNumber
}

func filterEmpty(slice []string) (filtered []string) {
	for _, val := range slice {
		if val != "" {
			filtered = append(filtered, val)
		}
	}
	return filtered
}

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var numbers []int
	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ",") {
			for _, number := range strings.Split(line, ",") {
				t, _ := strconv.Atoi(number)
				numbers = append(numbers, t)
			}
		} else if line == "" {
			boards = append(boards, &Board{})
			index = 0
		} else {
			numbers := filterEmpty(strings.Split(line, " "))
			for i, number := range numbers {
				t, _ := strconv.Atoi(number)
				position := &Position{t, false}
				boards[len(boards)-1].Grid[index][i] = position
				boards[len(boards)-1].Score += t
			}
			index++
		}
	}

	origJSON, err := json.Marshal(boards)

	if err != nil {
		log.Fatal(err)
	}

	var clone []*Board
	if err = json.Unmarshal(origJSON, &clone); err != nil {
		log.Fatal(err)
	}

	firstWin, finalNumber := getFirstWin(clone, numbers)
	fmt.Println(firstWin.Score * finalNumber)

	lastWin, lastNumber := getLastWin(boards, numbers)
	fmt.Println(lastWin.Score * lastNumber)
}
