package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

type RawEvent struct {
	timestamp time.Time
	event     string
}

type ByTime []RawEvent

func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool { return a[i].timestamp.Before(a[j].timestamp) }

type Log struct {
	ID     int
	Asleep time.Time
	Awake  time.Time
}

func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	events := []RawEvent{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		var year, month, day, hour, minute int
		_, err := fmt.Sscanf(s.Text()[0:18], "[%d-%d-%d %d:%d]", &year, &month, &day, &hour, &minute)
		if err != nil {
			log.Fatal(err)
		}

		ts := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.Local)

		events = append(events, RawEvent{timestamp: ts, event: s.Text()[19:]})
	}

	sort.Sort(ByTime(events))

	sleepLog := []Log{}

	var currentLog Log
	for _, e := range events {
		if e.event == "falls asleep" {
			currentLog.Asleep = e.timestamp
			continue
		}

		if e.event == "wakes up" {
			currentLog.Awake = e.timestamp
			sleepLog = append(sleepLog, currentLog)
			continue
		}

		_, err := fmt.Sscanf(e.event, "Guard #%d begins shift", &currentLog.ID)
		if err != nil {
			log.Fatal(err)
		}
	}

	mostSleepID := -1
	mostSleepTotal := 0
	sleepTotalID := map[int]int{}
	sleepTracker := map[int]map[int]int{}
	for _, l := range sleepLog {
		if _, ok := sleepTracker[l.ID]; !ok {
			sleepTracker[l.ID] = map[int]int{}
		}
		if _, ok := sleepTotalID[l.ID]; !ok {
			sleepTotalID[l.ID] = 0
		}

		sleepTotalID[l.ID] = sleepTotalID[l.ID] + l.Awake.Minute() - l.Asleep.Minute()
		if sleepTotalID[l.ID] > mostSleepTotal {
			mostSleepTotal = sleepTotalID[l.ID]
			mostSleepID = l.ID
		}

		for i := l.Asleep.Minute(); i < l.Awake.Minute(); i++ {
			if _, ok := sleepTracker[l.ID][i]; !ok {
				sleepTracker[l.ID][i] = 0
			}

			sleepTracker[l.ID][i] = sleepTracker[l.ID][i] + 1
		}
	}

	maxMin := -1
	maxMins := 0
	for i, min := range sleepTracker[mostSleepID] {
		if min > maxMins {
			maxMins = min
			maxMin = i
		}
	}

	fmt.Println(mostSleepID * maxMin)

}
