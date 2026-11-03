package terminal

import (
	"bufio"
	"fmt"
	"os"
	"pomodoro-timer/config"
	"strings"
)

func SwitchPhase(phase *string) {

	reader := bufio.NewReader(os.Stdin)

	if *phase == "Work phase" {
		*phase = "Break phase"
	} else {
		*phase = "Work phase"
	}

	for {
		fmt.Printf("Please confirm with \"yes\" to start the %s: ", *phase)

		input, err := reader.ReadString('\n')
		if err != nil && len(input) == 0 {
			fmt.Println("Invalid input!")
			continue
		}

		input = strings.TrimSpace(input)

		if input == "yes" {
			fmt.Println()
			break
		}
	}
}

func AutoStartTimer(settings *config.Settings) {
	if settings.AutoStartTimer {
		fmt.Println("🍅-Timer started\n")
	} else {
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Printf("Would you like to start the Pomodoro timer? (Type \"yes\" to continue): ")

			input, err := reader.ReadString('\n')
			if err != nil && len(input) == 0 {
				fmt.Println("Invalid input!")
				continue
			}

			input = strings.TrimSpace(input)

			if input == "yes" {
				fmt.Println()
				break
			}
		}
		fmt.Println("🍅-Timer started\n")
	}
}
