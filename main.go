package main

import (
	"fmt"
	"pomodoro-timer/config"
	"pomodoro-timer/terminal"
	"pomodoro-timer/timer"
	"time"
)

func main() {
	settings := config.Settings{
		WorkDuration:   -1 * time.Minute,
		BreakDuration:  -1 * time.Minute,
		AutoStartTimer: false,
	}
	currentPhase := "Work phase"

	fmt.Println("\tHello from 🍅-Timer\n")
	config.ReadSettings(&settings)

	terminal.AutoStartTimer(&settings)

	timer.Countdown(&settings, currentPhase)
	fmt.Printf("\n✅ %s finished!\n", currentPhase)
	terminal.SwitchPhase(&currentPhase)

	timer.Countdown(&settings, currentPhase)
	fmt.Printf("\n✅ %s finished!", currentPhase)
}
