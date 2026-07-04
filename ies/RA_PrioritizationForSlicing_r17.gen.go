// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RA-PrioritizationForSlicing-r17 (line 13092).

var rAPrioritizationForSlicingR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ra-PrioritizationSliceInfoList-r17"},
	},
}

type RA_PrioritizationForSlicing_r17 struct {
	Ra_PrioritizationSliceInfoList_r17 RA_PrioritizationSliceInfoList_r17
}

func (ie *RA_PrioritizationForSlicing_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rAPrioritizationForSlicingR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. ra-PrioritizationSliceInfoList-r17: ref
	{
		if err := ie.Ra_PrioritizationSliceInfoList_r17.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RA_PrioritizationForSlicing_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rAPrioritizationForSlicingR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. ra-PrioritizationSliceInfoList-r17: ref
	{
		if err := ie.Ra_PrioritizationSliceInfoList_r17.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
