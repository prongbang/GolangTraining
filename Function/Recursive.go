package main
import "fmt"

func main() {

	fmt.Println("5! = ", factirial(5))

}

func factirial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factirial(x - 1)
}