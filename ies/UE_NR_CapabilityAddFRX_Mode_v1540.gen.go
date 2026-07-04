// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-CapabilityAddFRX-Mode-v1540 (line 25923).

var uENRCapabilityAddFRXModeV1540Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ims-ParametersFRX-Diff", Optional: true},
	},
}

type UE_NR_CapabilityAddFRX_Mode_v1540 struct {
	Ims_ParametersFRX_Diff *IMS_ParametersFRX_Diff
}

func (ie *UE_NR_CapabilityAddFRX_Mode_v1540) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityAddFRXModeV1540Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ims_ParametersFRX_Diff != nil}); err != nil {
		return err
	}

	// 2. ims-ParametersFRX-Diff: ref
	{
		if ie.Ims_ParametersFRX_Diff != nil {
			if err := ie.Ims_ParametersFRX_Diff.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_NR_CapabilityAddFRX_Mode_v1540) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityAddFRXModeV1540Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ims-ParametersFRX-Diff: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ims_ParametersFRX_Diff = new(IMS_ParametersFRX_Diff)
			if err := ie.Ims_ParametersFRX_Diff.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
