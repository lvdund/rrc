// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultLoggingNR-r16 (line 3424).

var measResultLoggingNRR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "physCellId-r16"},
		{Name: "resultsSSB-Cell-r16"},
		{Name: "numberOfGoodSSB-r16", Optional: true},
	},
}

var measResultLoggingNRR16NumberOfGoodSSBR16Constraints = per.Constrained(1, common.MaxNrofSSBs_r16)

type MeasResultLoggingNR_r16 struct {
	PhysCellId_r16      PhysCellId
	ResultsSSB_Cell_r16 MeasQuantityResults
	NumberOfGoodSSB_r16 *int64
}

func (ie *MeasResultLoggingNR_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultLoggingNRR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.NumberOfGoodSSB_r16 != nil}); err != nil {
		return err
	}

	// 2. physCellId-r16: ref
	{
		if err := ie.PhysCellId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. resultsSSB-Cell-r16: ref
	{
		if err := ie.ResultsSSB_Cell_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. numberOfGoodSSB-r16: integer
	{
		if ie.NumberOfGoodSSB_r16 != nil {
			if err := e.EncodeInteger(*ie.NumberOfGoodSSB_r16, measResultLoggingNRR16NumberOfGoodSSBR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasResultLoggingNR_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultLoggingNRR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. physCellId-r16: ref
	{
		if err := ie.PhysCellId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. resultsSSB-Cell-r16: ref
	{
		if err := ie.ResultsSSB_Cell_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. numberOfGoodSSB-r16: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(measResultLoggingNRR16NumberOfGoodSSBR16Constraints)
			if err != nil {
				return err
			}
			ie.NumberOfGoodSSB_r16 = &v
		}
	}

	return nil
}
