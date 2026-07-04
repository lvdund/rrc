// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CA-ParametersNRDC-v1650 (line 18406).

var cAParametersNRDCV1650Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedCellGrouping-r16", Optional: true},
	},
}

var cAParametersNRDCV1650SupportedCellGroupingR16Constraints = per.SizeRange(1, common.MaxCellGroupings_r16)

type CA_ParametersNRDC_v1650 struct {
	SupportedCellGrouping_r16 *per.BitString
}

func (ie *CA_ParametersNRDC_v1650) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRDCV1650Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedCellGrouping_r16 != nil}); err != nil {
		return err
	}

	// 2. supportedCellGrouping-r16: bit-string
	{
		if ie.SupportedCellGrouping_r16 != nil {
			if err := e.EncodeBitString(*ie.SupportedCellGrouping_r16, cAParametersNRDCV1650SupportedCellGroupingR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNRDC_v1650) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRDCV1650Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedCellGrouping-r16: bit-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeBitString(cAParametersNRDCV1650SupportedCellGroupingR16Constraints)
			if err != nil {
				return err
			}
			ie.SupportedCellGrouping_r16 = &v
		}
	}

	return nil
}
