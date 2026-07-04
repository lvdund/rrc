// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: InterFreqCarrierFreqInfo-v1730 (line 4026).

var interFreqCarrierFreqInfoV1730Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "channelAccessMode2-r17", Optional: true},
	},
}

const (
	InterFreqCarrierFreqInfo_v1730_ChannelAccessMode2_r17_Enabled = 0
)

var interFreqCarrierFreqInfoV1730ChannelAccessMode2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{InterFreqCarrierFreqInfo_v1730_ChannelAccessMode2_r17_Enabled},
}

type InterFreqCarrierFreqInfo_v1730 struct {
	ChannelAccessMode2_r17 *int64
}

func (ie *InterFreqCarrierFreqInfo_v1730) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(interFreqCarrierFreqInfoV1730Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ChannelAccessMode2_r17 != nil}); err != nil {
		return err
	}

	// 2. channelAccessMode2-r17: enumerated
	{
		if ie.ChannelAccessMode2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ChannelAccessMode2_r17, interFreqCarrierFreqInfoV1730ChannelAccessMode2R17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *InterFreqCarrierFreqInfo_v1730) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(interFreqCarrierFreqInfoV1730Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. channelAccessMode2-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(interFreqCarrierFreqInfoV1730ChannelAccessMode2R17Constraints)
			if err != nil {
				return err
			}
			ie.ChannelAccessMode2_r17 = &idx
		}
	}

	return nil
}
