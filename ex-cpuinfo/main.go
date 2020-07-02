package main

import (
	"log"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func main() {

	var dur = 1 * time.Second

	for range time.Tick(dur) {
		cpuPercents, err := cpu.Percent(dur, false)
		if err != nil {
			log.Printf("Error getting cpu info: #{err}")
		}
		log.Printf("The current cpu percent: %.2f", cpuPercents[0])
	}

}
