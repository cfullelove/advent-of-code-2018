package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// boxIDs := []string{}

	exactly2s := 0
	exactly3s := 0

	ss := bufio.NewScanner(f)
	for ss.Scan() {
		// boxIDs = append(boxIDs, ss.Text())
		seen := map[rune]int{}
		for _, s := range ss.Text() {
			if _, ok := seen[s]; !ok {
				seen[s] = 0
			}
			seen[s] = seen[s] + 1
		}

		exactly2 := false
		exactly3 := false

		for _, c := range seen {
			if c == 2 {
				exactly2 = true
			}

			if c == 3 {
				exactly3 = true
			}
		}

		if exactly2 {
			exactly2s++
		}

		if exactly3 {
			exactly3s++
		}
	}

	fmt.Println(exactly2s * exactly3s)

}
