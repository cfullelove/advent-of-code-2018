package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const STEP_DURATION = 60
const NUM_WORKERS = 5

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

	workers := &WorkerPool{instructions: instructions}

	t := 0
	for workers.Tick() {
		t++
	}

	fmt.Println(t + 1)
	// 936

}

type WorkerPool struct {
	Completed    []string
	InProgress   [NUM_WORKERS]*step
	instructions *Instructions
}

func (wp *WorkerPool) Tick() bool {

	// Allocate Work
	for i, spot := range wp.InProgress {
		if spot != nil {
			// This spot already allocated
			continue
		}

		next, ok := wp.instructions.NextStep()
		if !ok {
			// No instructions ready to work on
			break
		}

		wp.InProgress[i] = next
		next.InProgress = true
	}

	// Do work
	for i, step := range wp.InProgress {
		if step == nil {
			continue
		}

		done := step.Work()
		if done {
			wp.Completed = append(wp.Completed, step.Label)
			wp.InProgress[i] = nil
			wp.instructions.FinishStep(step.Label)
		}
	}

	// Return value is true if instructions aren't completed (i.e. keep going)
	return !wp.instructions.IsFinished()
}

func (wp *WorkerPool) String() string {
	ss := fmt.Sprintln("Done:", wp.Completed)
	for i, s := range wp.InProgress {
		if s == nil {
			ss = ss + fmt.Sprintln(i, "N/A")
		} else {
			ss = ss + fmt.Sprintln(i, s.Label, s.TimeUntilDone)
		}
	}

	return ss
}

type Instructions map[string]*step

func (i Instructions) StepExists(label string) bool {
	for l, _ := range i {
		if l == label {
			return true
		}
	}

	return false
}

func (i Instructions) AddStep(label string, step *step) {
	i[label] = step
	i[label].TimeUntilDone = time.Duration(STEP_DURATION+strings.Index(ALPHABET, label)+1) * time.Second
	i[label].Label = label
}

func (i Instructions) GetStep(label string) *step {
	if _, ok := i[label]; !ok {
		i[label] = &step{}
	}

	return i[label]
}

func (i Instructions) IsFinished() bool {
	for _, step := range i {
		if !step.Done {
			return false
		}
	}

	return true
}

func (i Instructions) NextStep() (*step, bool) {

	available := []string{}
	// Find all steps without dependencies
	for label, step := range i {
		if len(step.Depends) == 0 && !step.Done && !step.InProgress {
			available = append(available, label)
		}
	}

	sort.Strings(available)

	if len(available) == 0 {
		return nil, false
	}

	return i[available[0]], true
}

func (i Instructions) FinishStep(label string) {
	i[label].Done = true
	for _, s := range i {
		s.Remove(label)
	}
}

func (i Instructions) String() string {

	labels := []string{}
	for l, _ := range i {
		labels = append(labels, l)
	}

	sort.Strings(labels)

	s := ""
	for _, step := range labels {
		s = s + fmt.Sprintln(step, i[step])
	}

	return s
}

type step struct {
	Label         string
	Done          bool
	InProgress    bool
	TimeUntilDone time.Duration
	Depends       []string
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

func (s *step) Work() bool {
	s.TimeUntilDone -= time.Second

	s.Done = s.TimeUntilDone <= 0

	return s.Done
}
