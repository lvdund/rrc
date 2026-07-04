// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasReportQuantity-r16 (line 13964).

var measReportQuantityR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cbr-r16"},
	},
}

type MeasReportQuantity_r16 struct {
	Cbr_r16 bool
}

func (ie *MeasReportQuantity_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measReportQuantityR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. cbr-r16: boolean
	{
		if err := e.EncodeBoolean(ie.Cbr_r16); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasReportQuantity_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measReportQuantityR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. cbr-r16: boolean
	{
		v0, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.Cbr_r16 = v0
	}

	return nil
}
