package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Result struct {
}

// getSearchQuery - takes the user provide input and returns the formatted query string
func getSearchQuery() string {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Q: ")
	searchQuery, _ := reader.ReadString('\n')
	searchQuery = searchQuery[:len(searchQuery)-1]
	return strings.ReplaceAll(searchQuery, " ", "+")
}

func main() {

}
