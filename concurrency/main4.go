package main

import (
	"time"
	"fmt"
)

func Intro(text string) {
	fmt.Println(text)
}

func GetMessages(channel chan string) {
	messages := []string{"The file 4 was closed!", "Notifications were turned off", "The house was cleaned", "The car was charged"}
	for _, msg := range messages {
		channel <- msg
	}
}	

func GiveValue(ch chan string) {
	for msg := range ch {
		SendMessage(msg)
	}
	close(ch)
}

func SendMessage(message string) {
	fmt.Println("\n-----Message sent to the user.\nMessage: ", message, "-------\n")
}

func main() {
	ch := make(chan string, 4)

	go Intro("Hello.")
	go Intro("What's up?")
	go Intro("How is it going?")
	go Intro("Hey!!!")
	go GetMessages(ch)
	go GiveValue(ch)
	time.Sleep(2 * time.Second)
}