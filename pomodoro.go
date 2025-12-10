package main

import (
	"fmt"
	"time"
)

func main() {
	var focusTimer int
	var breakTimer int
	var alertTimer int

	fmt.Print("Focus time (min)? ")
	_, err := fmt.Scan(&focusTimer)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if focusTimer <= 0 {
		fmt.Println("Error: Please enter a positive number.")
		return
	}

	fmt.Print("Break time (min)? ")
	_, err = fmt.Scan(&breakTimer)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if breakTimer <= 0 {
		fmt.Println("Error: Please enter a positive number")
		return
	}

	for focusTimer > 0 {
		fmt.Printf("%d min\n", focusTimer)
		time.Sleep(time.Second)

		focusTimer--
		alertTimer++

		if alertTimer == 30 {
			fmt.Println("⏰ Eye break! Switch desk to STANDING")
		}
		if alertTimer == 60 {
			fmt.Println("⏰ Eye break! Switch desk to SITTING")
		}
	}

	fmt.Printf("Focus done! Starting %d min break\n", breakTimer)

	for breakTimer > 0 {
		fmt.Printf("%d min\n", breakTimer)
		time.Sleep(time.Second)
		breakTimer--
	}

	fmt.Println("Break time is over! Ready for the next session?")
}