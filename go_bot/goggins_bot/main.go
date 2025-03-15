package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"goggins-bot/image_utils"

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

	var imageFiles []os.DirEntry
	for _, file := range files {
		if !file.IsDir() {
			ext := strings.ToLower(filepath.Ext(file.Name()))
			if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
				imageFiles = append(imageFiles, file)
			}
		}
	}

	if len(imageFiles) < count {
		log.Fatalf("Not enough images in the directory to pick %d images.", count)
	}

	rand.Shuffle(len(imageFiles), func(i, j int) { imageFiles[i], imageFiles[j] = imageFiles[j], imageFiles[i] })

	var images []string
	for i := 0; i < count; i++ {
		images = append(images, fmt.Sprintf("%s/%s", imageDir, imageFiles[i].Name()))
	}
	return images
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case "!ping":
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	case "!quote":
		quote := getRandomQuote()
		s.ChannelMessageSend(m.ChannelID, quote)
	case "!single":
		quote := getRandomQuote()
		images := getRandomImages("assets/images", 1)
		filePath := image_utils.GenerateSingleImageWithQuote(images[0], quote)
		sendImage(s, m.ChannelID, filePath)
	case "!window":
		quote := getRandomQuote()
		images := getRandomImages("assets/images", 4)
		filePath := image_utils.GenerateWindowImageWithQuote(images, quote)
		sendImage(s, m.ChannelID, filePath)
	case "!help":
		helpMsg := "**Goggins Bot Commands:**\n" +
			"- `!quote` - Get a random motivational quote\n" +
			"- `!single` - Generate a single image with a quote\n" +
			"- `!window` - Generate a 2x2 grid of images with a quote\n" +
			"- `!ping` - Check if the bot is running"
		s.ChannelMessageSend(m.ChannelID, helpMsg)
	default:
		// Either do nothing or respond with help message
		// s.ChannelMessageSend(m.ChannelID, "Unknown command. Try !help for available commands.")
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
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

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

	// Initialize fonts
	if err := image_utils.InitializeFonts(); err != nil {
		log.Printf("Warning: Failed to initialize fonts: %v", err)
		log.Println("Continuing with default fonts...")
	}

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

	fmt.Println("Goggins Bot is running. Press Ctrl+C to exit.")
	select {}
}
