package ies

import (
	"testing"

	"github.com/lvdund/asn1go/per"
)

// TestBetaOffsetsCrossPriSel_r17_RoundTrip verifies a top-level CHOICE whose
// alternative is itself a SEQUENCE OF. CHOICE alternative fields are pointers.
func TestBetaOffsetsCrossPriSel_r17_RoundTrip(t *testing.T) {
	dyn := []BetaOffsetsCrossPri_r17{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12},
	}
	stat := BetaOffsetsCrossPri_r17{13, 14, 15}
	cases := []struct {
		name string
		ie   BetaOffsetsCrossPriSel_r17
	}{
		{
			name: "dynamic-r17 alternative (4 x BetaOffsetsCrossPri-r17)",
			ie:   BetaOffsetsCrossPriSel_r17{Choice: BetaOffsetsCrossPriSel_r17_Dynamic_r17, Dynamic_r17: &dyn},
		},
		{
			name: "semiStatic-r17 alternative",
			ie:   BetaOffsetsCrossPriSel_r17{Choice: BetaOffsetsCrossPriSel_r17_SemiStatic_r17, SemiStatic_r17: &stat},
		},
	}

	for _, variant := range []per.Variant{per.APER, per.UPER} {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				enc := per.NewEncoder(variant)
				if err := c.ie.Encode(enc); err != nil {
					t.Fatalf("Encode (%v): %v", variant, err)
				}
				data := enc.Bytes()

				var got BetaOffsetsCrossPriSel_r17
				dec := per.NewDecoder(data, variant)
				if err := got.Decode(dec); err != nil {
					t.Fatalf("Decode (%v): %v", variant, err)
				}
				if got.Choice != c.ie.Choice {
					t.Fatalf("Choice mismatch (%v): got %d want %d",
						variant, got.Choice, c.ie.Choice)
				}
			})
		}
	}
}
