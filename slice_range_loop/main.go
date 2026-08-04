package main
import "fmt"
func main() 
{
	statusCodes := []int{200, 404, 500, 403}                            //dynamic slice
  
	for index, value := range statusCodes                       // Using 'range' to extract both index and value cleanly
  {
		fmt.Printfln("Index Position: %d",index)
    fmt.Printfln("Status Code: %d",value)
	}
}
