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

// concurrency is especially helpful when building telegram bots that integrate file sending. For example, if the code needs to download a file 
// from an URL and send it to a user, downloading a file takes time, and if concurrency is not used in this situation, it can lead to lags 
// when using the bot when it has to send multiple files at the same time. However, concurrency can solve this problem by allowing the bot 
// to serve other users while trying to download and send the files to the users.