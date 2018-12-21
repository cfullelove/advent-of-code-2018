package main

import (
	"flag"
	"fmt"
)

const GRID_SIZE = 300

var gridSerial int

func main() {

	flag.IntVar(&gridSerial, "serial", 0, "grid serial number")

	flag.Parse()

	grid := Grid(gridSerial)

	// Part 1:
	fmt.Println(grid.MaxPower(GRID_SIZE, 3))

	// 9005 => 20,32

	// Part 2
	fmt.Println(grid.MaxPowerAnySizeParallel(GRID_SIZE))

	// 9005 => 235,287,13

}

type Grid int

func (g Grid) PowerAt(topLeft Point, size int) int {

	serial := int(g)
	totalPower := 0

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			cell := Cell{topLeft.X + i, topLeft.Y + j}
			totalPower += cell.Power(serial)
		}
	}

	return totalPower
}

func (grid Grid) MaxPowerAnySize(gridSize int) (int, Point, int) {
	maxPower := 0
	var maxPowerSize int
	var maxPowerLocation Point

	for size := 0; size < gridSize; size++ {

		fmt.Println(size)
		power, loc := grid.MaxPower(gridSize, size)
		if power > maxPower {
			maxPower = power
			maxPowerLocation = loc
			maxPowerSize = size
		}
	}

	return maxPower, maxPowerLocation, maxPowerSize
}

func (grid Grid) MaxPower(gridSize, blockSize int) (int, Point) {
	maxPower := 0
	var maxPowerLocation Point
	for j := 0; j <= (gridSize - blockSize); j++ {
		for i := 0; i <= (gridSize - blockSize); i++ {
			loc := Point{i, j}
			power := grid.PowerAt(loc, blockSize)
			// fmt.Println(blockSize, loc, power)
			if power > maxPower {
				maxPower = power
				maxPowerLocation = loc
			}
		}
	}

	return maxPower, maxPowerLocation
}

type Cell Point

func (cell Cell) Power(serial int) int {
	rackID := cell.X + 10
	power := rackID * cell.Y
	power = power + serial
	power = power * rackID
	power = Hundreds(power)
	power = power - 5
	return power
}

func Hundreds(num int) int {
	num = num % 1000
	num = num / 100

	return num
}

type Point struct {
	X, Y int
}

type ParallelMaxPowerRangeSize int

func (grid Grid) MaxPowerAnySizeParallel(gridSize int) (int, Point, int) {

	type Result struct {
		Power    int
		Location Point
		Size     int
	}

	maxResult := Result{Power: 0}

	results := make(chan Result)
	allDone := make(chan bool)

	go func() {
		count := 0

		for count < gridSize {
			var result Result
			result = <-results
			if result.Power > maxResult.Power {
				maxResult = result
			}
			count++

			fmt.Printf("%.2f Complete\r", 100*float32(count)/float32(gridSize))
		}
		fmt.Println()

		allDone <- true
	}()

	for size := 0; size < gridSize; size++ {
		go func(_gridSize, _blockSize int) {
			power, location := grid.MaxPower(_gridSize, _blockSize)
			results <- Result{
				Power:    power,
				Location: location,
				Size:     _blockSize,
			}
		}(gridSize, size)
	}

	<-allDone

	return maxResult.Power, maxResult.Location, maxResult.Size
}
