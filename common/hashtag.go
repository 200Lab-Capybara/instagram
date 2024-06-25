package common

import (
	"regexp"
)

func GetHashtag(content string) []string {
	re := regexp.MustCompile(`#([\p{L}\p{Mn}_]+)`)
	matches := re.FindAllString(content, -1)
	return matches
}
