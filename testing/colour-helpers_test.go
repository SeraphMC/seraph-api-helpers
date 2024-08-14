package testing

import (
	"github.com/Clemintina/common_utilities-for-apis/src/colours"
	"testing"
)

func TestHexToInt(t *testing.T) {
	colour := "F0E1E1"
	got, _ := colours.HexToInt(colour)
	want := uint32(0xF0E1E1)

	t.Logf("Hex: %s -> Decimal: %d\n", colour, got)
	t.Logf("Expected: %d\n", want)

	if got == want {
		t.Log("The conversion is correct.")
	} else {
		t.Log(got, want)
		t.Errorf("got %q, wanted %q", got, want)

	}
}
