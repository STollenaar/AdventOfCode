package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Player struct {
	place, score int
}

var diceOutcomes map[int]int = map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}

func (p Player) doMove(roll int) Player {
	p.place += roll
	for p.place > 10 {
		p.place -= 10
	}
	p.score += p.place
	return p
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var playerPos []int
	for scanner.Scan() {
		line := scanner.Text()
		pos := strings.Split(line, " starting position: ")
		posN, _ := strconv.Atoi(pos[1])
		playerPos = append(playerPos, posN)
		if err != nil {
			log.Fatal(err)
		}
	}
	player1 := Player{place: playerPos[0]}
	player2 := Player{place: playerPos[1]}

	doPart1(player1, player2, start)
	doPart2(player1, player2, start)
}

func doPart1(player1, player2 Player, start time.Time) {
	dice := 1
	diceRolls := 0
	var losingPlayerScore int
	player1Turn := true
	for player1.score < 1000 && player2.score < 1000 {
		var roll int
		for i := 0; i <= 2; i++ {
			roll += dice
			dice++
			diceRolls++
			if dice > 100 {
				dice = 1
			}
		}
		if player1Turn {
			player1 = player1.doMove(roll)
		} else {
			player2 = player2.doMove(roll)
		}
		player1Turn = !player1Turn
	}
	if player1.score >= 1000 {
		losingPlayerScore = player2.score
	} else {
		losingPlayerScore = player1.score
	}

	totalScore := losingPlayerScore * diceRolls
	elapsed := time.Since(start)
	fmt.Println("Execution time for part 1: ", elapsed)
	fmt.Println("Total score for part 1: ", totalScore)
}

func doUniverse(current, other Player, player1Turn bool, universes int) (player1Wins, player2Wins int) {
	if other.score >= 21 {
		if player1Turn {
			return 0, universes
		} else {
			return universes, 0
		}
	}

	for roll, expansions := range diceOutcomes {
		player1Win, player2Win := doUniverse(other, current.doMove(roll), !player1Turn, universes*expansions)
		player1Wins += player1Win
		player2Wins += player2Win
	}
	return player1Wins, player2Wins
}

func doPart2(player1, player2 Player, start time.Time) {
	player1Winnings, player2Winnings := doUniverse(player1, player2, true, 1)

	var bigWinner int
	if player1Winnings > player2Winnings {
		bigWinner = player1Winnings
	} else {
		bigWinner = player2Winnings
	}

	elapsed := time.Since(start)
	fmt.Println("Execution time for part 2: ", elapsed)
	fmt.Println("Total score for part 2: ", bigWinner)
}
