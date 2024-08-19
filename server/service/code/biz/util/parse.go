package util

import (
	"strings"
)

func Parse(text string) (string, string) {
	texts := strings.SplitN(text, "\n", 2)
	if len(texts) != 2 {
		return "", ""
	}
	lan := strings.SplitN(texts[0], " ", 2)
	if len(lan) != 2 {
		return "", texts[1]
	}
	return strings.TrimSpace(lan[1]), texts[1]
}
