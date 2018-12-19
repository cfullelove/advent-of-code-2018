package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var data []int

func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		// fmt.Println(s.Text())
		num, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, num)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	node := Parse(data)

	fmt.Println(node)

	fmt.Println(node.Value())
	// 37067
}

func Parse(data []int) *Node {
	if len(data) < 2 {
		return nil
	}

	node := &Node{}

	numChildren := data[0]
	lenData := data[1]

	offSet := 2
	for i := 0; i < numChildren; i++ {
		child := Parse(data[offSet:])
		offSet += child.Length()
		node.Children = append(node.Children, child)
	}

	node.Data = data[offSet : offSet+lenData]

	return node
}

type Node struct {
	Children []*Node
	Data     []int
}

func (node *Node) Length() int {
	length := 2
	for _, child := range node.Children {
		length += child.Length()
	}

	length += len(node.Data)

	return length
}

func (node *Node) Value() int {

	if len(node.Children) == 0 {
		return node.Sum()
	}

	value := 0
	for _, i := range node.Data {
		i-- // Index starts at 1, not 0
		if i >= len(node.Children) {
			continue
		}

		value += node.Children[i].Value()
	}

	return value
}

func (node *Node) Sum() int {
	sum := 0

	for _, child := range node.Children {
		sum += child.Sum()
	}

	for _, d := range node.Data {
		sum += d
	}

	return sum
}

func (node *Node) String() string {
	w := &bytes.Buffer{}

	fmt.Fprintf(w, "Num Chilren: %v\n", len(node.Children))
	for _, child := range node.Children {
		s := bufio.NewScanner(strings.NewReader(child.String()))
		for s.Scan() {
			fmt.Fprintln(w, " ", s.Text())
		}
		if err := s.Err(); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Fprintln(w, "Data:", node.Data, "Value:", node.Value())

	return w.String()
}
