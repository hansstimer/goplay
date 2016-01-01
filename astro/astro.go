package main

import (
	"fmt"
	"time"

	"github.com/cpucycle/astrotime"
)

const latitude = float64(38.8895)
const longitude = float64(77.0352)

func main() {
	// t := time.Date(2015, time.December, 31, 23, 0, 0, 0, time.Local)
	t := time.Now()
	fmt.Printf("Go launched at %s\n", t.Local())

	nextSunrise := astrotime.NextSunrise(t, latitude, longitude)
	fmt.Printf("next sunrise at %s\n", nextSunrise)
	nextSunset := astrotime.NextSunset(t, latitude, longitude)
	fmt.Printf("next sunset at %s\n", nextSunset)

	if nextSunset.Before(nextSunrise) {
		fmt.Printf("It is daytime\n")
	} else {
		fmt.Printf("It is nighttime\n")
	}
}
