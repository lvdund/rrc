// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNR-v1560 (line 17296).

var cAParametersNRV1560Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "diffNumerologyWithinPUCCH-GroupLargerSCS", Optional: true},
	},
}

const (
	CA_ParametersNR_v1560_DiffNumerologyWithinPUCCH_GroupLargerSCS_Supported = 0
)

var cAParametersNRV1560DiffNumerologyWithinPUCCHGroupLargerSCSConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1560_DiffNumerologyWithinPUCCH_GroupLargerSCS_Supported},
}

type CA_ParametersNR_v1560 struct {
	DiffNumerologyWithinPUCCH_GroupLargerSCS *int64
}

func (ie *CA_ParametersNR_v1560) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRV1560Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DiffNumerologyWithinPUCCH_GroupLargerSCS != nil}); err != nil {
		return err
	}

	// 2. diffNumerologyWithinPUCCH-GroupLargerSCS: enumerated
	{
		if ie.DiffNumerologyWithinPUCCH_GroupLargerSCS != nil {
			if err := e.EncodeEnumerated(*ie.DiffNumerologyWithinPUCCH_GroupLargerSCS, cAParametersNRV1560DiffNumerologyWithinPUCCHGroupLargerSCSConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNR_v1560) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRV1560Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. diffNumerologyWithinPUCCH-GroupLargerSCS: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1560DiffNumerologyWithinPUCCHGroupLargerSCSConstraints)
			if err != nil {
				return err
			}
			ie.DiffNumerologyWithinPUCCH_GroupLargerSCS = &idx
		}
	}

	return nil
}
