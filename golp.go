package main

import "fmt"

func main() {

	fmt.Printf(`
-------------------------------- Golp -------
	`)

	for {

		userOptions()
		//-- get user choice
		var choice int 
		fmt.Print("Your selection: ")
		fmt.Scan( &choice ) 

		//--- condition
		if choice == 1 {
			fmt.Printf(`
- - - - - - - - - - - - - - - - - - - - - - - - - -
Go variables examples: string | int | float | bool

 data type inferred (integer)
  var x = 4    x := 4
        
 data type stated
  var x int = 4 

- - - - - - - - - - - - - - - - - - - - - - - - - -			
`)
		} else if choice == 9 {
			fmt.Print(" Exiting Golp \n")
			break
		} else {
			fmt.Print(" not an option - \n")
			return
		}



	} // end of for loop
	


}


func userOptions() {
	fmt.Printf(`
 what do you want to look up ?

 1. variables
 9. quit

 `)
}

