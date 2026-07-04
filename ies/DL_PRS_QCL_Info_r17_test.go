package ies

import (
	"testing"

	"github.com/lvdund/asn1go/per"
)

// TestDL_PRS_QCL_Info_r17_RoundTrip verifies a top-level CHOICE whose
// alternatives are inline SEQUENCEs that are themselves extensible
// (previously skipped). CHOICE alternative fields are pointers.
func TestDL_PRS_QCL_Info_r17_RoundTrip(t *testing.T) {
	ssbType := struct {
		Ssb_Index_r17 int64
		Rs_Type_r17   int64
	}{Ssb_Index_r17: 5, Rs_Type_r17: 0}
	dlType := struct {
		Qcl_DL_PRS_ResourceID_r17 NR_DL_PRS_ResourceID_r17
	}{Qcl_DL_PRS_ResourceID_r17: NR_DL_PRS_ResourceID_r17{}}
	cases := []struct {
		name string
		ie   DL_PRS_QCL_Info_r17
	}{
		{
			name: "ssb-r17 alternative",
			ie:   DL_PRS_QCL_Info_r17{Choice: DL_PRS_QCL_Info_r17_Ssb_r17, Ssb_r17: &ssbType},
		},
		{
			name: "dl-PRS-r17 alternative",
			ie:   DL_PRS_QCL_Info_r17{Choice: DL_PRS_QCL_Info_r17_Dl_PRS_r17, Dl_PRS_r17: &dlType},
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

				var got DL_PRS_QCL_Info_r17
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
