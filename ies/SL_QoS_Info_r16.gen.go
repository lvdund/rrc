// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-QoS-Info-r16 (line 2277).

var sLQoSInfoR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-QoS-FlowIdentity-r16"},
		{Name: "sl-QoS-Profile-r16", Optional: true},
	},
}

type SL_QoS_Info_r16 struct {
	Sl_QoS_FlowIdentity_r16 SL_QoS_FlowIdentity_r16
	Sl_QoS_Profile_r16      *SL_QoS_Profile_r16
}

func (ie *SL_QoS_Info_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLQoSInfoR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_QoS_Profile_r16 != nil}); err != nil {
		return err
	}

	// 2. sl-QoS-FlowIdentity-r16: ref
	{
		if err := ie.Sl_QoS_FlowIdentity_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. sl-QoS-Profile-r16: ref
	{
		if ie.Sl_QoS_Profile_r16 != nil {
			if err := ie.Sl_QoS_Profile_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_QoS_Info_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLQoSInfoR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-QoS-FlowIdentity-r16: ref
	{
		if err := ie.Sl_QoS_FlowIdentity_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. sl-QoS-Profile-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_QoS_Profile_r16 = new(SL_QoS_Profile_r16)
			if err := ie.Sl_QoS_Profile_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
