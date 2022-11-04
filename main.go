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
	fmt.Println("To stop this program please enter: ;;")
	fmt.Print("Please enter the first title: ")
	
	// var binge_list *List

	reader := bufio.NewReader(os.Stdin)
	link := List{}

	for {
		input, err := reader.ReadString(';')
	
		if (err != nil) {
			fmt.Println(err)
		} else if (input == ";") {
			break;
		}
	
		link.Insert(input)
		link.Display()
		fmt.Print("Please add a new entry: ")
	}

	fmt.Println("\n===============================")
	fmt.Println("|Thank you for using Binge-Path!|")
	fmt.Println("===============================\n")
	
}


