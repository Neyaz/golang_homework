package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("current time:", time.Now().Round(time.Nanosecond).String())
	fmt.Println("exact time:", exactTime.Round(time.Nanosecond).String())
}
