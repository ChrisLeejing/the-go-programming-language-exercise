package main

import (
	"fmt"
	"math"
)

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	distance := Point.Distance   // 函数表达式
	fmt.Println(distance(p, q))  // 5
	fmt.Printf("%T\n", distance) // func(main.Point, main.Point) float64

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale) // func(*main.Point, float64)
}

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

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

type Path []Point

// 使用一个值来代表多个方法中的一个, 且方法都属于同一个类型, 此时可以借助方法变量, 来处理不同的接收者.
func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset) // 调用p[i].Add(offset), 或者p[i].Sub(offset).
	}
}
