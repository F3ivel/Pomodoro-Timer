package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Settings struct {
	WorkDuration   time.Duration
	BreakDuration  time.Duration
	AutoStartTimer bool
}

func ReadSettings(settings *Settings) {
	file, err := os.Open("config/settings.txt")

	if err != nil {
		fmt.Println("[Error]: Application settings could not be read -> Default values for the settings will be used.\n")
		settings.WorkDuration = 25 * time.Minute
		settings.BreakDuration = 5 * time.Minute
		settings.AutoStartTimer = false
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

					if (err != nil) || (tmp <= 0) {
						fmt.Println("[INFO]: Invalid value or no value was set for the setting \"workDuration\" -> Default value \"25m\" will be used.")
						settings.WorkDuration = 25 * time.Minute
						continue
					} else {
						settings.WorkDuration = tmp
					}
				}
				// Reads the "breakDuration" setting
				if strings.HasPrefix(line, "breakDuration=") {
					_, value, _ := strings.Cut(line, "=")
					value = strings.ReplaceAll(value, "\"", "")
					tmp, err := time.ParseDuration(value)

					if (err != nil) || (tmp <= 0) {
						fmt.Println("[INFO]: Invalid value or no value was set for the setting \"breakDuration\" -> Default value \"5m\" will be used.")
						settings.BreakDuration = 5 * time.Minute
						continue
					} else {
						settings.BreakDuration = tmp
					}
				}
				// Reads the "autoStartTimer" setting
				if strings.HasPrefix(line, "autoStartTimer=") {
					_, value, _ := strings.Cut(line, "=")
					value = strings.ReplaceAll(value, "\"", "")

					if (value != "true") && (value != "false") {
						fmt.Println("[INFO]: Invalid value or no value was set for the setting \"autoStartTimer\" -> Default value \"false\" will be used.")
						settings.AutoStartTimer = false
					} else {
						if value == "true" {
							settings.AutoStartTimer = true
						} else {
							settings.AutoStartTimer = false
						}
					}
				}
			}
		} else {
			break
		}
	}

	fmt.Println()

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
