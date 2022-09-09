package main

import (
	// "context"

	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}

}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	app_token := os.Getenv("SLACK_APP_TOKEN")
	bot_token := os.Getenv("SLACK_BOT_TOKEN")

	bot := slacker.NewClient(bot_token, app_token)
	go printCommandEvents(bot.CommandEvents())

	bot.Command("my year of birth is: <year> ", &slacker.CommandDefinition{
		Description: "yob calculator",
		// Example:     "my year of birth is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			current_date := time.Now()
			current_year := current_date.Year()
			age := current_year - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
