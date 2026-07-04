// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MBS-Parameters-r17 (line 25942).

var mBSParametersR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxMRB-Add-r17", Optional: true},
	},
}

var mBSParametersR17MaxMRBAddR17Constraints = per.Constrained(1, 16)

type MBS_Parameters_r17 struct {
	MaxMRB_Add_r17 *int64
}

func (ie *MBS_Parameters_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSParametersR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaxMRB_Add_r17 != nil}); err != nil {
		return err
	}

	// 2. maxMRB-Add-r17: integer
	{
		if ie.MaxMRB_Add_r17 != nil {
			if err := e.EncodeInteger(*ie.MaxMRB_Add_r17, mBSParametersR17MaxMRBAddR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MBS_Parameters_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSParametersR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. maxMRB-Add-r17: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(mBSParametersR17MaxMRBAddR17Constraints)
			if err != nil {
				return err
			}
			ie.MaxMRB_Add_r17 = &v
		}
	}

	return nil
}
