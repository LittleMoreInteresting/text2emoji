package text2emoji

import (
	"encoding/base64"
	"math/rand"
	"strings"
)

func Encrypt(text string) string {
	base64Str := base64.StdEncoding.EncodeToString([]byte(text))
	emojiLen := len(emojis)
	randomInt := rand.Intn(10)
	var encrypted strings.Builder
	encrypted.WriteString(emojis[randomInt])
	for _, char := range base64Str {
		idx := (randomInt + base64CharIdx[string(char)]) % emojiLen
		encrypted.WriteString(emojis[idx])
	}
	return encrypted.String()
}

func Decrypt(emojiStr string) string {
	var base64Str strings.Builder
	emojiLen := len(emojis)
	var step int
	for i, emoji := range emojiStr {
		if i == 0 {
			step = emojiIdx[string(emoji)]
			continue
		}
		ei := emojiIdx[string(emoji)]
		char := base64Char[((ei-step)%emojiLen+emojiLen)%emojiLen]
		base64Str.WriteString(string(char))
	}
	decoded, _ := base64.StdEncoding.DecodeString(base64Str.String())
	return string(decoded)
}
