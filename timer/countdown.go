package timer

import (
	"fmt"
	"pomodoro-timer/config"
	"time"
)

func Countdown(settings *config.Settings, phase string) {
	var end time.Time

	fmt.Println("🍅-Phase:", phase)

	if phase == "Work phase" {
		end = time.Now().Add(settings.WorkDuration)
	} else {
		end = time.Now().Add(settings.BreakDuration)
	}

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
