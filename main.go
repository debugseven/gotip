package main

import (
	"fmt"
	"log"
	"os"

	"github.com/debugseven/gotip/lib"
)

func main() {
	//time_17 := uint32((lib.SecondsPerHour * 23) + (lib.SecondsPerHour * 1.49))

	args := os.Args[1:]
	var time lib.Time

	if len(args) > 1 {
		log.Fatal("Invalid arguments length")
	} else if len(args) == 0 {
		t, err := lib.Current()

		if err != nil {
			log.Fatal(err)
		}

		time = t
	} else {
		t, err := lib.FromTime(args[0])

		if err != nil {
			log.Fatal(err)
		}

		time = t
	}

	fmt.Printf("result: %v", time.Info())

}
