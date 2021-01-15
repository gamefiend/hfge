package hex

import (
	"math/rand"
)

var neighbors = map[int][]int{
	9: []int{10, 16, 13, 8, 4, 6},
}


type Flower struct {
	currentNode int
	content map[int]string
}

func NewFlower(content map[int]string) *Flower{
	return &Flower{
		currentNode: 9,
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
