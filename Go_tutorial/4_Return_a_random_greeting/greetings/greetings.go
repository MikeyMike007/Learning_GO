package greetings

import (
		"fmt"
		"math/rand"
		"time"
		"errors"
)

func Hello(name string) (string, error) {
		if name == "" {
				return name, errors.New("Empty name")
		}

		message := fmt.Sprintf(randomFormat(), name)

		return message, nil
}

func init() {
		rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
		formats := []string {
				"Hi %v. Welcome!",
				"Hi, great to see you %v",
				"Hail, %v! Well met!",
		}

		return formats[rand.Intn(len(formats))]
}



