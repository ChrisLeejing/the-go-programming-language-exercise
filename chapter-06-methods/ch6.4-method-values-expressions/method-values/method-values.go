package main

import (
	"fmt"
	"math"
	"time"
)

type Point struct {
	X, Y float64
}

// function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Point类型的方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type Rocket struct {
}

func (r *Rocket) Launch() {

}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFromP := p.Distance   // 方法变量
	fmt.Println(distanceFromP(q)) // 5

	var origin Point
	fmt.Println(distanceFromP(origin)) // 2.23606797749979. (0, 0) 到 (1, 2) 的距离

	r := new(Rocket)
	time.AfterFunc(10*time.Second, r.Launch) // 方法变量
	time.AfterFunc(10*time.Second, func() {
		r.Launch() // 常规写法
	})

	scaleP := p.ScaleBy
	scaleP(2)
	scaleP(3)
	scaleP(10)
	fmt.Println(p)
}
