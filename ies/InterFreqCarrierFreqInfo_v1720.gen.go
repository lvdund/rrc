// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: InterFreqCarrierFreqInfo-v1720 (line 4022).

var interFreqCarrierFreqInfoV1720Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "smtc4list-r17", Optional: true},
	},
}

type InterFreqCarrierFreqInfo_v1720 struct {
	Smtc4list_r17 *SSB_MTC4List_r17
}

func (ie *InterFreqCarrierFreqInfo_v1720) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(interFreqCarrierFreqInfoV1720Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Smtc4list_r17 != nil}); err != nil {
		return err
	}

	// 2. smtc4list-r17: ref
	{
		if ie.Smtc4list_r17 != nil {
			if err := ie.Smtc4list_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *InterFreqCarrierFreqInfo_v1720) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(interFreqCarrierFreqInfoV1720Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. smtc4list-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Smtc4list_r17 = new(SSB_MTC4List_r17)
			if err := ie.Smtc4list_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
