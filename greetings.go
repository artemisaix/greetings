package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("ERROR: Nombre vacio")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Greet(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hola, %v, bienvenido al modulo",
		"Que bueno verte, %v",
		"Saludos desde el modulo greetins, %v",
	}
	return formats[rand.Intn(len(formats))]
}
