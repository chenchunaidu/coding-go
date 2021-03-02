package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you %v!",
		"Hail %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

//Hellos returns a map that associates each of the names people
//with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	//A map associate names with messages
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
