package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "abc"
	s2 := strings.Clone(s1)
	b := strings.Builder{}

	b.WriteString("Hello darkness!")

	fmt.Println(b.String())
	fmt.Println(strings.ToLower(s1))
	fmt.Println(strings.ToUpper(s2))
	fmt.Println(strings.TrimSpace(s1))

	fmt.Println(strings.HasSuffix("test@gmail.com", "@gmail.com"))
}
