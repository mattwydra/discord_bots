# Goggins Bot

A Discord bot that generates motivational images with inspirational quotes. Named after David Goggins, a former Navy SEAL known for his motivational messages about mental toughness and perseverance.

## Overview

Goggins Bot listens for specific commands in Discord channels and responds with either plain text quotes or generates images with overlaid motivational quotes. It can create single images or a 2x2 grid of images with a quote centered on them.

## Features

- Random motivational quotes (`!quote`)
- Generate a single image with a motivational quote overlay (`!single`)
- Generate a 2x2 grid of images with a quote overlay (`!window`)
- Basic help command (`!ping` and `!help`)

## Project Structure

```
goggins-bot/
├── assets/
│   ├── fonts/         # Downloaded fonts will be stored here
│   ├── images/        # Place your motivational images here (.jpg, .jpeg, .png)
│   ├── quotes.txt     # Short motivational quotes
│   └── long_quotes.txt # Longer motivational quotes
├── image_utils/
│   ├── font_manager.go # Font management utilities
│   └── image_utils.go  # Image processing utilities
├── output/            # Generated images will be saved here
├── config.json        # Discord bot configuration
├── go.mod             # Go module definition
├── go.sum             # Go module checksums
└── main.go            # Main bot code
```

## Dependencies

The bot uses the following Go packages:
- [github.com/bwmarrin/discordgo](https://github.com/bwmarrin/discordgo) - Discord API wrapper
- [github.com/fogleman/gg](https://github.com/fogleman/gg) - 2D graphics library

## Setup and Installation

### Prerequisites

1. Go 1.20 or later
2. A Discord Bot Token (create one at [Discord Developer Portal](https://discord.com/developers/applications))
3. Proper folder structure with images in the `assets/images` directory

### Installation Steps

1. Clone this repository or download the source code

2. Create the necessary directories:
```sh
mkdir -p assets/fonts assets/images output
```

3. Add images to the `assets/images` directory:
   - Make sure there are at least 4 images for the `!window` command
   - Supported formats: JPG, JPEG, PNG

4. Create or edit the `config.json` file with your Discord bot token:
```json
{
  "token": "YOUR_DISCORD_BOT_TOKEN"
}
```

5. Install the dependencies:
```sh
go mod download
```

6. Build and run the bot:
```sh
go build
./goggins-bot  # On Windows: goggins-bot.exe
```

Alternatively, you can directly run without building:
```sh
go run main.go
```

## Usage

Once the bot is running and added to your Discord server, you can use the following commands:

- `!ping` - Check if the bot is running
- `!quote` - Get a random motivational quote
- `!single` - Generate a single image with a quote
- `!window` - Generate a 2x2 grid of images with a quote
- `!help` - Show available commands

## Troubleshooting

### Common Issues

1. **Bot not starting**
   - Ensure your `config.json` file exists and contains a valid Discord bot token
   - Check for any error messages in the console

2. **Font loading errors**
   - The bot will automatically download needed fonts on first run
   - Ensure your system has internet access for the initial font download
   - If font download fails, manually download fonts and place them in `assets/fonts` directory

3. **Image generation errors**
   - Ensure you have at least one image for `!single` and four images for `!window` in `assets/images`
   - Check file permissions on the `output` directory

4. **Bot can connect but commands don't work**
   - Ensure the bot has proper permissions in your Discord server
   - Check that quotes are properly loaded from `assets/quotes.txt`

### Debugging Tips

- Add more logging statements to `main.go` to track command execution
- Test image generation functions separately from Discord integration
- Ensure all required directories exist with proper permissions

## Customization

### Adding More Quotes

Simply add more quotes to the `assets/quotes.txt` file, with one quote per line.

### Changing Image Styles

Modify the constants in `image_utils/image_utils.go` to adjust:
- Image dimensions (`StandardWidth`, `StandardHeight`)
- Text styling (`FontSize`, `LineSpacing`, `MaxLineWidth`)
- Overlay styling (`ShadowOffset`, `TextPadding`)

### Adding Custom Fonts

Place custom TTF or OTF font files in the `assets/fonts` directory and update the font loading code in `image_utils/font_manager.go` if needed.

## Acknowledgments

- David Goggins for the motivational quotes and inspiration
- The creators of [discordgo](https://github.com/bwmarrin/discordgo) and [gg](https://github.com/fogleman/gg) libraries