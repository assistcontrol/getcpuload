package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/DataDog/gopsutil/cpu"
)

const refresh_time = 3 * time.Second

func main() {
	for {
		cpuPercents, err := cpu.Percent(refresh_time, false)
		if err != nil {
			log.Fatalln("Cannot get CPU percentage:", err)
		}

		percent := int(math.Round(cpuPercents[0]))
		fmt.Printf("%d%%\n", percent)
	}
}
