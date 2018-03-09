/* Sean Lachhander
   GoLang
   This program will ask the user for their name, accept user input, then greet and introduce itself.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* This function will generate a random number.
   @param: min - minimum int range.
   @param: max - maximum int range.
*/
func randomNumber(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func main() {
	var name string // store name
	names := [10]string{"Sean", "Bill", "Ian", "Bob", "PC", "Mac", "Apple", "Tom", "Kris", "Tony"}
	rand_num := randomNumber(0, len(names)-1)

	/* Interact with user, and ask for name */
	fmt.Println("Hello, welcome to GoLang. What is your name?")
	fmt.Scanf("%s", &name)

	// Interact with user
	fmt.Println("Hello " + name + ", my name is " + names[rand_num] + ". It is nice to meet you!")
}
