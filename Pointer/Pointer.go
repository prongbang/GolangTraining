package main

import "fmt"

func main() {

	x := 100
	var y int = 200

	var iop *int // ประกาศตัวแปร iop ที่เป็น pointer ไว้เก็บ memory address ของ ตัวแปรที่เก็บข้อมูลชนิด int
	poi := &x    // x ส่ง memory address ให้กับ poi โดย poi จะ มี type เป็นไปตามตัวแปรที่ส่ง address มาให้
	iop = &y

	one(iop)
	fmt.Println("iop(address y) =", iop)

	zero(poi)
	fmt.Println("x =", x)
	fmt.Println("x(address poi) =", poi)

	// การใช้ new
	xPttr := new(int) // ส่ง memory address ของ new(int) ไปให้ xPttr
	one(xPttr)
	*xPttr = 1
	fmt.Println("xPttr(address new(int)) =", xPttr)
	fmt.Println("xPttr =", *xPttr)

	// สรุป
	// *type และ new(type) คือการประกาศ pointer เพื่อเก็บ memory address
	// &variable ดึง memory address ของตัวแปรนั้น ๆ
	// *variable ดึง value ของ memory address นั้น ๆ
	// *variable = ค่าใหม่ คือการกำหนดค่าใหม่ให้กับตัวแปรผ่าน memory address

}

func zero(xPttr *int) {
	*xPttr = 0 // คือการกำหนดค่าใหม่ให้กับตัวแปรที่ส่ง memory address มาให้
}

func one(xPttr *int) {
	*xPttr = 1 // คือการกำหนดค่าใหม่ให้กับตัวแปรที่ส่ง memory address มาให้
}
