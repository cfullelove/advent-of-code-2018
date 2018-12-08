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

	sum := 0

	ss := bufio.NewScanner(f)
	for ss.Scan() {
		var n int
		_, err := fmt.Sscanf(ss.Text(), "%d", &n)
		if err != nil {
			log.Fatal(err)
		}

		sum += n
	}

	fmt.Println(sum)

}
