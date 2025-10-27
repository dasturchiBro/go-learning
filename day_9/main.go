package main

import (
	"fmt"
	"context"
	"log"
	"google.golang.org/genai"
	"bufio"
	"os"
	"time"
)

func TextAI(prompt string) string {
	ctx := context.Background()

	client, err := genai.NewClient(
		ctx, nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	ZeroConst := int32(0)
	config := &genai.GenerateContentConfig{
		ThinkingConfig: &genai.ThinkingConfig{
			ThinkingBudget: &ZeroConst,
		},
		SystemInstruction: genai.NewContentFromText("You are a renowned teacher who teaches Math and Physics to students at schools and universities. You explain things in a short and simple way so that even worse-performing students get you.", genai.RoleUser),
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(prompt),
		config,
	)
	if err != nil {
		log.Fatal(err)
	}
	return result.Text()
}

func main() {
	fmt.Print("***Welcome to an AI assistant app***\n\nEnter /exit to exit.\n\n")
	for true {
		fmt.Print("\n\tEnter you prompt: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		prompt := scanner.Text()
		if prompt == "/exit" {
			break
		} 
		fmt.Println(TextAI(prompt))
		time.Sleep(400 * time.Millisecond)
	}
}