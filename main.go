package main
  
import (
	"fmt"
	"bufio"
	"os"
)

// Main function
func main() {
	fmt.Println("\n============================")
	fmt.Println("|Welcome to the Binge-Path!|")
	fmt.Println("============================\n")
	fmt.Print("Please enter the first title: ")
	
	// var binge_list *List

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	
	if (err != nil) {
		fmt.Println(err)
	}

	fmt.Println("The input is: " + input)


}


