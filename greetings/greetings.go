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

// Hellos 引数に受け取った名前のスライスを使って、複数のメッセージを作成する
func Hellos(names []string) (map[string]string, error) {
	// メッセージを格納するためのマップを作成
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

// ランダムなフォーマットのメッセージを返す
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
