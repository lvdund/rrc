// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-CapabilityAddXDD-Mode-v1530 (line 25914).

var uENRCapabilityAddXDDModeV1530Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "eutra-ParametersXDD-Diff"},
	},
}

type UE_NR_CapabilityAddXDD_Mode_v1530 struct {
	Eutra_ParametersXDD_Diff EUTRA_ParametersXDD_Diff
}

func (ie *UE_NR_CapabilityAddXDD_Mode_v1530) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityAddXDDModeV1530Constraints)
	_ = seq

	// 1. eutra-ParametersXDD-Diff: ref
	{
		if err := ie.Eutra_ParametersXDD_Diff.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UE_NR_CapabilityAddXDD_Mode_v1530) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityAddXDDModeV1530Constraints)
	_ = seq

	// 1. eutra-ParametersXDD-Diff: ref
	{
		if err := ie.Eutra_ParametersXDD_Diff.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
