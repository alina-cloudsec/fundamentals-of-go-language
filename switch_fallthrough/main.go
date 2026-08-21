package main
import "fmt"
func main() {
	userTier := "Gold"
	fmt.Println("Checking unlocked rewards for your account:")

	switch userTier {
	case "Platinum":
		fmt.Println("Unlocked the Premium Lounge Access")
		fallthrough                                          // Forces the code to run the next case as well
	case "Gold":
		fmt.Println("Unlocked the 20% Discount Code")
		fallthrough
	case "Silver":
		fmt.Println("Unlocked the Standard Customer Support")
	default:
		fmt.Println("Unlocked: Basic Profile Features")
	}
}
