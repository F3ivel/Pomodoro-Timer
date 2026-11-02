package main

import (
	"fmt"
	"pomodoro-timer/config"
	"pomodoro-timer/terminal"
	"pomodoro-timer/timer"
	"time"
)

func main() {
	var workDuration time.Duration = -1 * time.Minute
	var breakDuration time.Duration = -1 * time.Minute
	var autoStartTimer bool = false
	currentPhase := "Work phase"

	fmt.Println("\tHello from 🍅-Timer\n")
	config.ReadSettings(&workDuration, &breakDuration, &autoStartTimer)

	terminal.StartTimer(&autoStartTimer)

	timer.Countdown(workDuration, currentPhase)
	fmt.Printf("\n✅ %s finished!\n", currentPhase)
	terminal.SwitchPhase(&currentPhase)

	timer.Countdown(breakDuration, currentPhase)
	fmt.Printf("\n✅ %s finished!", currentPhase)
}
