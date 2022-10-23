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

	for {
		cpuPercents, err := cpu.Percent(refresh_time, false)
		if err != nil {
			log.Fatalln("Cannot get CPU percentage:", err)
		}

		percent := int(math.Round(cpuPercents[0]))
		fmt.Printf("%d%%\n", percent)
	}
}
