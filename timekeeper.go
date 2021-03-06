package timekeeper

import (
	"log"
	"time"
)

// EndText will be used to format the elapsed time and EndTime
var EndText string = "calculation took %s"
var start time.Time

// StartTime starts the time keeping. Use it at the beginning of a function.
func StartTime() {
	start = time.Now()
}

// EndTime ends the time keeping and prints the elapsed time. Use `defer timekeeper.EndTime()` right after StartTime.
func EndTime() {
	elapsed := time.Since(start)
	log.Printf(EndText, elapsed)
}
