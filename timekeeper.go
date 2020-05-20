package timekeeper

import (
	"log"
	"time"
)

var start time.Time

func StartTime() {
	start = time.Now()
}

func EndTime() {
	elapsed := time.Since(start)
	log.Printf("calculation took %s", elapsed)
}
