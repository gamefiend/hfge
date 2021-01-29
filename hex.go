package hex

import (
	"math/rand"
	"time"
)

var neighbors = map[int][]int{
	1: {2,5,3,0,0,0},
	2: {4,7,5,1,0,0},
	3: {5,8,6,0,0,1},
	4: {0,9,7,2,0,0},
	5: {7,10,8,3,1,2},
	6: {8,11,0,0,0,3},
	7: {9,12,10,5,2,4},
	8: {10,13,11,6,3,5},
	9: {0,14,12,7,4,0},
	10: {12,15,13,8,5,7},
	11: {13,16,0,0,6,8},
	12: {14,17,15,10,7,9},
	13: {15,18,16,11,8,10},
	14: {0,0,17,12,9,0},
	15: {17,19,18,13,10,12},
	16: {18,0,0,0,11,13},
	17: {0,0,19,15,12,14},
	18: {19,0,0,16,13,15},
	19: {0,0,0,18,15,17},
}

type Flower struct {
	currentNode int
	content map[int]string
	navHex map[int]int
	Random *rand.Rand
}

func NewFlower(content map[int]string, nh map[int]int, start int) *Flower{
	return &Flower{
		currentNode: start,
		content: content,
		navHex: nh,
		Random: rand.New(rand.NewSource(time.Now().Unix())),
	}
}

// Move goes in a hex in a direction from 1(NW) to 6(SW).
// 0 = stand still
func (f *Flower) Move(direction int) {
	if direction != 0 {
		choices := f.Neighbors()
		validDirection := direction%(len(choices)-1)
		f.currentNode = choices[validDirection]
	}
}

func (f *Flower) CurrentHex() int {
	return f.currentNode
}

func (f *Flower) State() string {
	return f.content[f.currentNode]
}

func (f *Flower) MoveRandomly() {
	roll := f.Roll2d6()
	direction := f.navHex[roll]
	f.Move(direction)
}

func (f Flower) Roll2d6() int {
	d1 := f.Random.Intn(7) + 1
	d2 := f.Random.Intn(7) + 1
	return d1 + d2
}

func (f Flower) Neighbors() []int {
	rawNbs := neighbors[f.currentNode]
	result := make([]int, 0, len(rawNbs))
	for _, n := range rawNbs {
		if n != 0 {
			result = append(result, n)
		}
	}
	return result
}