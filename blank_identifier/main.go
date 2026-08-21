package main
import "fmt"
func getCoordinates() (int, int) {
	x := 45
	y := 82
	return x, y
}

func main() {
	posX, _ := getCoordinates() // This avoids the unused variables
	fmt.Printf("Current X Position: %d\n", posX)
}
