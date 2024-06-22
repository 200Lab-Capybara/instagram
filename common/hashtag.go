package common

import (
	"fmt"
	"regexp"
)

func GetHashtag(content string) []string {
	re := regexp.MustCompile(`#([\p{L}\p{Mn}_]+)`)
	matches := re.FindAllString(content, -1)
	fmt.Println(matches, "hashtags")
	return matches
}
