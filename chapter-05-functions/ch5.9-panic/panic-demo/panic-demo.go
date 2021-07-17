package main

import "fmt"

func main() {
	switch s := suit(); s {
	case "Club":
		fmt.Println("club 梅花")
	case "Diamond":
		fmt.Println("diamond 方块")
	case "Heart":
		fmt.Println("heart 红桃")
	case "Spade":
		fmt.Println("spade 黑桃")
	default:
		panic(fmt.Sprintf("invalid suit: %q", s))
	}
}

func suit() string {
	return drawCard()
}

// func drawCard() string {
// 	// "club 梅花；diamond 方块；heart 红桃；spade 黑桃；big joker 大王；little joker 小王"
// 	return "Club"
// }

func drawCard() string {
	// "club 梅花；diamond 方块；heart 红桃；spade 黑桃；big joker 大王；little joker 小王"
	return "joker"
}
