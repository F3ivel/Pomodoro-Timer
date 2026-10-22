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

func main() {
	fmt.Println("\tHello from 🍅-Timer\n")

	workDuration := 25 * time.Minute
	breakDuration := 5 * time.Minute
	currentPhase := "Work phase"

	fmt.Println("🍅-Timer started\n")

	countdown(workDuration, currentPhase)
	fmt.Printf("\n✅ %s finished!\n", currentPhase)
	switchPhase(&currentPhase)

	countdown(breakDuration, currentPhase)
	fmt.Printf("\n✅ %s finished!", currentPhase)
}
