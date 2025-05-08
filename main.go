package main


import (
        "fmt"
        "strings"
				"bufio"
				"os"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		answer := scanner.Text()
		words := cleanInput(answer)
		if len(words) > 0 {
			fmt.Printf("Your command was: %s\n", words[0])
		}
	}
}


func cleanInput(text string) []string {
  s := strings.TrimSpace(text)
  lowerChars := strings.ToLower(s)
  words := strings.Fields(lowerChars)
  return words
}

