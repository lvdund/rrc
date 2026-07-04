// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IMS-Parameters-v1700 (line 20864).

var iMSParametersV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ims-ParametersFR2-2-r17", Optional: true},
	},
}

type IMS_Parameters_v1700 struct {
	Ims_ParametersFR2_2_r17 *IMS_ParametersFR2_2_r17
}

func (ie *IMS_Parameters_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iMSParametersV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ims_ParametersFR2_2_r17 != nil}); err != nil {
		return err
	}

	// 2. ims-ParametersFR2-2-r17: ref
	{
		if ie.Ims_ParametersFR2_2_r17 != nil {
			if err := ie.Ims_ParametersFR2_2_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *IMS_Parameters_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iMSParametersV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ims-ParametersFR2-2-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ims_ParametersFR2_2_r17 = new(IMS_ParametersFR2_2_r17)
			if err := ie.Ims_ParametersFR2_2_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
