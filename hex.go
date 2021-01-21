package hex

import (
	"math/rand"
)

var neighbors = map[int][]int{
	1: []int{2,5,3,0,0,0},
	2: []int{4,7,5,1,0,0},
	3: []int{5,8,6,0,0,1},
	4: []int{0,9,7,2,0,0},
	5: []int{7,10,8,3,1,2},
	6: []int{8,11,0,0,0,3},
	7: []int{9,12,10,5,2,4},
	8: []int{10,13,11,6,3,5},
	9: []int{0,14,12,7,4,0},
	10: []int{12,15,13,8,5,7},
	11: []int{13,16,0,0,6,8},
	12: []int{14,17,15,10,7,9},
	13: []int{15,18,16,11,8,10},
	14: []int{0,0,17,12,9,0},
	15: []int{17,19,18,13,10,12},
	16: []int{18,0,0,0,11,13},
	17: []int{0,0,19,15,12,14},
	18: []int{19,0,0,16,13,15},
	19: []int{0,0,0,18,15,17},
}


type Flower struct {
	currentNode int
	content map[int]string
}

func NewFlower(content map[int]string) *Flower{
	return &Flower{
		currentNode: 10,
		content: content,
	}
}

func (f *Flower) GoToNext() {

	choices := neighbors[f.currentNode]

	f.currentNode = choices[rand.Intn(len(choices))]
}

func (f *Flower) State () string {
	return f.content[f.currentNode]
}
