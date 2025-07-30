package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting ntp time:", err)
		os.Exit(1)
	}

	fmt.Println("Exact time:", exactTime.Format(time.RFC1123))
}
