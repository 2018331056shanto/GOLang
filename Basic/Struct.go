

// A struct is a collection of fields/properties.
// However, note that only exported fields (capitalized) can be accessed from outside of a package.
// You don’t need to define getters and setters on struct fields, they can be accessed automatically.
package main

import (
	"fmt"
	"time"
)

type Bootcamp struct {
	Lat, Lon float64
	Date     time.Time
}

func main() {
	event := Bootcamp{
		Lat: 34.012836,
		Lon: -118.495338,
	}
	event.Date = time.Now()
	fmt.Printf("Event on %s, location (%f, %f)",
		event.Date, event.Lat, event.Lon)

}