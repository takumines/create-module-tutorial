package greetings

import "fmt"

// Hello 引数に受け取った名前を使ってメッセージを作成する
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
