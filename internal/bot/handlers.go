package bot

import (
	"fmt"
	"regexp"
	"strings"

	"gopkg.in/telebot.v3"
)

func handleFile(ctx telebot.Context) error {

	inputFile := ctx.Message().Document.File
	user := ctx.Message().Sender

	if !strings.HasSuffix(inputFile.FileLocal, ".torrent") {
		return ctx.Send("Incorrect file BTW")
	}

	//do some with file

	//repository.SetUser(user, fileId) для ответа по скаченному файлу
	fmt.Println(user)

	return ctx.Send("File start download")
}

func handleMagnetLink(ctx telebot.Context) error {
	magnetLink := ctx.Message().Text
	user := ctx.Message().Sender

	if !validateMagnetLink(magnetLink) {
		return ctx.Send("Incorrect file BTW")
	}

	//do some with magnetLink

	//repository.SetUser(user, fileId) для ответа по скаченному файлу
	fmt.Println(user)

	return ctx.Send("File start download")
}

// это надо вынести в сервисы хэндлеров что бы тут были только функции отвечающие за поведение в ответ по энтпойнту
// а валилдация может быть где-то в тулзах или ещё чём
func validateMagnetLink(link string) bool {
	// idk work or not this patterns
	magnetPattern := []string{
		`^magnet:\?xt=urn:btih:[a-zA-Z0-9]{32,40}(&dn=.+)?(&tr=.+)*$`,
		`^magnet:\?xt=urn:(btih|sha1|sha256):[a-zA-Z0-9]{32,64}(&dn=.+)?(&tr=.+)*$`,
		`^magnet:\?(xt=urn:(btih|sha1|sha256|md5):[a-zA-Z0-9]{32,64}|dn=.+|kt=.+|tr=.+|xs=.+|as=.+|mt=.+)+$`,
	}
	for _, pattern := range magnetPattern {
		re := regexp.MustCompile(pattern)
		if re.MatchString(link) {
			return true
		}
	}
	return false
}
