// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SI-SchedulingInfo-v1740 (line 15071).

var sISchedulingInfoV1740Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "si-RequestConfigRedCap-r17", Optional: true},
	},
}

type SI_SchedulingInfo_v1740 struct {
	Si_RequestConfigRedCap_r17 *SI_RequestConfig
}

func (ie *SI_SchedulingInfo_v1740) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sISchedulingInfoV1740Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Si_RequestConfigRedCap_r17 != nil}); err != nil {
		return err
	}

	// 2. si-RequestConfigRedCap-r17: ref
	{
		if ie.Si_RequestConfigRedCap_r17 != nil {
			if err := ie.Si_RequestConfigRedCap_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SI_SchedulingInfo_v1740) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sISchedulingInfoV1740Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. si-RequestConfigRedCap-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Si_RequestConfigRedCap_r17 = new(SI_RequestConfig)
			if err := ie.Si_RequestConfigRedCap_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
