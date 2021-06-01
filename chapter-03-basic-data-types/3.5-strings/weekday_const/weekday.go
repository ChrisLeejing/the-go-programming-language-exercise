package main

import "fmt"

type Weekday int
type ByteSize int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

const (
	_   = iota
	KiB = 1 << (10 * iota)
	MiB
	GiB
	TiB
	PiB
)

// refer: https://medium.com/swlh/iota-create-effective-constants-in-golang-b399f94aac31
// Sunday:  0
// Monday:  1
// Tuesday:  2
// Wednesday:  3
// Thursday:  4
// Friday:  5
// Saturday:  6
// -------------
// KB:  1024
// MB:  1048576
// GB:  1073741824
// TB:  1099511627776
// PB:  1125899906842624
// -------------
// KiB:  1024
// MiB:  1048576
// GiB:  1073741824
// TiB:  1099511627776
// PiB:  1125899906842624
func main() {
	fmt.Println("Sunday: ", Sunday)
	fmt.Println("Monday: ", Monday)
	fmt.Println("Tuesday: ", Tuesday)
	fmt.Println("Wednesday: ", Wednesday)
	fmt.Println("Thursday: ", Thursday)
	fmt.Println("Friday: ", Friday)
	fmt.Println("Saturday: ", Saturday)

	fmt.Println("-------------")

	fmt.Println("KB: ", KB)
	fmt.Println("MB: ", MB)
	fmt.Println("GB: ", GB)
	fmt.Println("TB: ", TB)
	fmt.Println("PB: ", PB)

	fmt.Println("-------------")

	fmt.Println("KiB: ", KiB)
	fmt.Println("MiB: ", MiB)
	fmt.Println("GiB: ", GiB)
	fmt.Println("TiB: ", TiB)
	fmt.Println("PiB: ", PiB)
}
