package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	var alertTimer int

	focusTimer, err := getValidatedInput("Focus time (min)? ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	breakTimer, err := getValidatedInput("Break time (min)? ")
	if err != nil {
		fmt.Println("Error:", err)
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

func getValidatedInput(prompt string) (int, error) {
	var timeInMin int

	fmt.Print(prompt)

	_, err := fmt.Scan(&timeInMin)
	if err != nil {
		return 0, err
	}

    err = validateNumber(timeInMin)
    if err != nil {
        return 0, err
    }

	return timeInMin, nil
}

func validateNumber(n int) error {
    if n <= 0 {
        return errors.New("input must be positive")
    }
    return nil
}