package main
import (
	"math"
	"fmt"
)

func main() {

	// 9.1 Structs
	c := new(Circle)
	c.r = 40
	fmt.Println("function : ",areaCircle(c))
	fmt.Println("method : ",c.area())

	r := new(Rectangle)
	r.x1 = 30
	r.y1 = 232
	fmt.Println("Rect : ",r.area())
}

type Circle struct  {
	x float64
	y float64
	r float64
}

type Rectangle struct  {
	x1,y1,x2,y2 float64
}

func areaCircle(c *Circle) float64{
	return math.Pi * c.r * c.r
}


// Method ของ Circle
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// method rectangle
func (r *Rectangle) area() float64 {

	return r.x1 * r.y1
}