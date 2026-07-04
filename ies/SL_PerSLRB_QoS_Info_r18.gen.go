// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PerSLRB-QoS-Info-r18 (line 2318).

var sLPerSLRBQoSInfoR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RemoteUE-SLRB-Identity-r18"},
		{Name: "sl-QoS-ProfilePerSLRB-r18", Optional: true},
	},
}

type SL_PerSLRB_QoS_Info_r18 struct {
	Sl_RemoteUE_SLRB_Identity_r18 SLRB_Uu_ConfigIndex_r16
	Sl_QoS_ProfilePerSLRB_r18     *SL_QoS_Profile_r16
}

func (ie *SL_PerSLRB_QoS_Info_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPerSLRBQoSInfoR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_QoS_ProfilePerSLRB_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-RemoteUE-SLRB-Identity-r18: ref
	{
		if err := ie.Sl_RemoteUE_SLRB_Identity_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. sl-QoS-ProfilePerSLRB-r18: ref
	{
		if ie.Sl_QoS_ProfilePerSLRB_r18 != nil {
			if err := ie.Sl_QoS_ProfilePerSLRB_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_PerSLRB_QoS_Info_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPerSLRBQoSInfoR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-RemoteUE-SLRB-Identity-r18: ref
	{
		if err := ie.Sl_RemoteUE_SLRB_Identity_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. sl-QoS-ProfilePerSLRB-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_QoS_ProfilePerSLRB_r18 = new(SL_QoS_Profile_r16)
			if err := ie.Sl_QoS_ProfilePerSLRB_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
