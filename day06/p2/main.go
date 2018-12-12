package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	places := []pos{}

	s := bufio.NewScanner(f)
	for s.Scan() {

		var p pos
		_, err := fmt.Sscanf(s.Text(), "%d, %d", &p.X, &p.Y)
		if err != nil {
			log.Fatal(err)
		}

		places = append(places, p)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(SafeRegionSize(places, 10000))
}

func SafeRegionSize(places []pos, maxSum int) int {
	topLeft, bottomRight := Bounds(places)

	size := 0
	for i := topLeft.X; i <= bottomRight.X; i++ {
		for j := topLeft.Y; j <= bottomRight.Y; j++ {
			distanceSum := 0
			for _, loc := range places {
				distanceSum = distanceSum + loc.dist(pos{i, j})
			}
			if distanceSum < maxSum {
				size++
			}
		}
	}
	return size
}

func Bounds(places []pos) (pos, pos) {
	minX := math.MaxInt16
	minY := math.MaxInt16
	maxX := math.MinInt16
	maxY := math.MinInt16

	for _, p := range places {
		if p.X < minX {
			minX = p.X
		}

		if p.X > maxX {
			maxX = p.X
		}

		if p.Y < minY {
			minY = p.Y
		}

		if p.Y > maxY {
			maxY = p.Y
		}
	}

	return pos{minX, minY}, pos{maxX, maxY}
}

type pos struct{ X, Y int }

func (p pos) dist(p2 pos) int {
	dist := Abs(p2.X-p.X) + Abs(p2.Y-p.Y)
	return dist
}

func (p pos) circle(r int) []pos {
	if r < 1 {
		return []pos{}
	}

	ring := []pos{}

	for x := -r; x <= r; x++ {
		y := r - Abs(x)
		ring = append(ring, pos{p.X + x, p.Y + y})
		if y != 0 {
			ring = append(ring, pos{p.X + x, p.Y - y})
		}
	}

	return ring

}

func Abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
