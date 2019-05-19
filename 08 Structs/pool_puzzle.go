package main

import (
	"./geo"
	"fmt"
)

func main() {
	location := geo.Landmark{
		Name: "The Googleplex",
	}
	location.Latitude = 37.42
	location.Longitude = -122.08
	fmt.Println(location)
}
