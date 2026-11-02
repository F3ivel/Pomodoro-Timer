package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func ReadSettings(workDuration, breakDuration *time.Duration, autoStartTimer *bool) {
	file, err := os.Open("config/settings.txt")

	if err != nil {
		fmt.Println("[Error]: Application settings could not be read -> Default values for the settings will be used.\n")
		*workDuration = 25 * time.Minute
		*breakDuration = 5 * time.Minute
		*autoStartTimer = false
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
						*workDuration = 25 * time.Minute
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

					if (err != nil) || (tmp <= 0) {
						fmt.Println("[INFO]: Invalid value or no value was set for the setting \"breakDuration\" -> Default value \"5m\" will be used.")
						*breakDuration = 5 * time.Minute
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
						fmt.Println("[INFO]: Invalid value or no value was set for the setting \"autoStartTimer\" -> Default value \"false\" will be used.")
						*autoStartTimer = false
					} else {
						if value == "true" {
							*autoStartTimer = true
						} else {
							*autoStartTimer = false
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
