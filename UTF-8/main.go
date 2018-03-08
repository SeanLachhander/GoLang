package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Decimal: %d :: Binary: %b :: Hexadecimal: %x \n", 42, 42, 42)
	fmt.Printf("Decimal: %d \t Binary: %b \t Hexadecimal (With 0X): %#X \n", 42, 42, 42)

	fmt.Println("\nNow, I will print out the first 200 characters in UTF-8")
	fmt.Printf("-------------------------------------------------------\n")

	for x := 0; x < 200; x++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", x, x, x, x)
	}
}
