package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Point struct {
	x, y int
}

type Robot struct {
	pos, vel Point
}

const (
	maxX, maxY = 101, 103
	rH, uH     = maxX / 2, maxY / 2
)

var (
	robots []*Robot
)

func main() {
	lines := internal.Reader()

	for _, line := range lines {
		l := strings.Split(strings.ReplaceAll(strings.ReplaceAll(line, "p=", ""), "v=", ""), " ")
		p := strings.Split(l[0], ",")
		px, _ := strconv.Atoi(p[0])
		py, _ := strconv.Atoi(p[1])

		v := strings.Split(l[1], ",")
		vx, _ := strconv.Atoi(v[0])
		vy, _ := strconv.Atoi(v[1])

		robots = append(robots, &Robot{
			pos: Point{
				x: px,
				y: py,
			},
			vel: Point{
				x: vx,
				y: vy,
			},
		})
	}
	// sec, dangerLvl := 0, math.MaxInt
	for sec := 0; sec < 10000; sec++ {
		quads := make(map[int]int)
		for _, robot := range robots {
			robot.move()
			quads[robot.getQuadrant()]++
		}
		if sec == 99 {
			fmt.Printf("Part 1: %d\n", quads[1]*quads[2]*quads[3]*quads[4])
		}
		if sec+1 == 7572 {
			fmt.Println(quads[0] * quads[1] * quads[2] * quads[3] * quads[4])
		}
	}
}

func (r *Robot) move() {
	r.pos.x += r.vel.x
	r.pos.y += r.vel.y

	if r.pos.x >= maxX {
		r.pos.x -= maxX
	} else if r.pos.x < 0 {
		r.pos.x += maxX
	}
	if r.pos.y >= maxY {
		r.pos.y -= maxY
	} else if r.pos.y < 0 {
		r.pos.y += maxY
	}
}

func (r *Robot) getQuadrant() int {

	if r.pos.x < rH {
		if r.pos.y < uH {
			return 1
		} else if r.pos.y > uH {
			return 3
		}
	} else if r.pos.x > rH {
		if r.pos.y < uH {
			return 2
		} else if r.pos.y > uH {
			return 4
		}
	}
	return 0
}
