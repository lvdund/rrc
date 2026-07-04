package ies

import (
	"testing"

	"github.com/lvdund/asn1go/per"
)

// TestDL_CCCH_MessageType_RoundTrip verifies a top-level CHOICE with a NESTED
// inline CHOICE alternative (c1). Both outer and inner CHOICE structs have
// pointer alt fields. Uses the NULL alts (spare1/spare2/messageClassExtension)
// which don't require deep struct population.
func TestDL_CCCH_MessageType_RoundTrip(t *testing.T) {
	cases := []struct {
		name string
		ie   DL_CCCH_MessageType
	}{
		{
			name: "c1.spare1 (nested CHOICE NULL alt)",
			ie: DL_CCCH_MessageType{
				Choice: DL_CCCH_MessageType_C1,
				C1: &struct {
					Choice    int
					RrcReject *RRCReject
					RrcSetup  *RRCSetup
					Spare2    *struct{}
					Spare1    *struct{}
				}{
					Choice: DL_CCCH_MessageType_C1_Spare1,
					Spare1: &struct{}{},
				},
			},
		},
		{
			name: "c1.spare2 (nested CHOICE NULL alt)",
			ie: DL_CCCH_MessageType{
				Choice: DL_CCCH_MessageType_C1,
				C1: &struct {
					Choice    int
					RrcReject *RRCReject
					RrcSetup  *RRCSetup
					Spare2    *struct{}
					Spare1    *struct{}
				}{
					Choice: DL_CCCH_MessageType_C1_Spare2,
					Spare2: &struct{}{},
				},
			},
		},
		{
			name: "messageClassExtension (outer NULL alt)",
			ie: DL_CCCH_MessageType{
				Choice:                DL_CCCH_MessageType_MessageClassExtension,
				MessageClassExtension: &struct{}{},
			},
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

				var got DL_CCCH_MessageType
				dec := per.NewDecoder(data, variant)
				if err := got.Decode(dec); err != nil {
					t.Fatalf("Decode (%v): %v", variant, err)
				}
				if got.Choice != c.ie.Choice {
					t.Fatalf("outer Choice mismatch (%v): got %d want %d",
						variant, got.Choice, c.ie.Choice)
				}
				if c.ie.C1 != nil && got.C1 != nil {
					if got.C1.Choice != c.ie.C1.Choice {
						t.Fatalf("inner Choice mismatch (%v): got %d want %d",
							variant, got.C1.Choice, c.ie.C1.Choice)
					}
				}
			})
		}
	}
}
