package main

import (
	"fmt"
	"hex"
)

func main() {
	var content = map[int]string{
		1: "Fine",
		2: "Cloudy",
		3: "Rainy",
		4: "Stormy",
		5: "Cold",
		6: "Sunny",
		7: "Hail",
		8: "Stormy Seas",
		9: "Misty",
		10: "Monsoon",
		11: "Hot",
		12: "Changeable",
		13: "Snow",
		14: "Sleet",
		15: "Foggy",
		16: "Thunder",
		17: "Same as yesterday",
		18: "Moderate",
		19: "Humid",
	}
	var nh = map[int]int{
		2: 3,
		3: 3,
		4: 4,
		5: 4,
		6: 5,
		7: 5,
		8: 6,
		9: 6,
		10: 1,
		11: 1,
		12: 2,
	}
	hf := hex.NewFlower(content, nh, 10)
	for {
		fmt.Println(hf.State())
		fmt.Scanln()
		hf.MoveRandomly()
	}
}