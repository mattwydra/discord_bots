package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"

	// "io/ioutil"
	"log"

	"math/rand"
	"os"
	"strings"

	"goggins-bot/image_utils"

	// "time"

	"github.com/bwmarrin/discordgo"
)

var quotes []string

func loadQuotes(filePath string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading quotes file: %v", err)
	}
	quotes = strings.Split(string(content), "\n")
}

func getRandomQuote() string {
	return quotes[rand.Intn(len(quotes))]
}

func getRandomImages(imageDir string, count int) []string {
	files, err := os.ReadDir(imageDir)
	if err != nil {
		log.Fatalf("Error reading image directory: %v", err)
	}

	if len(files) < count {
		log.Fatalf("Not enough images in the directory to pick %d images.", count)
	}

	rand.Shuffle(len(files), func(i, j int) { files[i], files[j] = files[j], files[i] })

	var images []string
	for i := 0; i < count; i++ {
		images = append(images, fmt.Sprintf("%s/%s", imageDir, files[i].Name()))
	}
	return images
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Example: Respond to a specific command
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	switch m.Content {
	case "!quote":
		quote := getRandomQuote()
		s.ChannelMessageSend(m.ChannelID, quote)
	case "!single":
		quote := getRandomQuote()
		images := getRandomImages("assets/images", 1)
		filePath := image_utils.GenerateSingleImageWithQuote(images[0], quote)
		sendImage(s, m.ChannelID, filePath)
	// case "!single":
	// 	quote := getRandomQuote()
	// 	images := getRandomImages("assets/images", 1)
	// 	filePath := generateSingleImageWithQuote(images[0], quote)
	// 	sendImage(s, m.ChannelID, filePath)

	case "!window":
		quote := getRandomQuote()
		images := getRandomImages("assets/images", 4)
		filePath := image_utils.GenerateWindowImageWithQuote(images, quote)
		sendImage(s, m.ChannelID, filePath)

	default:
		s.ChannelMessageSend(m.ChannelID, "Get back to work. Try !quote, !single, or !window.")
	}
}

func sendImage(s *discordgo.Session, channelID, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return
	}
	defer file.Close()

	s.ChannelFileSend(channelID, "motivation.png", file)
}

func main() {
	// Replace with your bot token
	// Load config
	config := struct {
		Token string `json:"token"`
	}{}
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer configFile.Close()

	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}

	// Load quotes
	loadQuotes("assets/quotes.txt")

	// Start bot
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}
	defer dg.Close()

	dg.AddHandler(messageHandler)

	if err := dg.Open(); err != nil {
		log.Fatalf("Error opening Discord session: %v", err)
	}

	fmt.Println("Bot is running. Press Ctrl+C to exit.")
	select {}
}

// func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
// 	// Ignore bot's own messages
// 	if m.Author.ID == s.State.User.ID {
// 		return
// 	}

// 	// Example: Respond to a specific command
// 	if m.Content == "!ping" {
// 		s.ChannelMessageSend(m.ChannelID, "Pong!")
// 	}
// }
