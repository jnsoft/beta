package stringutil

import (
	"encoding/base64"
	"encoding/hex"
	"strings"
)

func ToBase64(data []byte, lineBreak int) string {
	encoded := base64.StdEncoding.EncodeToString(data)
	if lineBreak > 0 {
		return insertLineBreaks(encoded, lineBreak)
	}
	return encoded
}

func FromBase64(data string) ([]byte, error) {
	cleanedData := strings.ReplaceAll(data, "\n", "")
	return base64.StdEncoding.DecodeString(cleanedData)
}

func ToHex(data []byte, lineBreak int) string {
	encoded := hex.EncodeToString(data)
	if lineBreak > 0 {
		return insertLineBreaks(encoded, lineBreak)
	}
	return encoded
}

func FromHex(data string) ([]byte, error) {
	cleanedData := strings.ReplaceAll(data, "\n", "")
	return hex.DecodeString(cleanedData)
}

func insertLineBreaks(s string, n int) string {
	var result string
	for i := 0; i < len(s); i += n {
		if i+n < len(s) {
			result += s[i:i+n] + "\n"
		} else {
			result += s[i:]
		}
	}
	return result
}
