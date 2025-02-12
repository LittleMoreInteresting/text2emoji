package text2emoji

import (
	"encoding/base64"
	"fmt"
	"strings"
)

var (
	char2emoji = map[string]string{
		"A": "😀",
		"B": "😃",
		"C": "😄",
		"D": "😁",
		"E": "😆",
		"F": "😅",
		"G": "😂",
		"H": "🤣",
		"I": "😊",
		"J": "😇",
		"K": "🐶",
		"L": "🐱",
		"M": "🐭",
		"N": "🐹",
		"O": "🐰",
		"P": "🦊",
		"Q": "🐻",
		"R": "🐼",
		"S": "🐨",
		"T": "🐯",
		"U": "🍎",
		"V": "🍐",
		"W": "🍊",
		"X": "🍋",
		"Y": "🍌",
		"Z": "🍉",
		"a": "🍇",
		"b": "🍓",
		"c": "🍈",
		"d": "🍒",
		"e": "⚽",
		"f": "🏀",
		"g": "🏈",
		"h": "⚾",
		"i": "🎾",
		"j": "🏐",
		"k": "🏉",
		"l": "🎱",
		"m": "🏓",
		"n": "🏸",
		"o": "🚗",
		"p": "🚕",
		"q": "🚙",
		"r": "🚌",
		"s": "🚎",
		"t": "🏎",
		"u": "🚓",
		"v": "🚑",
		"w": "🚒",
		"x": "🚐",
		"y": "⌚",
		"z": "📱",
		"0": "📲",
		"1": "💻",
		"2": "⭐",
		"3": "🐈",
		"4": "🐕",
		"5": "🐖",
		"6": "🐪",
		"7": "💾",
		"8": "💰",
		"9": "💴",
		"+": "💵",
		"/": "💶",
		"=": "💷",
	}
	emoji2char = map[string]string{
		"😀": "A",
		"😃": "B",
		"😄": "C",
		"😁": "D",
		"😆": "E",
		"😅": "F",
		"😂": "G",
		"🤣": "H",
		"😊": "I",
		"😇": "J",
		"🐶": "K",
		"🐱": "L",
		"🐭": "M",
		"🐹": "N",
		"🐰": "O",
		"🦊": "P",
		"🐻": "Q",
		"🐼": "R",
		"🐨": "S",
		"🐯": "T",
		"🍎": "U",
		"🍐": "V",
		"🍊": "W",
		"🍋": "X",
		"🍌": "Y",
		"🍉": "Z",
		"🍇": "a",
		"🍓": "b",
		"🍈": "c",
		"🍒": "d",
		"⚽": "e",
		"🏀": "f",
		"🏈": "g",
		"⚾": "h",
		"🎾": "i",
		"🏐": "j",
		"🏉": "k",
		"🎱": "l",
		"🏓": "m",
		"🏸": "n",
		"🚗": "o",
		"🚕": "p",
		"🚙": "q",
		"🚌": "r",
		"🚎": "s",
		"🏎": "t",
		"🚓": "u",
		"🚑": "v",
		"🚒": "w",
		"🚐": "x",
		"⌚": "y",
		"📱": "z",
		"📲": "0",
		"💻": "1",
		"⭐": "2",
		"🐈": "3",
		"🐕": "4",
		"🐖": "5",
		"🐪": "6",
		"💾": "7",
		"💰": "8",
		"💴": "9",
		"💵": "+",
		"💶": "/",
		"💷": "=",
	}
)

func Encrypt(text string) string {
	base64Str := base64.StdEncoding.EncodeToString([]byte(text))
	fmt.Println(base64Str)
	var encrypted strings.Builder
	for _, char := range base64Str {
		encrypted.WriteString(char2emoji[string(char)])
	}
	return encrypted.String()
}

func Decrypt(emojiStr string) string {
	var base64Str strings.Builder
	for _, emoji := range emojiStr {
		char, ok := emoji2char[fmt.Sprintf("%c", emoji)]
		if !ok {
			return ""
		}
		base64Str.WriteString(char)
	}
	decoded, _ := base64.StdEncoding.DecodeString(base64Str.String())
	return string(decoded)
}
