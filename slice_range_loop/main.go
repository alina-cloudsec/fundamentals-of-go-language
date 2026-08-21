package main
import "fmt"
func main() {
	statusCodes := []int{200, 404, 500, 403}
	
	for index, value := range statusCodes {
		fmt.Println("Index Position:", index)
		fmt.Println("Status Code:", value)
	}
}
