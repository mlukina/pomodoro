package main

import (
	"fmt"
	"time"
)

func main() {
	var timer int

	fmt.Print("How long would you like the focus time to be in minutes? ")
	_, err := fmt.Scan(&timer)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for i := timer; i > 0; i-- {
		fmt.Printf("%d min\n", i)
		time.Sleep(time.Minute)
	}

	fmt.Println("Done!")
}