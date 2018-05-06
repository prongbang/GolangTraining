package main
import (
	"fmt"
	"sort"
)


func main() {

	// #Map
	x := make(map[string]int)
	fmt.Println(x)	// map[]
	// add
	x["key"] = 100
	x["data"] = 200
	fmt.Println(x) // map[key:100]
	fmt.Println(x["key"]) // 100
	// delete
	delete(x,"key")
	fmt.Println(x)	// map[]

	// ตรวจสอบว่ามีค่ามั้ย
	name, ok := x["data"]
	fmt.Println(name,ok) // 200 true
	name2, ok2 := x["key"]
	fmt.Println(name2,ok2) // 0 false

	// การเข้าถึงค่าของ map
	if name3, ok3 := x["data"]; ok3{
		fmt.Println("if : ",name3, ok3) // 200 true
	}

	if name3, ok3 := x["key"]; ok3{
		fmt.Println(name3, ok3)
	} else {
		fmt.Println("else : ", name3, ok3) // 0 false
	}

	// # การเขียน map ให้สั้นลง
	m := map[string]string{
		"A":"AA",
		"B":"BB",
		"C":"CC",
		"E":"EE",
		"D":"DD", // ** อย่าลืม ,
	}
	fmt.Println(m)	// map[A:AA B:BB C:CC E:EE D:DD]

	// map ซ้อน map
	m2m := map[string]map[string]string{
		"A":map[string]string{
			"1" : "A",
			"2" : "A",
		},
		"B":map[string]string{
			"1" : "B",
			"2" : "B",
		},
		"C":map[string]string{
			"1" : "C",
			"2" : "C",
		},
		"E":map[string]string{
			"1" : "E",
			"2" : "E",
		},
		"D":map[string]string{
			"1" : "D",
			"2" : "D",
		},
	}
	fmt.Println(m2m) // map[C:map[1:C 2:C] E:map[1:E 2:E] D:map[2:D 1:D] A:map[1:A 2:A] B:map[1:B 2:B]]


	// find min value in array
	xx := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17, }
	sort.Ints(xx)
	fmt.Println("min : ", xx[0]) // 9

}