package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welccome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for Our Pizza : ")

	// comma ok syntax => err ok syntax

	input, _ := reader.ReadString('\n')
	// _,err := reader.ReadString('\n') -> if error is there we can use this
	fmt.Println("Thankss for Rating , ", input)
	fmt.Printf("Type of rating is %T ", input)
}
/*
Welccome to user input
Enter the rating for Our Pizza : 
9
Thankss for Rating ,  9

Type of rating is string 
*/
