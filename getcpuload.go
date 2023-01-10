package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/DataDog/gopsutil/cpu"
)

func main() {
	secs := flag.Int("r", 3, "refresh time (seconds)")
	flag.Parse()

	refreshTime := time.Duration(*secs) * time.Second

	// Show an initial result quickly
	fmt.Println(cpuPercent(1 * time.Second))

	for {
		fmt.Println(cpuPercent(refreshTime))
	}
}

func cpuPercent(refreshTime time.Duration) string {
	cpuPercents, err := cpu.Percent(refreshTime, false)
	if err != nil {
		log.Fatalln("Cannot get CPU percentage:", err)
	}

	percent := int(math.Round(cpuPercents[0]))
	return fmt.Sprintf("%d%%", percent)
}
