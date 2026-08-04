package main
import "fmt"
var currentenv = "Development"
func main() 
{
	userName:="Alina"
	userID:=9091
	valid:=true

	fmt.Println("\t\t\tUSER PROFILE STATUS")
	fmt.Printf("User: %s\n",userName)
  fmt.Printf("ID: %d\n",userID)
	fmt.Printfln("Verification: %t, Enviroment: %s\n",valid,currentenv)
}
