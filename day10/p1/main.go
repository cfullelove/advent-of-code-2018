package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
)

var sky = &Sky{}

func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	// s.Split(bufio.ScanWords)
	for s.Scan() {
		// fmt.Println(s.Text())
		var posX, posY, velX, velY int
		_, err := fmt.Sscanf(s.Text(), "position=<%d,%d> velocity=<%d,%d>", &posX, &posY, &velX, &velY)
		if err != nil {
			log.Fatal(err)
		}

		sky.Stars = append(sky.Stars, &Star{
			Pos: Point{posX, posY},
			Vx:  velX,
			Vy:  velY,
		})

	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	t := 0
	prevArea := math.MaxInt64

	for {
		area := sky.Area()
		if area > prevArea {
			sky.Tick(-1)
			fmt.Println(sky)
			return
		}
		prevArea = area

		sky.Tick(1)
		t++
	}

	// Part 1: ERKECKJJ
	// Part 2: 10645

}

type Sky struct {
	Stars []*Star
}

func (sky *Sky) Area() int {

	p1, p2 := sky.Bounds()

	return (p2.X - p1.X) * (p2.Y - p1.Y)
}

func (sky *Sky) Bounds() (P1, P2 Point) {
	// P1 is top left (so minX minY)
	P1.X = math.MaxInt16
	P1.Y = math.MaxInt16

	// P2 is bottom right (so maxX maxY)
	P2.Y = math.MinInt16
	P2.X = math.MinInt16

	for _, star := range sky.Stars {
		if star.Pos.X < P1.X {
			P1.X = star.Pos.X
		}

		if star.Pos.Y < P1.Y {
			P1.Y = star.Pos.Y
		}

		if star.Pos.X > P2.X {
			P2.X = star.Pos.X
		}

		if star.Pos.Y > P2.Y {
			P2.Y = star.Pos.Y
		}

	}

	return P1, P2
}

func (sky *Sky) Tick(dT int) {
	for _, star := range sky.Stars {
		star.Tick(dT)
	}
}

func (sky *Sky) String() string {

	w := &bytes.Buffer{}

	grid := map[Point]bool{}

	for _, star := range sky.Stars {
		grid[star.Pos] = true
	}

	p1, p2 := sky.Bounds()

	for j := p1.Y; j <= p2.Y; j++ {
		fmt.Fprint(w, "| ")
		for i := p1.X; i <= p2.X; i++ {
			if _, ok := grid[Point{i, j}]; ok {
				fmt.Fprintf(w, "><")
			} else {
				fmt.Fprintf(w, "  ")
			}
		}
		fmt.Fprintln(w, " |")
	}

	return w.String()

}

type Star struct {
	Pos    Point // Position
	Vx, Vy int   //Velocity
}

func (star *Star) Tick(dT int) {
	star.Pos.X += dT * star.Vx
	star.Pos.Y += dT * star.Vy
}

type Point struct {
	X, Y int
}
