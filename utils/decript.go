package utils

import (
	"strings"
)

// uses the message map to build the message in accordance with the map's keys
// return the message string
func buildMessageFromMap(mappedWords map[string]int) string {
	message := make([]string, len(mappedWords))
	for key, value := range mappedWords {
		if key != "" {
			message[value] = key
		}
	}
	return strings.Join(message, " ")
}

// input: el mensaje tal cual es recibido en cada satélite
// output: el mensaje tal cual lo genera el emisor del mensaje
func GetMessage(messages ...[]string) (msg string) {
	var mappedMessage = make(map[string]int)
	for _, message := range messages {
		var lastWord string
		for i, word := range message {
			if value, ok := mappedMessage[word]; ok {
				if value < i {
					mappedMessage[word] = i
					mappedMessage[lastWord] = i - 1
				}
			} else {
				mappedMessage[word] = i
			}
			lastWord = word

		}
	}
	msg = buildMessageFromMap(mappedMessage)
	return
}
