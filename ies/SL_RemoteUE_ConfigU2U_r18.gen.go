// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-RemoteUE-ConfigU2U-r18 (line 27856).

var sLRemoteUEConfigU2UR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RSRP-ThreshU2U-r18", Optional: true},
		{Name: "sl-HystMinU2U-r18", Optional: true},
		{Name: "sd-RSRP-ThreshU2U-r18", Optional: true},
		{Name: "sd-FilterCoefficientU2U-r18", Optional: true},
		{Name: "sd-HystMinU2U-r18", Optional: true},
	},
}

type SL_RemoteUE_ConfigU2U_r18 struct {
	Sl_RSRP_ThreshU2U_r18       *SL_RSRP_Range_r16
	Sl_HystMinU2U_r18           *Hysteresis
	Sd_RSRP_ThreshU2U_r18       *SL_RSRP_Range_r16
	Sd_FilterCoefficientU2U_r18 *FilterCoefficient
	Sd_HystMinU2U_r18           *Hysteresis
}

func (ie *SL_RemoteUE_ConfigU2U_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLRemoteUEConfigU2UR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_RSRP_ThreshU2U_r18 != nil, ie.Sl_HystMinU2U_r18 != nil, ie.Sd_RSRP_ThreshU2U_r18 != nil, ie.Sd_FilterCoefficientU2U_r18 != nil, ie.Sd_HystMinU2U_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-RSRP-ThreshU2U-r18: ref
	{
		if ie.Sl_RSRP_ThreshU2U_r18 != nil {
			if err := ie.Sl_RSRP_ThreshU2U_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. sl-HystMinU2U-r18: ref
	{
		if ie.Sl_HystMinU2U_r18 != nil {
			if err := ie.Sl_HystMinU2U_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sd-RSRP-ThreshU2U-r18: ref
	{
		if ie.Sd_RSRP_ThreshU2U_r18 != nil {
			if err := ie.Sd_RSRP_ThreshU2U_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sd-FilterCoefficientU2U-r18: ref
	{
		if ie.Sd_FilterCoefficientU2U_r18 != nil {
			if err := ie.Sd_FilterCoefficientU2U_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. sd-HystMinU2U-r18: ref
	{
		if ie.Sd_HystMinU2U_r18 != nil {
			if err := ie.Sd_HystMinU2U_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_RemoteUE_ConfigU2U_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLRemoteUEConfigU2UR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-RSRP-ThreshU2U-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_RSRP_ThreshU2U_r18 = new(SL_RSRP_Range_r16)
			if err := ie.Sl_RSRP_ThreshU2U_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. sl-HystMinU2U-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_HystMinU2U_r18 = new(Hysteresis)
			if err := ie.Sl_HystMinU2U_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sd-RSRP-ThreshU2U-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sd_RSRP_ThreshU2U_r18 = new(SL_RSRP_Range_r16)
			if err := ie.Sd_RSRP_ThreshU2U_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sd-FilterCoefficientU2U-r18: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Sd_FilterCoefficientU2U_r18 = new(FilterCoefficient)
			if err := ie.Sd_FilterCoefficientU2U_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. sd-HystMinU2U-r18: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Sd_HystMinU2U_r18 = new(Hysteresis)
			if err := ie.Sd_HystMinU2U_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
