package hex_test

import "testing"
import "hex"



func TestNew(t *testing.T) {
	var hf *hex.Flower
	var content = map[int]string{
		8: "Stormy Seas",
		10: "Calm Day",
	}
	hf = hex.NewFlower(content)

	if hf.State() != content[10] {
		t.Errorf("Expected %q, got %q\n",content[10], hf.State())
	}
	hf.GoToNext()
	if hf.State() != content[8] {
		t.Errorf("Expected %q, got %q\n", content[8], hf.State())
	}
}

