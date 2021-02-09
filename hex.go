package hex

import (
	"github.com/oleiade/reflections"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type HexFlowerContent struct {
	Start     int    `yaml:"start"`
	Hexes     struct {
		Num1  string `yaml:"1"`
		Num2  string `yaml:"2"`
		Num3  string `yaml:"3"`
		Num4  string `yaml:"4"`
		Num5  string `yaml:"5"`
		Num6  string `yaml:"6"`
		Num7  string `yaml:"7"`
		Num8  string `yaml:"8"`
		Num9  string `yaml:"9"`
		Num10 string `yaml:"10"`
		Num11 string `yaml:"11"`
		Num12 string `yaml:"12"`
		Num13 string `yaml:"13"`
		Num14 string `yaml:"14"`
		Num15 string `yaml:"15"`
		Num16 string `yaml:"16"`
		Num17 string `yaml:"17"`
		Num18 string `yaml:"18"`
		Num19 string `yaml:"19"`
	} `yaml:"hexes"`
}
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

var NavHex = map[int]int{
	2: 2,
	3: 2,
	4: 3,
	5: 3,
	6: 4,
	7: 4,
	8: 5,
	9: 5,
	10: 0,
	11: 0,
	12: 1,
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

func LoadContent(filename string) (map[int]string, int){
	var hfc HexFlowerContent
	yamlFile, err := ioutil.ReadFile(filename)
	if err!=nil {
		log.Fatalf("can't open file: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &hfc)
	if err!= nil {
		log.Fatalf("can't parse yaml file: %v", err)
	}

	// remove the Num that we used to import and convert to a map[int]string
	content := make(map[int]string, 19)
	for i := 1; i < 20; i++ {
		numStr := "Num" + strconv.Itoa(i)
		value, err := reflections.GetField(hfc.Hexes, numStr)
		if err != nil {
			log.Fatalf("cannot translate yaml to content: %v",err)
		}
		s := value.(string)
		content[i] = s
	}
	return content, hfc.Start

}
// Move goes in a hex in a direction from 0(NW) to 5(SW).
// 0 = stand still
func (f *Flower) Move(direction int) {
	if direction != -1 {
		choices := f.Neighbors()
		validDirection := direction%(len(choices))
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

