package ies

import (
	"testing"

	"github.com/lvdund/asn1go/per"
)

// TestPDSCH_HARQ_ACK_CodebookList_r16_RoundTrip verifies a top-level
// SEQUENCE (SIZE(1..2)) OF ENUMERATED {semiStatic, dynamic} round-trips
// across both PER variants. This exercises the previously-missing
// "SEQUENCE OF with a non-scalar element" code path (enumerated element).
func TestPDSCH_HARQ_ACK_CodebookList_r16_RoundTrip(t *testing.T) {
	cases := []struct {
		name string
		ie   PDSCH_HARQ_ACK_CodebookList_r16
	}{
		{name: "single element semiStatic", ie: PDSCH_HARQ_ACK_CodebookList_r16{0}},
		{name: "two elements [semiStatic, dynamic]", ie: PDSCH_HARQ_ACK_CodebookList_r16{0, 1}},
		{name: "two elements [dynamic, dynamic]", ie: PDSCH_HARQ_ACK_CodebookList_r16{1, 1}},
	}

	for _, variant := range []per.Variant{per.APER, per.UPER} {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				enc := per.NewEncoder(variant)
				if err := c.ie.Encode(enc); err != nil {
					t.Fatalf("Encode (%v): %v", variant, err)
				}
				data := enc.Bytes()

				var got PDSCH_HARQ_ACK_CodebookList_r16
				dec := per.NewDecoder(data, variant)
				if err := got.Decode(dec); err != nil {
					t.Fatalf("Decode (%v): %v", variant, err)
				}
				if len(got) != len(c.ie) {
					t.Fatalf("len mismatch (%v): got %d want %d", variant, len(got), len(c.ie))
				}
				for i := range c.ie {
					if got[i] != c.ie[i] {
						t.Fatalf("elem %d mismatch (%v): got %d want %d",
							i, variant, got[i], c.ie[i])
					}
				}
			})
		}
	}
}
