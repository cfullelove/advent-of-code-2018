package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func react(units []byte) int {
	for {
		reacted := false

		var newUnits []byte
		skipNext := false
		for i, a := range units {
			if skipNext {
				// Previous unit reacted, with the current one
				skipNext = false
				continue
			}

			if i == len(units)-1 {
				// Last unit, nothing more to react with
				newUnits = append(newUnits, a)
				break
			}

			b := units[i+1]

			if a == b {
				// Same type and charge, can't react
				newUnits = append(newUnits, a)
				continue
			}

			if strings.ToUpper(string(a)) == strings.ToUpper(string(b)) {
				// Different charge, same type, it react!
				skipNext = true
				reacted = true
				continue
			}

			newUnits = append(newUnits, a)
		}

		units = newUnits

		if !reacted {
			break
		}
	}

	return len(units)
}

var types = []byte("abcdefghijklmnopqrstuvwxyz")

func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var units []byte

	for {
		r := make([]byte, 128)
		l, err := f.Read(r)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		units = append(units, r[0:l]...)

		if err == io.EOF {
			break
		}
	}

	minLength := -1
	for _, t := range types {
		filtered := []byte{}

		for _, u := range units {
			if strings.ToLower(string(u)) == string(t) {
				continue
			}

			filtered = append(filtered, u)
		}

		l := react(filtered)
		if l < minLength || minLength == -1 {
			minLength = l
		}

	}

	fmt.Println(minLength)

}
