package ies

import (
	"testing"

	"github.com/lvdund/asn1go/per"
)

// TestQFI_List_r18_RoundTrip verifies the generated top-level SEQUENCE OF
// (SIZE(1..maxNrofQFIs) OF QFI) round-trips across both PER variants, with
// QFI being a wrapper-struct scalar (INTEGER 0..maxQFI).
func TestQFI_List_r18_RoundTrip(t *testing.T) {
	cases := []struct {
		name string
		ie   QFI_List_r18
	}{
		{name: "single element", ie: QFI_List_r18{{Value: 0}}},
		{name: "several elements", ie: QFI_List_r18{{Value: 1}, {Value: 5}, {Value: 63}}},
	}

	for _, variant := range []per.Variant{per.APER, per.UPER} {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				enc := per.NewEncoder(variant)
				if err := c.ie.Encode(enc); err != nil {
					t.Fatalf("Encode (%v): %v", variant, err)
				}
				data := enc.Bytes()

				var got QFI_List_r18
				dec := per.NewDecoder(data, variant)
				if err := got.Decode(dec); err != nil {
					t.Fatalf("Decode (%v): %v", variant, err)
				}
				if len(got) != len(c.ie) {
					t.Fatalf("len mismatch (%v): got %d want %d", variant, len(got), len(c.ie))
				}
				for i := range c.ie {
					if got[i].Value != c.ie[i].Value {
						t.Fatalf("elem %d mismatch (%v): got %d want %d",
							i, variant, got[i].Value, c.ie[i].Value)
					}
				}
			})
		}
	}
}
