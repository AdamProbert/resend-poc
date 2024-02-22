package main

import (
	"fmt"
	"log"
	"os"

	"github.com/resend/resend-go/v2"
)

func main() {
	apiKey := os.Getenv("RESEND_API_KEY")

	client := resend.NewClient(apiKey)
	emailHtml := readHTMLFile("email.html")

	params := &resend.SendEmailRequest{
		From:    "Mad Mushroom <onboarding@resend.dev>",
		To:      []string{"adster999@hotmail.com"},
		Html:    emailHtml,
		Subject: "Hello from Mad Mushroom!",
		Cc:      []string{"cc@example.com"},
		Bcc:     []string{"bcc@example.com"},
		ReplyTo: "replyto@example.com",
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(sent.Id)

	// Get email by id
	email, err := client.Emails.Get(sent.Id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(email.LastEvent)
}

func readHTMLFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
