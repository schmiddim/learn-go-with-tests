package iteration

import "strings"

//https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration

func Repeat(char string, repeat int) (result string) {

	result = strings.Repeat(char, repeat)
	return
}
