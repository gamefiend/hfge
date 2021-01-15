package hex_test

import "testing"
import "hex"



func TestNew(t *testing.T) {
	var hf *hex.Flower
	var content = map[int]string{
		6: "Stormy Seas",
		9: "Calm Day",
	}
	hf = hex.NewFlower(content)

	if hf.State() != content[9] {
		t.Errorf("Expected %q, got %q\n",content[9], hf.State())
	}
	hf.GoToNext()
	if hf.State() != content[6] {
		t.Errorf("Expected %q, got %q\n", content[6], hf.State())
	}
}

