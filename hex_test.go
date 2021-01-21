package hex_test

import "testing"
import "hex"



func TestNewFlower(t *testing.T) {
	var hf *hex.Flower

	var content = map[int]string{
		8: "Stormy Seas",
		10: "Calm Day",
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
	hf = hex.NewFlower(content, nh, 10)
	if hf.State() != content[10] {
		t.Errorf("Expected %q, got %q\n",content[10], hf.State())
	}
	// stay still
	hf.GoToNextDirection(0)
	if hf.State() != content[10] {
		t.Errorf("Expected %q, got %q\n",content[10], hf.State())
	}
	// move to SE hex
	hf.GoToNextDirection(4)
	if hf.State() != content[8] {
		t.Errorf("Expected %q, got %q\n", content[8], hf.State())
	}
}



