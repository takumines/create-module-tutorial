package greetings

import (
	"fmt"
	"math/rand"
)

// Hello 引数に受け取った名前を使ってメッセージを作成する
func Hello(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

// ランダムなフォーマットのメッセージを返す
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
