// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultL1-r19 (line 3525).

var measResultL1R19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "resultsSSB-Indexes-r19", Optional: true},
	},
}

type MeasResultL1_r19 struct {
	ResultsSSB_Indexes_r19 *ResultsPerSSB_IndexList
}

func (ie *MeasResultL1_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultL1R19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ResultsSSB_Indexes_r19 != nil}); err != nil {
		return err
	}

	// 3. resultsSSB-Indexes-r19: ref
	{
		if ie.ResultsSSB_Indexes_r19 != nil {
			if err := ie.ResultsSSB_Indexes_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasResultL1_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultL1R19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. resultsSSB-Indexes-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.ResultsSSB_Indexes_r19 = new(ResultsPerSSB_IndexList)
			if err := ie.ResultsSSB_Indexes_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
