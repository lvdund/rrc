// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-ReselectionConfig-r17 (line 27847).

var sLReselectionConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RSRP-Thresh-r17", Optional: true},
		{Name: "sl-FilterCoefficientRSRP-r17", Optional: true},
		{Name: "sl-HystMin-r17", Optional: true},
	},
}

type SL_ReselectionConfig_r17 struct {
	Sl_RSRP_Thresh_r17           *SL_RSRP_Range_r16
	Sl_FilterCoefficientRSRP_r17 *FilterCoefficient
	Sl_HystMin_r17               *Hysteresis
}

func (ie *SL_ReselectionConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLReselectionConfigR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_RSRP_Thresh_r17 != nil, ie.Sl_FilterCoefficientRSRP_r17 != nil, ie.Sl_HystMin_r17 != nil}); err != nil {
		return err
	}

	// 2. sl-RSRP-Thresh-r17: ref
	{
		if ie.Sl_RSRP_Thresh_r17 != nil {
			if err := ie.Sl_RSRP_Thresh_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. sl-FilterCoefficientRSRP-r17: ref
	{
		if ie.Sl_FilterCoefficientRSRP_r17 != nil {
			if err := ie.Sl_FilterCoefficientRSRP_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-HystMin-r17: ref
	{
		if ie.Sl_HystMin_r17 != nil {
			if err := ie.Sl_HystMin_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_ReselectionConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLReselectionConfigR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-RSRP-Thresh-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_RSRP_Thresh_r17 = new(SL_RSRP_Range_r16)
			if err := ie.Sl_RSRP_Thresh_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. sl-FilterCoefficientRSRP-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_FilterCoefficientRSRP_r17 = new(FilterCoefficient)
			if err := ie.Sl_FilterCoefficientRSRP_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-HystMin-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_HystMin_r17 = new(Hysteresis)
			if err := ie.Sl_HystMin_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
