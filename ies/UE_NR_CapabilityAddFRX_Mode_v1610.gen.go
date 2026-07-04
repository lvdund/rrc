// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-CapabilityAddFRX-Mode-v1610 (line 25927).

var uENRCapabilityAddFRXModeV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "powSav-ParametersFRX-Diff-r16", Optional: true},
		{Name: "mac-ParametersFRX-Diff-r16", Optional: true},
	},
}

type UE_NR_CapabilityAddFRX_Mode_v1610 struct {
	PowSav_ParametersFRX_Diff_r16 *PowSav_ParametersFRX_Diff_r16
	Mac_ParametersFRX_Diff_r16    *MAC_ParametersFRX_Diff_r16
}

func (ie *UE_NR_CapabilityAddFRX_Mode_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityAddFRXModeV1610Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PowSav_ParametersFRX_Diff_r16 != nil, ie.Mac_ParametersFRX_Diff_r16 != nil}); err != nil {
		return err
	}

	// 2. powSav-ParametersFRX-Diff-r16: ref
	{
		if ie.PowSav_ParametersFRX_Diff_r16 != nil {
			if err := ie.PowSav_ParametersFRX_Diff_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. mac-ParametersFRX-Diff-r16: ref
	{
		if ie.Mac_ParametersFRX_Diff_r16 != nil {
			if err := ie.Mac_ParametersFRX_Diff_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_NR_CapabilityAddFRX_Mode_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityAddFRXModeV1610Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. powSav-ParametersFRX-Diff-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.PowSav_ParametersFRX_Diff_r16 = new(PowSav_ParametersFRX_Diff_r16)
			if err := ie.PowSav_ParametersFRX_Diff_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. mac-ParametersFRX-Diff-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Mac_ParametersFRX_Diff_r16 = new(MAC_ParametersFRX_Diff_r16)
			if err := ie.Mac_ParametersFRX_Diff_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
