package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/mackerelio/go-osstat/memory"
	"github.com/shirou/gopsutil/v3/cpu"
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
	cpuPercents, err := cpu.Percent(refreshTime, false)
	if err != nil {
		log.Fatalln("Cannot get CPU percentage:", err)
	}

	percent := int(math.Round(cpuPercents[0]))
	return fmt.Sprintf("%d%%", percent)
}

func getMem() string {
	mem, _ := memory.Get()

	used := bytesToGB(mem.Used)
	usedPercent := int(math.Round(float64(mem.Used) / float64(mem.Total) * 100))

	return fmt.Sprintf("%d%%(%dG)", usedPercent, used)
}

func bytesToGB(bytes uint64) int {
	return int(math.Round(float64(bytes) / 1024 / 1024 / 1024))
}
