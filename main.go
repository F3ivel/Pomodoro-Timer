package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func countdown(duration time.Duration, phase string) {
	fmt.Println("🍅-Phase:", phase)

	end := time.Now().Add(duration)

	for {
		remaining := time.Until(end)

		if remaining <= 0 {
			break
		}

		minutes := int(remaining.Seconds()) / 60
		seconds := int(remaining.Seconds()) % 60

		fmt.Printf("\r⏳ %02d:%02d remaining", minutes, seconds)

		time.Sleep(1 * time.Second)
	}

	//fmt.Print("\r⏳ 00:00 remaining\n")

	// Play terminal sound
	for i := 0; i < 4; i++ {
		fmt.Print("\a")
		time.Sleep(2000 * time.Millisecond)
	}
}

func switchPhase(phase *string) {

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

func readSettings(workDuration, breakDuration *time.Duration, autoStartTimer *bool) {
	file, err := os.Open("settings.txt")

	if err != nil {
		fmt.Println("Error: File could not be opened!", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for {
		if scanner.Scan() {
			line := scanner.Text()

			// Filters empty lines, lines with no assignement and comments out
			if line == "" || !strings.Contains(line, "=") || strings.HasPrefix(line, "#") {
				continue
			} else {
				// Reads the "workDuration" setting
				if strings.HasPrefix(line, "workDuration=") {
					_, value, _ := strings.Cut(line, "=")
					value = strings.ReplaceAll(value, "\"", "")
					tmp, err := time.ParseDuration(value)

					if err != nil {
						fmt.Println("Invalid value for the setting \"workDuration\" detected.", err)
						continue
					} else {
						*workDuration = tmp
					}
				}
				// Reads the "breakDuration" setting
				if strings.HasPrefix(line, "breakDuration=") {
					_, value, _ := strings.Cut(line, "=")
					value = strings.ReplaceAll(value, "\"", "")
					tmp, err := time.ParseDuration(value)

					if err != nil {
						fmt.Println("Invalid value for the setting \"breakDuration\" detected.", err)
						continue
					} else {
						*breakDuration = tmp
					}
				}
				// Reads the "autoStartTimer" setting
				if strings.HasPrefix(line, "autoStartTimer=") {
					_, value, _ := strings.Cut(line, "=")
					value = strings.ReplaceAll(value, "\"", "")

					if (value != "true") && (value != "false") {
						fmt.Println("[INFO]: Invalid value or no value was set for the setting \"autoStartTimer\". Default value \"false\" will be used.\n")
					} else {
						// Set autoStartTimer to true, otherwise default value false will be used
						if value == "true" {
							*autoStartTimer = true
						}
					}
				}
			}
		} else {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	if *workDuration <= 0 {
		fmt.Println("Invalid value or no value was set for the setting \"workDuration\". Default value \"25m\" will be used.")
		*workDuration = 25 * time.Minute
	}

	if *breakDuration <= 0 {
		fmt.Println("Invalid value or no value was set for the setting \"breakDuration\". Default value \"5m\" will be used.")
		*breakDuration = 5 * time.Minute
	}
}

func startTimer(autoStartTimer *bool) {
	if *autoStartTimer {
		fmt.Println("🍅-Timer started\n")
	} else {
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Printf("Would you like to start the Pomodoro timer? (yes or no): ")

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

func main() {
	var workDuration time.Duration = -1 * time.Minute
	var breakDuration time.Duration = -1 * time.Minute
	var autoStartTimer bool = false
	currentPhase := "Work phase"

	fmt.Println("\tHello from 🍅-Timer\n")
	readSettings(&workDuration, &breakDuration, &autoStartTimer)

	startTimer(&autoStartTimer)

	countdown(workDuration, currentPhase)
	fmt.Printf("\n✅ %s finished!\n", currentPhase)
	switchPhase(&currentPhase)

	countdown(breakDuration, currentPhase)
	fmt.Printf("\n✅ %s finished!", currentPhase)
}
