// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-RemoteUE-Config-r17 (line 27841).

var sLRemoteUEConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "threshHighRemote-r17", Optional: true},
		{Name: "hystMaxRemote-r17", Optional: true},
		{Name: "sl-ReselectionConfig-r17", Optional: true},
	},
}

type SL_RemoteUE_Config_r17 struct {
	ThreshHighRemote_r17     *RSRP_Range
	HystMaxRemote_r17        *Hysteresis
	Sl_ReselectionConfig_r17 *SL_ReselectionConfig_r17
}

func (ie *SL_RemoteUE_Config_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLRemoteUEConfigR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ThreshHighRemote_r17 != nil, ie.HystMaxRemote_r17 != nil, ie.Sl_ReselectionConfig_r17 != nil}); err != nil {
		return err
	}

	// 2. threshHighRemote-r17: ref
	{
		if ie.ThreshHighRemote_r17 != nil {
			if err := ie.ThreshHighRemote_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. hystMaxRemote-r17: ref
	{
		if ie.HystMaxRemote_r17 != nil {
			if err := ie.HystMaxRemote_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-ReselectionConfig-r17: ref
	{
		if ie.Sl_ReselectionConfig_r17 != nil {
			if err := ie.Sl_ReselectionConfig_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_RemoteUE_Config_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLRemoteUEConfigR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. threshHighRemote-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.ThreshHighRemote_r17 = new(RSRP_Range)
			if err := ie.ThreshHighRemote_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. hystMaxRemote-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.HystMaxRemote_r17 = new(Hysteresis)
			if err := ie.HystMaxRemote_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-ReselectionConfig-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_ReselectionConfig_r17 = new(SL_ReselectionConfig_r17)
			if err := ie.Sl_ReselectionConfig_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
