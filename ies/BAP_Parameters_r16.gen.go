// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BAP-Parameters-r16 (line 25932).

var bAPParametersR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "flowControlBH-RLC-ChannelBased-r16", Optional: true},
		{Name: "flowControlRouting-ID-Based-r16", Optional: true},
	},
}

const (
	BAP_Parameters_r16_FlowControlBH_RLC_ChannelBased_r16_Supported = 0
)

var bAPParametersR16FlowControlBHRLCChannelBasedR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BAP_Parameters_r16_FlowControlBH_RLC_ChannelBased_r16_Supported},
}

const (
	BAP_Parameters_r16_FlowControlRouting_ID_Based_r16_Supported = 0
)

var bAPParametersR16FlowControlRoutingIDBasedR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BAP_Parameters_r16_FlowControlRouting_ID_Based_r16_Supported},
}

type BAP_Parameters_r16 struct {
	FlowControlBH_RLC_ChannelBased_r16 *int64
	FlowControlRouting_ID_Based_r16    *int64
}

func (ie *BAP_Parameters_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bAPParametersR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FlowControlBH_RLC_ChannelBased_r16 != nil, ie.FlowControlRouting_ID_Based_r16 != nil}); err != nil {
		return err
	}

	// 2. flowControlBH-RLC-ChannelBased-r16: enumerated
	{
		if ie.FlowControlBH_RLC_ChannelBased_r16 != nil {
			if err := e.EncodeEnumerated(*ie.FlowControlBH_RLC_ChannelBased_r16, bAPParametersR16FlowControlBHRLCChannelBasedR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. flowControlRouting-ID-Based-r16: enumerated
	{
		if ie.FlowControlRouting_ID_Based_r16 != nil {
			if err := e.EncodeEnumerated(*ie.FlowControlRouting_ID_Based_r16, bAPParametersR16FlowControlRoutingIDBasedR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BAP_Parameters_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bAPParametersR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. flowControlBH-RLC-ChannelBased-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(bAPParametersR16FlowControlBHRLCChannelBasedR16Constraints)
			if err != nil {
				return err
			}
			ie.FlowControlBH_RLC_ChannelBased_r16 = &idx
		}
	}

	// 3. flowControlRouting-ID-Based-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(bAPParametersR16FlowControlRoutingIDBasedR16Constraints)
			if err != nil {
				return err
			}
			ie.FlowControlRouting_ID_Based_r16 = &idx
		}
	}

	return nil
}
