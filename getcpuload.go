package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/DataDog/gopsutil/cpu"
)

var secs = flag.Int("r", 3, "refresh time (seconds)")

func main() {
	flag.Parse()
	refresh_time := time.Duration(*secs) * time.Second

	// Show an initial result quickly
	fmt.Println(cpuPercent(1 * time.Second))

	for {
		fmt.Println(cpuPercent(refresh_time))
	}
}

func cpuPercent(refresh_time time.Duration) string {
	cpuPercents, err := cpu.Percent(refresh_time, false)
	if err != nil {
		log.Fatalln("Cannot get CPU percentage:", err)
	}

	percent := int(math.Round(cpuPercents[0]))
	return fmt.Sprintf("%d%%", percent)
}
