package hex_test

import (
	"hex"
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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

func newTestFlower() *hex.Flower {
	return hex.NewFlower(content,
		map[int]int{
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
		}, 10)
}

func TestNewFlower(t *testing.T) {
	hf := newTestFlower()
	if hf.State() != content[10] {
		t.Errorf("Expected %q, got %q\n",content[10], hf.State())
	}
	// stay still
	hf.Move(0)
	if hf.State() != content[10] {
		t.Errorf("Expected %q, got %q\n",content[10], hf.State())
	}
	// move to SE hex
	hf.Move(4)
	if hf.State() != content[8] {
		t.Errorf("Expected %q, got %q\n", content[8], hf.State())
	}
}

func TestMoveRandomly(t *testing.T) {
	hf := newTestFlower()
	hf.Random = rand.New(rand.NewSource(1))
	hf.MoveRandomly()
	wantState := content[12]
	if wantState != hf.State() {
		t.Errorf("want state %q, got %q", wantState, hf.State())
	}
}

func TestRestrictedMove(t *testing.T) {
	hf := newTestFlower()
	hf.Move(1)
	hf.Move(1)
	hf.Move(6)
	wantHex := 9
	if wantHex != hf.CurrentHex() {
		t.Errorf("want current hex %d, got %d", wantHex, hf.CurrentHex())
	}
}

func TestBoundsChecking(t *testing.T) {
	hf := newTestFlower()
	hf.Move(1)
	hf.Move(1)
	wantHex := 14
	if wantHex != hf.CurrentHex() {
		t.Fatalf("want current hex %d, got %d", wantHex, hf.CurrentHex())
	}
	wantNbs := []int{17,12,9}
	if !cmp.Equal(wantNbs, hf.Neighbors()) {
		t.Error(cmp.Diff(wantNbs, hf.Neighbors()))
	}
	hf.Move(4)
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
	want := []int{12,15,13,8,5,7}
	got := hf.Neighbors()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}