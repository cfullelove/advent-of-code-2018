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

	nums := []int{}

	ss := bufio.NewScanner(f)
	for ss.Scan() {
		var n int
		_, err := fmt.Sscanf(ss.Text(), "%d", &n)
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, n)
	}

	sum := 0
	seen := map[int]bool{0: true}
	for {
		for _, n := range nums {
			sum += n

			if b, ok := seen[sum]; ok && b {
				fmt.Println(sum)
				return
			}

			seen[sum] = true
		}
	}
}
