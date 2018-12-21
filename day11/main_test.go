package main

import (
	"fmt"
	"testing"
)

func TestHundreds(t *testing.T) {
	cases := [][2]int{
		{100, 1},
		{1234, 2},
		{10, 0},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%d", c[0]), func(t *testing.T) {
			if Hundreds(c[0]) != c[1] {
				t.Errorf("Expected %d, got %d", c[1], Hundreds(c[0]))
			}
		})
	}
}

func TestCellPower(t *testing.T) {

	type testCase struct {
		Cell   Cell
		Serial int
		Power  int
	}

	cases := []testCase{
		{
			Cell:   Cell{3, 5},
			Serial: 8,
			Power:  4,
		},
		{
			Cell:   Cell{122, 79},
			Serial: 57,
			Power:  -5,
		},
		{
			Cell:   Cell{217, 196},
			Serial: 39,
			Power:  0,
		},
		{
			Cell:   Cell{101, 153},
			Serial: 71,
			Power:  4,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v_%v", c.Cell, c.Serial), func(t *testing.T) {
			power := c.Cell.Power(c.Serial)
			if power != c.Power {
				t.Errorf("Expected %d, got %d", c.Power, power)
			}
		})
	}
}

func TestMaxPower(t *testing.T) {
	type testCase struct {
		Serial   int
		MaxPower int
		Location Point
	}

	cases := []testCase{
		{
			Serial:   18,
			MaxPower: 29,
			Location: Point{33, 45},
		},
		{
			Serial:   42,
			MaxPower: 30,
			Location: Point{21, 61},
		},
		// Actual Problem 1 input:
		{
			Serial:   9005,
			MaxPower: 31,
			Location: Point{20, 32},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c), func(t *testing.T) {
			power, loc := Grid(c.Serial).MaxPower(300, 3)
			if power != c.MaxPower || loc != c.Location {
				fmt.Errorf("Expected: %v, %v got %v, %v", c.MaxPower, c.Location, power, loc)
			}
		})
	}
}

func TestMaxPowerAnySize(t *testing.T) {
	type testCase struct {
		Serial   int
		Size     int
		MaxPower int
		Location Point
	}

	cases := []testCase{
		{
			Serial:   18,
			MaxPower: 113,
			Size:     16,
			Location: Point{90, 16},
		},
		{
			Serial:   42,
			MaxPower: 119,
			Size:     12,
			Location: Point{232, 251},
		},
		// Actual Problem 2 input:
		{
			Serial:   9005,
			MaxPower: 148,
			Size:     13,
			Location: Point{235, 287},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c), func(t *testing.T) {
			power, loc, size := Grid(c.Serial).MaxPowerAnySize(300)
			if power != c.MaxPower || loc != c.Location || size != c.Size {
				fmt.Errorf("Expected: %v, %v, %v got %v, %v, %v", c.MaxPower, c.Location, c.Size, power, loc, size)
			}
		})
	}
}
