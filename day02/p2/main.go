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

	boxIDs := []string{}

	ss := bufio.NewScanner(f)
	for ss.Scan() {
		boxIDs = append(boxIDs, ss.Text())
	}

	// Probably not efficient
	for _, s1 := range boxIDs {
		for _, s2 := range boxIDs {
			wrongTimes := 0
			wrongIndex := -1
			for k, _ := range s1 {
				if s1[k] != s2[k] {
					wrongTimes++
					wrongIndex = k
				}

				if wrongTimes > 1 {
					break
				}
			}
			if wrongTimes == 1 {
				fmt.Println(s1[0:wrongIndex] + s1[wrongIndex+1:])
				return
			}
		}
	}

}
