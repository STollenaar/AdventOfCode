package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

const (
	cOrder  = "AKQJT98765432"
	cJOrder = "AKQT98765432J"
)

var (
	jokerActive = false
)

type HandType string

type Hand struct {
	cards string
	bid   int
}

func (h *Hand) getHandType() (HandType, int) {
	var pairHold, threeHold int
	for i, l := range h.cards {
		switch strings.Count(h.cards, string(l)) {
		case 5:
			return "Five", 0
		case 4:
			if jokerActive {
				if string(l) != "J" {
					if strings.Count(h.cards, "J") > 0 {
						return "Five", 0
					}
				} else {
					return "Five", 0
				}
			}
			return "Four", 1
		case 3:
			if jokerActive {
				if string(l) != "J" {
					if strings.Count(h.cards, "J") == 2 {
						return "Five", 0
					} else if strings.Count(h.cards, "J") == 1 {
						return "Four", 1
					}
				} else {
					temp := strings.ReplaceAll(h.cards, "J", "")
					if temp[0] == temp[1] {
						return "Five", 0
					} else {
						return "Four", 1
					}
				}
			}
			if i == strings.Index(h.cards, string(l)) {
				threeHold++
			}
		case 2:
			if jokerActive {
				if string(l) != "J" {
					if strings.Count(h.cards, "J") == 3 {
						return "Five", 0
					} else if strings.Count(h.cards, "J") == 2 {
						return "Four", 1
					} else if strings.Count(h.cards, "J") == 1 && i == strings.Index(h.cards, string(l)) {
						threeHold++
					} else if i == strings.Index(h.cards, string(l)) {
						pairHold++
					}
				} else {
					temp := strings.ReplaceAll(h.cards, "J", "")
					tempHand := Hand{temp, 0}
					_, tempHandTypeIndex := tempHand.getHandType()
					// fmt.Println(h.cards, tempHand, tempHandType)
					switch tempHandTypeIndex {
					case 3:
						return "Five", 0
					case 1:
						return "Four", 1
					default:
						if i == strings.Index(h.cards, string(l)) {
							threeHold++
						}
					}
				}
			} else {
				if i == strings.Index(h.cards, string(l)) {
					pairHold++
				}
			}
		}
	}
	if (threeHold == 1 && pairHold == 1) || threeHold == 2 {
		return "Full", 2
	} else if threeHold == 1 {
		return "Three", 3
	} else if pairHold == 2 {
		return "Two", 4
	} else if pairHold == 1 {
		return "One", 5
	} else if pairHold == 0 && threeHold == 0 && strings.Count(h.cards, "J") == 1 {
		return "One", 5
	}

	return "High", 6
}

func isWinningHand(a, op Hand) int {
	_, hIndex := a.getHandType()
	_, opIndex := op.getHandType()

	if hIndex < opIndex {
		return 1
	} else if hIndex > opIndex {
		return -1
	}
	for i, card := range a.cards {
		if string(card) != string(op.cards[i]) {
			var cIndex, opCIndex int
			if jokerActive {
				cIndex = strings.Index(cJOrder, string(card))
				opCIndex = strings.Index(cJOrder, string(op.cards[i]))
			} else {
				cIndex = strings.Index(cOrder, string(card))
				opCIndex = strings.Index(cOrder, string(op.cards[i]))
			}
			if cIndex < opCIndex {
				return 1
			} else if cIndex > opCIndex {
				return -1
			} else {
				return 0
			}
		}
	}
	panic("This shouldn't happen")
}

func main() {
	lines := internal.Reader()

	var hands []Hand
	for _, line := range lines {
		hand, bid := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
		b, _ := strconv.Atoi(strings.TrimSpace(bid))
		hands = append(hands, Hand{cards: hand, bid: b})
	}
	slices.SortFunc[[]Hand](hands, isWinningHand)

	var total int
	for i, hand := range hands {
		total += (i + 1) * hand.bid
	}
	fmt.Printf("Solution part 1: %d\n", total)

	jokerActive = true
	slices.SortFunc[[]Hand](hands, isWinningHand)

	total = 0
	for i, hand := range hands {
		total += (i + 1) * hand.bid
	}
	fmt.Printf("Solution part 2: %d\n", total)
}
