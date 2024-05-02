package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func GetTime() {
	timeNTP, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	currentTime := time.Now()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	} else {
		fmt.Println("Текущее время:", currentTime)
		fmt.Println("Точное время (NTP):", timeNTP)
	}
}

func main() {
		GetTime()
}