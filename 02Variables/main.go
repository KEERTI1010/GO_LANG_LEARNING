package main

import "fmt"

const LoginToken string = "Hehaha" //public

func main() {
	var userName string = "KKS"
	fmt.Println(userName)
	fmt.Printf("Varibles is of type : %T \n ", userName)

	var isLogggedIn bool = false
	fmt.Println(isLogggedIn)
	fmt.Printf("Varibles is of type : %T \n ", isLogggedIn)

	var smallVal uint8 = 255 //cant do 256..bcz..uint8 liit is upto3256
	fmt.Println(smallVal)
	fmt.Printf("Varibles is of type : %T \n ", smallVal)

	var smallFloat float32 = 255.45555332 //255.45555
	fmt.Println(smallFloat)
	fmt.Printf("Variables is of type : %T \n", smallFloat)

	/*
		var smallFloat float64 = 255.45555332 //255.45555332
		fmt.Println(smallFloat)
		fmt.Printf("Variables is of type : %T \n", smallFloat)
	*/

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n ", anotherVariable)

	// implicit type

	var website = "google.com"
	fmt.Println(website)

	// no var style

	numOfUsers := 300000.0
	fmt.Println(numOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type %T \n", LoginToken)
}


/*

OUTPUT 


Varibles is of type : uint8 
255.45555
Variables is of type : float32 
0
Variable is of type : int 
google.com
300000
Hehaha
Variable is of type string 


*/