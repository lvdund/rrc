// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: OtherConfig-v1830 (line 26357).

var otherConfigV1830Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PRS-AssistanceConfigNR-r18", Optional: true},
	},
}

const (
	OtherConfig_v1830_Sl_PRS_AssistanceConfigNR_r18_True = 0
)

var otherConfigV1830SlPRSAssistanceConfigNRR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OtherConfig_v1830_Sl_PRS_AssistanceConfigNR_r18_True},
}

type OtherConfig_v1830 struct {
	Sl_PRS_AssistanceConfigNR_r18 *int64
}

func (ie *OtherConfig_v1830) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(otherConfigV1830Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PRS_AssistanceConfigNR_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-PRS-AssistanceConfigNR-r18: enumerated
	{
		if ie.Sl_PRS_AssistanceConfigNR_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_PRS_AssistanceConfigNR_r18, otherConfigV1830SlPRSAssistanceConfigNRR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *OtherConfig_v1830) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(otherConfigV1830Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-PRS-AssistanceConfigNR-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(otherConfigV1830SlPRSAssistanceConfigNRR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PRS_AssistanceConfigNR_r18 = &idx
		}
	}

	return nil
}
