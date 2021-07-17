package main

import "fmt"

func main() {
	switch s := suit(); s {
	case "club":
		fmt.Println("club 梅花")
	case "diamond":
		fmt.Println("diamond 方块")
	case "heart":
		fmt.Println("heart 红桃")
	case "spade":
		fmt.Println("spade 黑桃")
	default:
		fmt.Printf("not suit: %s\n", s)
	}
}

// "club 梅花；diamond 方块；heart 红桃；spade 黑桃；big joker 大王；little joker 小王"
func suit() string {
	return drawCard()
}

// func drawCard() string {
// 	return "club"
// }

func drawCard() string {
	return "joker"
}
