package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Robot struct {
	pos, vel internal.Point[int]
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
			pos: internal.Point[int]{
				X: px,
				Y: py,
			},
			vel: internal.Point[int]{
				X: vx,
				Y: vy,
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
	r.pos.X += r.vel.X
	r.pos.Y += r.vel.Y

	if r.pos.X >= maxX {
		r.pos.X -= maxX
	} else if r.pos.X < 0 {
		r.pos.X += maxX
	}
	if r.pos.Y >= maxY {
		r.pos.Y -= maxY
	} else if r.pos.Y < 0 {
		r.pos.Y += maxY
	}
}

func (r *Robot) getQuadrant() int {

	if r.pos.X < rH {
		if r.pos.Y < uH {
			return 1
		} else if r.pos.Y > uH {
			return 3
		}
	} else if r.pos.X > rH {
		if r.pos.Y < uH {
			return 2
		} else if r.pos.Y > uH {
			return 4
		}
	}
	return 0
}
