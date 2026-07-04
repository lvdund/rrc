// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MAC-Parameters-v17c0 (line 20961).

var mACParametersV17c0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "directSCellActivationWithTCI-r17", Optional: true},
	},
}

const (
	MAC_Parameters_V17c0_DirectSCellActivationWithTCI_r17_Supported = 0
)

var mACParametersV17c0DirectSCellActivationWithTCIR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_Parameters_V17c0_DirectSCellActivationWithTCI_r17_Supported},
}

type MAC_Parameters_V17c0 struct {
	DirectSCellActivationWithTCI_r17 *int64
}

func (ie *MAC_Parameters_V17c0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mACParametersV17c0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DirectSCellActivationWithTCI_r17 != nil}); err != nil {
		return err
	}

	// 2. directSCellActivationWithTCI-r17: enumerated
	{
		if ie.DirectSCellActivationWithTCI_r17 != nil {
			if err := e.EncodeEnumerated(*ie.DirectSCellActivationWithTCI_r17, mACParametersV17c0DirectSCellActivationWithTCIR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MAC_Parameters_V17c0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mACParametersV17c0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. directSCellActivationWithTCI-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(mACParametersV17c0DirectSCellActivationWithTCIR17Constraints)
			if err != nil {
				return err
			}
			ie.DirectSCellActivationWithTCI_r17 = &idx
		}
	}

	return nil
}
