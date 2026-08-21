package main
import "fmt"
var currentenv = "Development"
func main() {
	userName := "Alina"
	userID := 9091
	valid := true

	fmt.Println("\t\t\tUSER PROFILE STATUS")
	fmt.Println("User:", userName)
	fmt.Println("ID:", userID)
	fmt.Println("Verification:", valid, ", Environment:", currentenv)
}
