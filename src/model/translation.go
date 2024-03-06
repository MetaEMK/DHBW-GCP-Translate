package model

import (
	"crypto/sha512"
	"encoding/hex"

	"golang.org/x/text/language"
)

type TranslateRequest struct {
    Text            string          `json:"text"`
    TargetLanguage  language.Tag    `json:"targetLanguage"`
}

type TranslationResult struct {
    Text            string          `json:"text"`
    TargetLanguage  language.Tag    `json:"targetLanguage"`
    SourceLanguage  language.Tag    `json:"sourceLanguage"`
    Translation     string          `json:"translation"`
}

func (t *TranslateRequest) HashText() string {
    return hashString(t.Text)
}

func (t *TranslationResult) HashText() string {
    return hashString(t.Text)
}

func hashString(input string) string {
    hash := sha512.Sum512([]byte(input))
	hashString := hex.EncodeToString(hash[:])

    return hashString
}
