package hex

import (
	"math/rand"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type HexFlowerContent struct {
	Start int            `yaml:"start"`
	Hexes map[int]string `yaml:hexes"`
}

var neighbors = map[int][]int{
	1:  {2, 5, 3, 0, 0, 0},
	2:  {4, 7, 5, 1, 0, 0},
	3:  {5, 8, 6, 0, 0, 1},
	4:  {0, 9, 7, 2, 0, 0},
	5:  {7, 10, 8, 3, 1, 2},
	6:  {8, 11, 0, 0, 0, 3},
	7:  {9, 12, 10, 5, 2, 4},
	8:  {10, 13, 11, 6, 3, 5},
	9:  {0, 14, 12, 7, 4, 0},
	10: {12, 15, 13, 8, 5, 7},
	11: {13, 16, 0, 0, 6, 8},
	12: {14, 17, 15, 10, 7, 9},
	13: {15, 18, 16, 11, 8, 10},
	14: {0, 0, 17, 12, 9, 0},
	15: {17, 19, 18, 13, 10, 12},
	16: {18, 0, 0, 0, 11, 13},
	17: {0, 0, 19, 15, 12, 14},
	18: {19, 0, 0, 16, 13, 15},
	19: {0, 0, 0, 18, 15, 17},
}

type NavHex map[int]int
var DefaultNavHex = NavHex{
	2:  2,
	3:  2,
	4:  3,
	5:  3,
	6:  4,
	7:  4,
	8:  5,
	9:  5,
	10: 0,
	11: 0,
	12: 1,
}

type Flower struct {
	currentNode int
	content     map[int]string
	NavHex      map[int]int
	Random      *rand.Rand
}

func NewFlower(content HexFlowerContent) *Flower {
	return &Flower{
		currentNode: content.Start,
		content:     content.Hexes,
		NavHex:      DefaultNavHex,
		Random:      rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func LoadContent(filename string) (HexFlowerContent, error) {
	var hfc HexFlowerContent
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return HexFlowerContent{}, err
	}
	err = yaml.Unmarshal(yamlFile, &hfc)
	if err != nil {
		return HexFlowerContent{}, err
	}

	return hfc, nil

}

func NewFlowerFromFile(filename string) (*Flower, error) {
	hfc, err := LoadContent(filename)
	if err != nil {
		return &Flower{}, nil
	}
	return NewFlower(hfc), nil
}

// Move goes in a hex in a direction from 0(NW) to 5(SW).
// -1 = stand still
func (f *Flower) Move(direction int) {
	choices := f.Neighbors()
	//if we stand still, our choice is to stay put.
	if direction == -1 {
		choices = []int{f.currentNode}
	}
	validDirection := direction % (len(choices))
	f.currentNode = choices[validDirection]
}

func (f *Flower) CurrentHex() int {
	return f.currentNode
}

func (f *Flower) SetHex(hex int) {
	f.currentNode = hex
}

func (f *Flower) State() string {
	return f.content[f.currentNode]
}

func (f *Flower) MoveRandomly() {
	roll := f.Roll2d6()
	direction := f.NavHex[roll]
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
