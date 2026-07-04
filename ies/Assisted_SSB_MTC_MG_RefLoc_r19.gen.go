// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Assisted-SSB-MTC-MG-RefLoc-r19 (line 4537).

var assistedSSBMTCMGRefLocR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "refLoc-r19", Optional: true},
	},
}

type Assisted_SSB_MTC_MG_RefLoc_r19 struct {
	RefLoc_r19 *ReferenceLocation_r17
}

func (ie *Assisted_SSB_MTC_MG_RefLoc_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(assistedSSBMTCMGRefLocR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RefLoc_r19 != nil}); err != nil {
		return err
	}

	// 2. refLoc-r19: ref
	{
		if ie.RefLoc_r19 != nil {
			if err := ie.RefLoc_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *Assisted_SSB_MTC_MG_RefLoc_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(assistedSSBMTCMGRefLocR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. refLoc-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.RefLoc_r19 = new(ReferenceLocation_r17)
			if err := ie.RefLoc_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
