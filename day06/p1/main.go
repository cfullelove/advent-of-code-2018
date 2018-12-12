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

	_, area := largestArea(places)
	fmt.Println(area)
}

func findClosestPlace(places []pos, p pos) (pos, bool) {
	minDist := math.MaxInt16
	var minPlace pos
	multiple := false

	for _, place := range places {
		d := p.dist(place)

		if d == minDist {
			multiple = true
		}
		if d < minDist {
			minDist = d
			minPlace = place
			multiple = false
		}
	}

	return minPlace, !multiple
}

func largestArea(places []pos) (pos, int) {

	maxRadius := 500
	var maxLocation pos
	maxCount := 0

	for _, loc := range places {
		count := 0
		r := 1
		isInfinite := false
		for {
			if r >= maxRadius {
				isInfinite = true
				break
			}

			locInList := false
			for _, place := range loc.circle(r) {
				closest, ok := findClosestPlace(places, place)
				if !ok {
					continue
				}

				if loc == closest {
					locInList = true
					count++
				}
			}

			if !locInList {
				break
			}
			r++
		}

		if count > maxCount && !isInfinite {
			maxCount = count
			maxLocation = loc
		}

	}

	return maxLocation, maxCount + 1
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
