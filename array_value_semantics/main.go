package main
import "fmt"
func main() 
{
	originalScores := [3]int{90, 85, 95}                       //array
	copiedScores := originalScores
	copiedScores[0] = 100

	fmt.Println("--- Array Copy Test ---")
	fmt.Printf("Original Array: %v\n", originalScores)               //outputOriginal Array: [90 85 95]
	fmt.Printf("Copied Array:   %v\n", copiedScores)                  //Copied Array:   [100 85 95]
}
