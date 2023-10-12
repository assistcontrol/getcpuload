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

var MEM_TOTAL = 0.0

func init() {
	mem, err := memory.Get()
	if err != nil {
		log.Fatal(err)
	}
	MEM_TOTAL = float64(mem.Total)
}

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
	mem, err := memory.Get()
	if err != nil {
		log.Fatalln("Cannot get memory info:", err)
	}

	used := mem.Used
	usedPercent := int(math.Round(float64(used) / MEM_TOTAL * 100))

	return fmt.Sprintf("%d%%(%dG)", usedPercent, bytesToGB(used))
}

func bytesToGB(bytes uint64) int {
	return int(math.Round(float64(bytes) / 1024 / 1024 / 1024))
}
