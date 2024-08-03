package i18n

import (
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

const (
	TR = "tr"
	EN = "en"
)

func InitBundle(lang string) {
	bundle = i18n.NewBundle(language.Turkish)

	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	var languages = []string{
		lang + "/en.json",
		lang + "/tr.json",
	}

	for _, lang := range languages {
		bundle.MustLoadMessageFile(lang)
	}
}
