package main
import "fmt"

func main() {

	/// ความแตกต่างของ slice กับ array คือ ไม่มีการกำหนดความยาวของ array

	var x[]float32;
	fmt.Println(x) // []

	/**
	 ควรใช้ make สร้าง slice
	 */
	// 5 คือ ควายาวของ array
	x1 := make([]float64, 5) // [0 0 0 0 0]
	fmt.Println(x1)

	// 10 คือ ความยาวของ array ที่ slice อ้างถึง
	x2 := make([]float64, 5, 10)
	fmt.Println(x2)// [0 0 0 0 0]


	// สร้าง slice โดยใช้นิพจน์ [low:high]
	arr := [5]float64{1, 2, 3, 4, 5}
	x3 := arr[0:3]	// [1 2 3]
	x4 := arr[0:len(arr)]
	x5 := arr[:]
	x6 := arr[:5]
	fmt.Println(x3, x4, x5, x6) // [1 2 3]


	// ฟังก์ชันจัดการกับ slice มี append และ copy

	// #append
	s1 := []int{1,2}
	s2 := append(s1,3,4,5,6)
	fmt.Println(s2) // [1 2 3 4 5 6]

	// #copy
	s3 := make([]int,3)
	copy(s3,s1) // copy ข้อมูลเอามาใส่ s3
	fmt.Println(s3, s1) // [1 2 0] [1 2]

}