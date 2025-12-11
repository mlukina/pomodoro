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

	fmt.Println()
	fmt.Println("========================================")
	fmt.Printf("       FOCUS SESSION (%d min)\n", focusTimer)
	fmt.Println("========================================")
	fmt.Println()

	for focusTimer > 0 {
		// Added 4 spaces to clear line from leftover characters
		fmt.Printf("\r%d min    ", focusTimer)
		time.Sleep(time.Minute)

		focusTimer--
		alertTimer++

		switch alertTimer {
		case 30:
			fmt.Println()
			fmt.Println()
			fmt.Println("====================================")
			fmt.Println("      👀 Eye break reminder")
			fmt.Println("     🧍 Switch desk to STANDING")
			fmt.Println("====================================")
			fmt.Println()
		case 60:
			fmt.Println()
			fmt.Println()
			fmt.Println("====================================")
			fmt.Println("       👀 Eye break reminder")
			fmt.Println("      💺 Switch desk to SITTING")
			fmt.Println("====================================")
			fmt.Println()
		}
	}

	fmt.Printf("\n\n✅ Focus done! Starting %d min break\n", breakTimer)

	fmt.Println()
	fmt.Println("========================================")
	fmt.Printf("       BREAK TIME (%d min)\n", breakTimer)
	fmt.Println("========================================")
	fmt.Println()

	for breakTimer > 0 {
		fmt.Printf("\r%d min    ", breakTimer)
		time.Sleep(time.Minute)
		breakTimer--
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("Break time is over! Ready for the next session?")
}