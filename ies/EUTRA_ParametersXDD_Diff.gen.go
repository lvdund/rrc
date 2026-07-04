// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EUTRA-ParametersXDD-Diff (line 20922).

var eUTRAParametersXDDDiffConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rsrqMeasWidebandEUTRA", Optional: true},
	},
}

const (
	EUTRA_ParametersXDD_Diff_RsrqMeasWidebandEUTRA_Supported = 0
)

var eUTRAParametersXDDDiffRsrqMeasWidebandEUTRAConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EUTRA_ParametersXDD_Diff_RsrqMeasWidebandEUTRA_Supported},
}

type EUTRA_ParametersXDD_Diff struct {
	RsrqMeasWidebandEUTRA *int64
}

func (ie *EUTRA_ParametersXDD_Diff) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eUTRAParametersXDDDiffConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RsrqMeasWidebandEUTRA != nil}); err != nil {
		return err
	}

	// 3. rsrqMeasWidebandEUTRA: enumerated
	{
		if ie.RsrqMeasWidebandEUTRA != nil {
			if err := e.EncodeEnumerated(*ie.RsrqMeasWidebandEUTRA, eUTRAParametersXDDDiffRsrqMeasWidebandEUTRAConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *EUTRA_ParametersXDD_Diff) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eUTRAParametersXDDDiffConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. rsrqMeasWidebandEUTRA: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(eUTRAParametersXDDDiffRsrqMeasWidebandEUTRAConstraints)
			if err != nil {
				return err
			}
			ie.RsrqMeasWidebandEUTRA = &idx
		}
	}

	return nil
}
