package hex_test

import (
	"hex"
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var content = hex.HexFlowerContent{
	Start: 10,
	Hexes: map[int]string{
		1:  "Fine",
		2:  "Cloudy",
		3:  "Rainy",
		4:  "Stormy",
		5:  "Cold",
		6:  "Sunny",
		7:  "Hail",
		8:  "Stormy Seas",
		9:  "Misty",
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
	},
}

func newTestFlower() *hex.Flower {
	return hex.NewFlower(content)
}

func TestNewFlower(t *testing.T) {
	hf := newTestFlower()
	if hf.State() != content.Hexes[10] {
		t.Errorf("Expected %q, got %q\n", content.Hexes[10], hf.State())
	}
	// stay still
	hf.Move(-1)
	if hf.State() != content.Hexes[10] {
		t.Errorf("Expected %q, got %q\n", content.Hexes[10], hf.State())
	}
	// move to SE hex
	hf.Move(3)
	if hf.State() != content.Hexes[8] {
		t.Errorf("Expected %q, got %q\n", content.Hexes[8], hf.State())
	}
}

func TestMoveRandomly(t *testing.T) {
	hf := newTestFlower()
	hf.Random = rand.New(rand.NewSource(1))
	hf.MoveRandomly()
	wantState := content.Hexes[12]
	if wantState != hf.State() {
		t.Errorf("want state %q, got %q", wantState, hf.State())
	}
}

func TestRestrictedMove(t *testing.T) {
	hf := newTestFlower()
	hf.Move(0)
	hf.Move(0)
	hf.Move(5)
	wantHex := 9
	if wantHex != hf.CurrentHex() {
		t.Errorf("want current hex %d, got %d", wantHex, hf.CurrentHex())
	}
}

func TestBoundsChecking(t *testing.T) {
	hf := newTestFlower()
	hf.Move(0)
	hf.Move(0)
	wantHex := 14
	if wantHex != hf.CurrentHex() {
		t.Fatalf("want current hex %d, got %d", wantHex, hf.CurrentHex())
	}
	wantNbs := []int{17, 12, 9}
	if !cmp.Equal(wantNbs, hf.Neighbors()) {
		t.Error(cmp.Diff(wantNbs, hf.Neighbors()))
	}
	hf.Move(3)
	wantHex = 17
	if wantHex != hf.CurrentHex() {
		t.Errorf("want current hex %d, got %d", wantHex, hf.CurrentHex())
	}
}

func TestRoll2d6(t *testing.T) {
	hf := newTestFlower()
	hf.Random = rand.New(rand.NewSource(1))
	want := 10
	got := hf.Roll2d6()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestNeighbors(t *testing.T) {
	hf := newTestFlower()
	want := []int{12, 15, 13, 8, 5, 7}
	got := hf.Neighbors()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestLoadHexFlower(t *testing.T) {
	filename := "./testdata/hextest.yaml"
	content, err := hex.LoadContent(filename)
	if err != nil {
		t.Fatalf("Can't load file: %v", err)
	}
	hf := hex.NewFlower(content)
	hf.Move(-1)
	if hf.State() != "special" {
		t.Errorf("Expected %q, got %q\n", content.Hexes[10], hf.State())
	}
}

func TestNewFlowerFromFile(t *testing.T) {
	flower, err := hex.NewFlowerFromFile("testdata/hextest.yaml")
	if err != nil {
		t.Fatalf("Can't load file: %v", err)
	}
	if flower.CurrentHex() != 10 {
		t.Errorf("Expected %d, got %d\n", flower.CurrentHex(), 10)
	}
	if flower.State() != "special" {
		t.Errorf("Expected %q, got %q\n", content.Hexes[10], flower.State())
	}
}

func TestListContents(t *testing.T){
	got, err := hex.GetContentsList("testdata/content")
	want := "terrain\nweather\n"
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got){
		t.Error(cmp.Diff(want, got))
	}
}