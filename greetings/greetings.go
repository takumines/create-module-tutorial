package greetings

import (
	"errors"
	"fmt"
)

// Hello 引数に受け取った名前を使ってメッセージを作成する
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message, nil
}
