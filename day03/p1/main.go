package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type claim struct {
	ID, X, Y, Width, Height int
}

var maxWidth = 1000
var maxHeight = 1000

func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fabric := map[int]int{}

	claims := []claim{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		var c claim
		_, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &c.ID, &c.X, &c.Y, &c.Width, &c.Height)
		if err != nil {
			log.Fatal(err)
		}

		for i := c.X; i < (c.X + c.Width); i++ {
			for j := c.Y; j < (c.Y + c.Height); j++ {
				idx := j*(maxWidth) + i
				if _, ok := fabric[idx]; !ok {
					fabric[idx] = 0
				}

				fabric[idx] = fabric[idx] + 1
				// fmt.Println("idx", idx, fabric[idx])
			}
		}

		claims = append(claims, c)
	}

	count := 0
	for _, c := range fabric {
		if c > 1 {
			count++
		}
	}

	fmt.Println(count)

}
