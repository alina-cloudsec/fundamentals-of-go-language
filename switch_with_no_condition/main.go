package main
import "fmt"
func main() {
	score := 85
	switch {
	case score >= 90:
		fmt.Println("A+ grade, Excellent performance")
	case score >= 80:
		fmt.Println("A grade, Very good")
	case score >= 70:
		fmt.Println("B grade, Good job")
	case score >= 60:
		fmt.Println("C grade, Needs improvement")
	default:
		fmt.Println("F grade, Failed")
	}
}

