package main
import "fmt"
func main() 
{
	counter := 1
	for 
  {
		fmt.Printf("Loop iteration number: %d\n", counter)
		 if counter >= 5 
      {
			  fmt.Println("Target reached! Breaking out of the loop.")
			  break 
		  }
		counter++ 
	}
	fmt.Println("Program finished.")
}
