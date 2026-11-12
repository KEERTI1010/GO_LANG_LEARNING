package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Our Pizza App")
	fmt.Println("Please rate our Pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for Rating,", input)

	//	numRating, err := strconv.ParseFloat(input, 64) => strconv.ParseFloat: parsing "3\r\n": invalid syntax
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
		// panic(err)
	} else {
		fmt.Println("Added 1 to ur rating", numRating+1)
	}
}

/*
Welcome to Our Pizza App
Please rate our Pizza between 1 and 5
3
Thanks for Rating, 3

Added 1 to ur rating 4
*/
