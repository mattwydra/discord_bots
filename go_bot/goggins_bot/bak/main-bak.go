package mainbak

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var quotes []string

func loadQuotes(filePath string) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading quotes file: %v", err)
	}
	quotes = strings.Split(string(content), "\n")
}

func getRandomQuote() string {
	return quotes[rand.Intn(len(quotes))]
}

func getRandomImages(imageDir string, count int) []string {
	files, err := ioutil.ReadDir(imageDir)
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

	switch m.Content {
	case "!quote":
		quote := getRandomQuote()
		s.ChannelMessageSend(m.ChannelID, quote)

	case "!single":
		quote := getRandomQuote()
		images := getRandomImages("assets/images", 1)
		filePath := generateSingleImageWithQuote(images[0], quote)
		sendImage(s, m.ChannelID, filePath)

	case "!window":
		quote := getRandomQuote()
		images := getRandomImages("assets/images", 4)
		filePath := generateWindowImageWithQuote(images, quote)
		sendImage(s, m.ChannelID, filePath)

	default:
		s.ChannelMessageSend(m.ChannelID, "Unknown command. Try !quote, !single, or !window.")
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
	// New(NewSource(time.Now().UnixNano()))

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
