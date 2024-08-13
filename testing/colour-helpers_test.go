package testing

import (
	"github.com/Clemintina/common_utilities-for-apis/src/colours"
	"testing"
)

func TestHexToInt(t *testing.T) {
	got, _ := colours.HexToInt("#f71b1bff")
	want := uint32(0xf71b1bff)

	t.Logf("Hex: %s -> Decimal: %d\n", "#437284", got)
	t.Logf("Expected: %d\n", want)

	if got == want {
		t.Log("The conversion is correct.")
	} else {
		t.Log(got, want)
		t.Errorf("got %q, wanted %q", got, want)

	}
}
