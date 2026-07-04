// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LTM-CSI-IM-ResourceSet-r19 (line 8938).

var lTMCSIIMResourceSetR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ltm-CSI-IM-ResourceSetId-r19"},
	},
}

type LTM_CSI_IM_ResourceSet_r19 struct {
	Ltm_CSI_IM_ResourceSetId_r19 CSI_IM_ResourceSetId
}

func (ie *LTM_CSI_IM_ResourceSet_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTMCSIIMResourceSetR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. ltm-CSI-IM-ResourceSetId-r19: ref
	{
		if err := ie.Ltm_CSI_IM_ResourceSetId_r19.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *LTM_CSI_IM_ResourceSet_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTMCSIIMResourceSetR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. ltm-CSI-IM-ResourceSetId-r19: ref
	{
		if err := ie.Ltm_CSI_IM_ResourceSetId_r19.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
