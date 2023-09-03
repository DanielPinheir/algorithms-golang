package main

import (
	"fmt"
	"math"
)

func main() {
	var fileSize float64
	var internetSpeed float64
	
	fmt.Print("Enter the download file size (in MB): ")
	fmt.Scan(&fileSize)

	fmt.Print("Enter your internet speed (in Mbps): ")
	fmt.Scan(&internetSpeed)

	downloadTime := (fileSize * 8) / internetSpeed 
	downloadTimeMinutes := math.Ceil(downloadTime / 60)

	if downloadTimeMinutes <= 1 {
		fmt.Printf("Approximate download time is %.2f minutes.\n", downloadTimeMinutes)
	} else {
		fmt.Printf("Approximate download time is %.2f minutes.\n", downloadTimeMinutes)
	}

}