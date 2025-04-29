package bot

import (
	"fmt"
	"strings"

	"gopkg.in/telebot.v3"
)

func handleFile(ctx telebot.Context) error {
	inputFile := ctx.Message().Document.File
	user := ctx.Message().Sender

	if !strings.HasSuffix(inputFile.FileLocal, ".torrent") {
		return ctx.Send("Incorrect file BTW")
	}

	// do some work with file

	// repository.SetUser(user, fileId) для ответа по скаченному файлу
	fmt.Println(user)

	return ctx.Send("File start download")
}

func handleMagnetLink(ctx telebot.Context) error {
	magnetLink := ctx.Message().Text
	user := ctx.Message().Sender

	if !validateMagnetLink(magnetLink) {
		return ctx.Send("Incorrect file BTW")
	}

	// do some work with magnetLink

	// repository.SetUser(user, fileId) для ответа по скаченному файлу
	fmt.Println(user)

	return ctx.Send("File start download")
}
