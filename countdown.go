package main

import (
	"fmt"
	"time"
)

func countdown(minutes int, label string) {
	totalSeconds := minutes * 60

	for i := totalSeconds; i >= 0; i-- {
		fmt.Printf("\r%s: %02d:%02d", label, i/60, i%60)
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
}
