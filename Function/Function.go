package main
import "fmt"

func main() {

	xs := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	avg := average(xs)
	fmt.Println("avg : ", avg)
	fmt.Println(f())
	s, e := f2()
	fmt.Println(s, e)
	fmt.Println("total : ", add(1, 2, 3, 4, 5))

	x := []int{1,2,3,4,5}
	fmt.Println("total : ", add(x...))	// ส่งข้อมูลไปแบบ slice โดยเติม ...

}

/**
 * @func average
 * @commect name(param type) return type
 */
func average(xs []float64) float64 {
	fmt.Println(xs)
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// การกำหนดค่าตัวแปรที่ต้องการจะ return ได้
// r int คือ type ที่ return
func f() (r int) {
	r = 1        // กำหนดค่าให้ return
	return
}

// return หลายค่า มักจะนำไปใช้ส่งข้อมูลความผิดพลาดของ function
func f2() (int, bool) {
	// ค่า , error
	return 5, true
}

// ฟังก์ชันรับค่าแปบไม่จำกัดเหมือน fmt.Pringln(,,)
func add(args ...int) int {
	fmt.Println(args)
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}