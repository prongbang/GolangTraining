package main
import "fmt"


func main() {

	// คือฟังก์ชันที่ไม่จำเป็นต้องมีชื่อ
	add := func(x, y int) int{
		return x + y
	}

	fmt.Println("1 + 2 = ",add(1,2))

}