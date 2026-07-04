// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParametersMRDC-FRX-Diff (line 21465).

var measAndMobParametersMRDCFRXDiffConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "simultaneousRxDataSSB-DiffNumerology", Optional: true},
	},
}

const (
	MeasAndMobParametersMRDC_FRX_Diff_SimultaneousRxDataSSB_DiffNumerology_Supported = 0
)

var measAndMobParametersMRDCFRXDiffSimultaneousRxDataSSBDiffNumerologyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_FRX_Diff_SimultaneousRxDataSSB_DiffNumerology_Supported},
}

type MeasAndMobParametersMRDC_FRX_Diff struct {
	SimultaneousRxDataSSB_DiffNumerology *int64
}

func (ie *MeasAndMobParametersMRDC_FRX_Diff) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersMRDCFRXDiffConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SimultaneousRxDataSSB_DiffNumerology != nil}); err != nil {
		return err
	}

	// 2. simultaneousRxDataSSB-DiffNumerology: enumerated
	{
		if ie.SimultaneousRxDataSSB_DiffNumerology != nil {
			if err := e.EncodeEnumerated(*ie.SimultaneousRxDataSSB_DiffNumerology, measAndMobParametersMRDCFRXDiffSimultaneousRxDataSSBDiffNumerologyConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasAndMobParametersMRDC_FRX_Diff) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersMRDCFRXDiffConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. simultaneousRxDataSSB-DiffNumerology: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(measAndMobParametersMRDCFRXDiffSimultaneousRxDataSSBDiffNumerologyConstraints)
			if err != nil {
				return err
			}
			ie.SimultaneousRxDataSSB_DiffNumerology = &idx
		}
	}

	return nil
}
