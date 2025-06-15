package utils

import (
	"fmt"
	"regexp"
)

func Regexp() {
	var pattern *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	isMatch := pattern.MatchString("eko")
	fmt.Println(isMatch)

	findString := pattern.FindAllString("eko ego eki eSo emo ejo eao", 5)
	fmt.Println(findString)
}
