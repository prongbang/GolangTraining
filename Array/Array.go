package main

import "fmt"

func main() {

	var x1[5]int
	x1[0] = 1

	fmt.Println("len : ", len(x1),", data : ", x1) // [1 0 0 0 0]

	var x2[5][2]int
	x2[0][0] = 2

	fmt.Println("len : ", len(x2),", len(x2[0]) : ", len(x2[0]),", data : ", x2) // [[2 0] [0 0] [0 0] [0 0] [0 0]]

	const size = int(len(x2))
	fmt.Println("size : ", size)

	// For range
	// _ คือ บอก compiler ว่าเราไม่ต้องการใช้ตัวแปรนั้น
	var total int = 0
	for _,value := range x1 {
		total += value
	}
	fmt.Println("total : ",total)


	// กำหนดค่าเริ่มแรก
	x := [5]float32{
		1,2,3,4,5, // ต้องเติม , ต่อหลังด้วย
	}
	fmt.Println("Array Fixed Data : ", x)

}
