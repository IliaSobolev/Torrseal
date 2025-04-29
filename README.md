# Torrseal

## Overview
Torrseal is a Telegram bot designed to simplify the process of downloading torrent files directly through Telegram. Users can send torrent links to the bot, and it will download the files and send back the results as a compressed archive.

## Navigation
- [How to Run](#how-to-run)
- [Warnings](#warnings)

## How to Run
To set up and run the Torrseal bot, follow these steps:

1. **Copy the environment file**:
   ```bash
   cp app.env.example .app.env
   ```
2. **Add your Telegram Bot Token**:
   Open the `.app.env` file and insert your Telegram bot token obtained from BotFather.
3. **Run the application**:
   Use Docker Compose to build and run the bot in detached mode:
   ```bash
   docker compose up --build -d
   ```

## ⚠️Warnings
Before running the bot, ensure you update the `DeveloperID` constant in the `utils/const.go` file. Replace the existing value with your own Telegram user ID to enable error message notifications via Telegram:
```go
const DeveloperID = "YOUR_TELEGRAM_USER_ID"
```
Failure to update this may result in restricted access or unintended behavior.
