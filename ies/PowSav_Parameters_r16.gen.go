// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PowSav-Parameters-r16 (line 23441).

var powSavParametersR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "powSav-ParametersCommon-r16", Optional: true},
		{Name: "powSav-ParametersFRX-Diff-r16", Optional: true},
	},
}

type PowSav_Parameters_r16 struct {
	PowSav_ParametersCommon_r16   *PowSav_ParametersCommon_r16
	PowSav_ParametersFRX_Diff_r16 *PowSav_ParametersFRX_Diff_r16
}

func (ie *PowSav_Parameters_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(powSavParametersR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PowSav_ParametersCommon_r16 != nil, ie.PowSav_ParametersFRX_Diff_r16 != nil}); err != nil {
		return err
	}

	// 3. powSav-ParametersCommon-r16: ref
	{
		if ie.PowSav_ParametersCommon_r16 != nil {
			if err := ie.PowSav_ParametersCommon_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. powSav-ParametersFRX-Diff-r16: ref
	{
		if ie.PowSav_ParametersFRX_Diff_r16 != nil {
			if err := ie.PowSav_ParametersFRX_Diff_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PowSav_Parameters_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(powSavParametersR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. powSav-ParametersCommon-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.PowSav_ParametersCommon_r16 = new(PowSav_ParametersCommon_r16)
			if err := ie.PowSav_ParametersCommon_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. powSav-ParametersFRX-Diff-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.PowSav_ParametersFRX_Diff_r16 = new(PowSav_ParametersFRX_Diff_r16)
			if err := ie.PowSav_ParametersFRX_Diff_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
