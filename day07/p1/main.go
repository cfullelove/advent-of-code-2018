package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	instructions := &Instructions{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		var first, before string
		_, err := fmt.Sscanf(s.Text(), "Step %s must be finished before step %s can begin.", &first, &before)
		if err != nil {
			log.Fatal(err)
		}

		if !instructions.StepExists(before) {
			instructions.AddStep(before, &step{})
		}

		if !instructions.StepExists(first) {
			instructions.AddStep(first, &step{})
		}

		instructions.GetStep(before).AddDepends(first)

	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	order := ""

	for {
		nextStep, keepGoing := instructions.NextStep()

		if !keepGoing {
			break
		}

		order = order + nextStep
	}

	fmt.Println(order)

	// LAPFCRGHVZOTKWENBXIMSUDJQY
}

type Instructions map[string]*step

func (i Instructions) GetMap() map[string]*step {
	return map[string]*step(i)
}

func (i Instructions) StepExists(label string) bool {
	for l, _ := range i.GetMap() {
		if l == label {
			return true
		}
	}

	return false
}

func (i Instructions) AddStep(label string, step *step) {
	i.GetMap()[label] = step
}

func (i Instructions) GetStep(label string) *step {
	if _, ok := i.GetMap()[label]; !ok {
		i.GetMap()[label] = &step{}
	}

	return i.GetMap()[label]
}

func (i Instructions) NextStep() (string, bool) {

	available := []string{}
	// Find all steps without dependencies
	for label, step := range i.GetMap() {
		if len(step.Depends) == 0 && !step.Done {
			available = append(available, label)
		}
	}

	sort.Strings(available)

	if len(available) == 0 {
		return "", false
	}

	nextStep := available[0]

	// Mark step done
	i.GetStep(available[0]).Done = true

	// Remove from others' dependencies
	for _, step := range i.GetMap() {
		step.Remove(nextStep)
	}

	return available[0], true
}

func (i Instructions) String() string {
	s := ""
	for step, i := range i.GetMap() {
		s = s + fmt.Sprintln(step, i)
	}

	return s
}

type step struct {
	Done    bool
	Depends []string
}

func (s *step) AddDepends(i string) {
	s.Depends = append(s.Depends, i)
}

func (s *step) Remove(i string) {
	depends := []string{}

	for _, d := range s.Depends {
		if d != i {
			depends = append(depends, d)
		}
	}

	s.Depends = depends
}
