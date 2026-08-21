package main
import "fmt"
var x int
var p *int

func change() {
	*p = 100                             // 3. Go inside the address and change value to 100
}

func main() {
	x = 20
	fmt.Println("Before:", x)
	p = &x                             //p stores this memory address of x
	change()                           // Call function to change the value
	fmt.Println("After:", x)
}
