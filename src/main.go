package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	godotenv.Load()
	ctx := context.Background()
	// Get Bot token from environment variables
	botToken := os.Getenv("TOKEN")

	// Create bot and enable debugging info
	// Note: Please keep in mind that default logger may expose sensitive information,
	// use in development only
	// (more on configuration in examples/configuration/main.go)
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Call method getMe (https://core.telegram.org/bots/api#getme)
	botUser, err := bot.GetMe(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Print Bot information
	fmt.Printf("Bot user: %+v\n", botUser)

	// Get updates channel
	// (more on configuration in examples/updates_long_polling/main.go)
	updates, _ := bot.UpdatesViaLongPolling(context.Background(), nil)

	// Loop through all updates when they came
	for update := range updates {
		fmt.Printf("Update: %+v\n", update)
		fmt.Printf("🤠 Message: %s \n", update.Message.Text)
		fmt.Printf("👽 ID: %d ", update.Message.Chat.ID) //1575433858 soy yo
		chatID := update.Message.Chat.ID

		// Call method sendMessage.
		// Send a message to sender with the same text (echo bot).
		// (https://core.telegram.org/bots/api#sendmessage)
		sentMessage, _ := bot.SendMessage(ctx,
			tu.Message(
				tu.ID(chatID), //1575433858 <- soy ese
				fmt.Sprintf("Hola miamors, soy una respuesta automatizada de tu amorcito ❤️. \n Te voy a dejar a continuación un número que wa necesitar después, así que ahí te lo encargo plox 🤠 Chat ID: %d", update.Message.Chat.ID),
			),
		)

		fmt.Printf("Sent Message: %v\n", sentMessage)
	}
}
