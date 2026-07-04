package ies

import (
	"testing"

	"github.com/lvdund/asn1go/per"
)

// TestCarrierState_r17_RoundTrip verifies the generated top-level CHOICE
// encodes/decodes back to an equal value across both PER variants, exercising
// a NULL alternative and a constrained-INTEGER alternative (whose bound
// resolves to common.MaxNrofBWPs). CHOICE alternative fields are pointers.
func TestCarrierState_r17_RoundTrip(t *testing.T) {
	ip := func(v int64) *int64 { return &v }
	cases := []struct {
		name string
		ie   CarrierState_r17
	}{
		{
			name: "null alternative (deActivated)",
			ie:   CarrierState_r17{Choice: CarrierState_r17_DeActivated_r17, DeActivated_r17: &struct{}{}},
		},
		{
			name: "integer alternative (activeBWP=3)",
			ie:   CarrierState_r17{Choice: CarrierState_r17_ActiveBWP_r17, ActiveBWP_r17: ip(3)},
		},
		{
			name: "integer alternative at max bound",
			ie:   CarrierState_r17{Choice: CarrierState_r17_ActiveBWP_r17, ActiveBWP_r17: ip(4)},
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

				var got CarrierState_r17
				dec := per.NewDecoder(data, variant)
				if err := got.Decode(dec); err != nil {
					t.Fatalf("Decode (%v): %v", variant, err)
				}
				if got.Choice != c.ie.Choice {
					t.Fatalf("Choice mismatch (%v): got %d want %d",
						variant, got.Choice, c.ie.Choice)
				}
				switch c.ie.Choice {
				case CarrierState_r17_ActiveBWP_r17:
					if got.ActiveBWP_r17 == nil || *got.ActiveBWP_r17 != *c.ie.ActiveBWP_r17 {
						t.Fatalf("ActiveBWP mismatch (%v): got %v want %d",
							variant, got.ActiveBWP_r17, *c.ie.ActiveBWP_r17)
					}
				}
			})
		}
	}
}
