// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-RelayUE-ConfigU2U-r18 (line 27826).

var sLRelayUEConfigU2UR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RSRP-Thresh-DiscConfig-r18", Optional: true},
		{Name: "sd-RSRP-ThreshDiscConfig-r18", Optional: true},
		{Name: "sd-hystMaxRelay-r18", Optional: true},
	},
}

type SL_RelayUE_ConfigU2U_r18 struct {
	Sl_RSRP_Thresh_DiscConfig_r18 *SL_RSRP_Range_r16
	Sd_RSRP_ThreshDiscConfig_r18  *SL_RSRP_Range_r16
	Sd_HystMaxRelay_r18           *Hysteresis
}

func (ie *SL_RelayUE_ConfigU2U_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLRelayUEConfigU2UR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_RSRP_Thresh_DiscConfig_r18 != nil, ie.Sd_RSRP_ThreshDiscConfig_r18 != nil, ie.Sd_HystMaxRelay_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-RSRP-Thresh-DiscConfig-r18: ref
	{
		if ie.Sl_RSRP_Thresh_DiscConfig_r18 != nil {
			if err := ie.Sl_RSRP_Thresh_DiscConfig_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. sd-RSRP-ThreshDiscConfig-r18: ref
	{
		if ie.Sd_RSRP_ThreshDiscConfig_r18 != nil {
			if err := ie.Sd_RSRP_ThreshDiscConfig_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sd-hystMaxRelay-r18: ref
	{
		if ie.Sd_HystMaxRelay_r18 != nil {
			if err := ie.Sd_HystMaxRelay_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_RelayUE_ConfigU2U_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLRelayUEConfigU2UR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-RSRP-Thresh-DiscConfig-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_RSRP_Thresh_DiscConfig_r18 = new(SL_RSRP_Range_r16)
			if err := ie.Sl_RSRP_Thresh_DiscConfig_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. sd-RSRP-ThreshDiscConfig-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sd_RSRP_ThreshDiscConfig_r18 = new(SL_RSRP_Range_r16)
			if err := ie.Sd_RSRP_ThreshDiscConfig_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sd-hystMaxRelay-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sd_HystMaxRelay_r18 = new(Hysteresis)
			if err := ie.Sd_HystMaxRelay_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
