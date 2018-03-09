package main

import(
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main(){

	// Ask user for a sentence
	fmt.Println("Enter the sentence you would like to make an acronym out of!")

	// Store the sentence into a scanner object.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sentence := scanner.Text()

	// Split the sentence into individual words separated by the space character.
	words := strings.Split(sentence, " ")

	// Declare a string array to store the acronym in.
	new_acronym := []string{}

	// Take the first letter of each word in the sentence, and append it to the variable: "new_acronym"
	for i:=0; i<len(words); i++{
		new_acronym = append(new_acronym, words[i][0:1])
	}

	// Store the acronym in result, and display in console.
	result := strings.Join(new_acronym, "")
	fmt.Println("The acronym for " + sentence + " is: " + result)

}