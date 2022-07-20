package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

// print all the command events ,give to slack

func commandEvents(analytics <-chan *slacker.CommandEvent) {
	// basically you want to print the timestamp and the events
	for event := range analytics {
		fmt.Println("Command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}

}

func main() {
	// Environment path is providing both bot and app
	os.Setenv("GOLANG_SLACK_BOT_TOKEN", "xoxb-3834357035104-3810658054819-dWWcRzghne2VIbSyBPt88nn8")
	os.Setenv("GOLANG_SLACK_APP_TOKEN", "xapp-1-A03PE07DBST-3816605232451-00a06c27907fab18c1eab3b1e5199af4a5f5f6901dfb0c7351a66325b42da54b")
	// slacker is library that we are using
	//  here os taking the both environment
	bot := slacker.NewClient(os.Getenv("GOLANG_SLACK_BOT_TOKEN"), os.Getenv("GOLANG_SLACK_APP_TOKEN"))

	go commandEvents(bot.CommandEvents())

	//  when you request any message it will respond
	bot.Command("Hi", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Hello")
		},
	})

	bot.Command("How are you? ", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("I am fine. How about you?")
		},
	})

	bot.Command("What can I help you", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Please tell me you number")
		},
	})

	bot.Command("8173xxxxxx is my number ", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("It's ok.")
		},
	})

	bot.Command("What is your qualification.", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {

			response.Reply("I have done MCA(Master of Computer Application)")
		},
	})

	bot.Command("What is your age", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("I am 24 years old.")
		},
	})

	bot.Command("What is your favourite game?", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("I love to play cricket.")
		},
	})

	bot.Command("Where are you from?", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("I live in Allahabad, Uttar Pradesh")
		},
	})

	bot.Command("Why Do you want to create Slack Bot", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Because it reply automatically you don't need to type manually.")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}

}
