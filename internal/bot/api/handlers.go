package api

import (
	"fmt"
	"strings"
	"github.com/IliaSobolev/Torrseal/internal/bot/services"

	"gopkg.in/telebot.v3"
)

func HandleFile(ctx telebot.Context) error {
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

func HandleMagnetLink(ctx telebot.Context) error {
	magnetLink := ctx.Message().Text
	user := ctx.Message().Sender

	if !services.ValidateMagnetLink(magnetLink) {
		return ctx.Send("Incorrect file BTW")
	}

	// do some work with magnetLink

	// repository.SetUser(user, fileId) для ответа по скаченному файлу
	fmt.Println(user)

	return ctx.Send("File start download")
}
