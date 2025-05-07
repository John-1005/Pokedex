package main


import (
        "fmt"
        "strings"
)


func main() {
	fmt.Println("Hello, World!")
}


func cleanInput(text string) []string {
  s := strings.TrimSpace(text)
  lowerChars := strings.ToLower(s)
  words := strings.Fields(lowerChars)
  return words
}

