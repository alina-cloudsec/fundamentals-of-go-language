package main
import "fmt"

func total(num ...int) int                          // Variadic function to add 
{
	sum:=0
	for _, val:=range num
  {
		sum=sum+val
	}
	return sum
}

func main() 
{
	var list []int
	var x int

	fmt.Println("\t\t\t Simple Go Calculator \t\t\t")
	fmt.Println("Type a number one by one and press Enter. To stop, just press ENTER:")

	for 
  {
		fmt.Print("Enter number: ")
		
		count, _ :=fmt.Scanln(&x)              //if user presses empty enter, count zero (here & memory operator) count is a counter
		if count == 0 
    {
			break
		}
		list = append(list, x)                           //to add anew no in the list (x ko list me add)
	}
  
  ans:=total(list...)
	fmt.Printf("Total Sum is: %d\n",ans)
}
