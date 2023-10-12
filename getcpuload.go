package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

const REFRESH_TIME = 3 * time.Second

func main() {
	// Show an initial result quickly
	fmt.Println(get(1 * time.Second))

	for {
		fmt.Println(get(REFRESH_TIME))
	}
}

func get(refreshTime time.Duration) string {
	curCPU := getCPU(refreshTime)
	curMem := getMem()
	return fmt.Sprintf("%s %s", curCPU, curMem)
}

func getCPU(refreshTime time.Duration) string {
	c, err := cpu.Percent(refreshTime, false)
	if err != nil {
		log.Fatalln("Cannot get CPU percentage:", err)
	}

	return percentString(c[0])
}

func getMem() string {
	m, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalln("Cannot get memory info:", err)
	}

	return fmt.Sprintf("%s(%dG)", percentString(m.UsedPercent), bytesToGB(m.Used))
}

func bytesToGB(b uint64) int {
	return int(math.Round(float64(b) / 1024 / 1024 / 1024))
}

// Returns a string with a rounded integer and a percent sign
func percentString(f float64) string {
	return fmt.Sprintf("%d%%", int(math.Round(f)))
}
