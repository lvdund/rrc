// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PowSav-Parameters-v1700 (line 23447).

var powSavParametersV1700Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "powSav-ParametersFR2-2-r17", Optional: true},
	},
}

type PowSav_Parameters_v1700 struct {
	PowSav_ParametersFR2_2_r17 *PowSav_ParametersFR2_2_r17
}

func (ie *PowSav_Parameters_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(powSavParametersV1700Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PowSav_ParametersFR2_2_r17 != nil}); err != nil {
		return err
	}

	// 3. powSav-ParametersFR2-2-r17: ref
	{
		if ie.PowSav_ParametersFR2_2_r17 != nil {
			if err := ie.PowSav_ParametersFR2_2_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PowSav_Parameters_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(powSavParametersV1700Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. powSav-ParametersFR2-2-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.PowSav_ParametersFR2_2_r17 = new(PowSav_ParametersFR2_2_r17)
			if err := ie.PowSav_ParametersFR2_2_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
