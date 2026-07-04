// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-RelayUE-Config-r17 (line 27808).

var sLRelayUEConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "threshHighRelay-r17", Optional: true},
		{Name: "threshLowRelay-r17", Optional: true},
		{Name: "hystMaxRelay-r17", Optional: true},
		{Name: "hystMinRelay-r17", Optional: true},
	},
}

type SL_RelayUE_Config_r17 struct {
	ThreshHighRelay_r17 *RSRP_Range
	ThreshLowRelay_r17  *RSRP_Range
	HystMaxRelay_r17    *Hysteresis
	HystMinRelay_r17    *Hysteresis
}

func (ie *SL_RelayUE_Config_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLRelayUEConfigR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ThreshHighRelay_r17 != nil, ie.ThreshLowRelay_r17 != nil, ie.HystMaxRelay_r17 != nil, ie.HystMinRelay_r17 != nil}); err != nil {
		return err
	}

	// 2. threshHighRelay-r17: ref
	{
		if ie.ThreshHighRelay_r17 != nil {
			if err := ie.ThreshHighRelay_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. threshLowRelay-r17: ref
	{
		if ie.ThreshLowRelay_r17 != nil {
			if err := ie.ThreshLowRelay_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. hystMaxRelay-r17: ref
	{
		if ie.HystMaxRelay_r17 != nil {
			if err := ie.HystMaxRelay_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. hystMinRelay-r17: ref
	{
		if ie.HystMinRelay_r17 != nil {
			if err := ie.HystMinRelay_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_RelayUE_Config_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLRelayUEConfigR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. threshHighRelay-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.ThreshHighRelay_r17 = new(RSRP_Range)
			if err := ie.ThreshHighRelay_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. threshLowRelay-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.ThreshLowRelay_r17 = new(RSRP_Range)
			if err := ie.ThreshLowRelay_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. hystMaxRelay-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.HystMaxRelay_r17 = new(Hysteresis)
			if err := ie.HystMaxRelay_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. hystMinRelay-r17: ref
	{
		if seq.IsComponentPresent(3) {
			ie.HystMinRelay_r17 = new(Hysteresis)
			if err := ie.HystMinRelay_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
