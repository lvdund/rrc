// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultServingCell-r16 (line 3038).

var measResultServingCellR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resultsSSB-Cell"},
		{Name: "resultsSSB", Optional: true},
	},
}

type MeasResultServingCell_r16 struct {
	ResultsSSB_Cell MeasQuantityResults
	ResultsSSB      *struct {
		Best_Ssb_Index   SSB_Index
		Best_Ssb_Results MeasQuantityResults
		NumberOfGoodSSB  int64
	}
}

func (ie *MeasResultServingCell_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultServingCellR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ResultsSSB != nil}); err != nil {
		return err
	}

	// 2. resultsSSB-Cell: ref
	{
		if err := ie.ResultsSSB_Cell.Encode(e); err != nil {
			return err
		}
	}

	// 3. resultsSSB: sequence
	{
		if ie.ResultsSSB != nil {
			c := ie.ResultsSSB
			if err := c.Best_Ssb_Index.Encode(e); err != nil {
				return err
			}
			if err := c.Best_Ssb_Results.Encode(e); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.NumberOfGoodSSB, per.Constrained(1, common.MaxNrofSSBs_r16)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasResultServingCell_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultServingCellR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. resultsSSB-Cell: ref
	{
		if err := ie.ResultsSSB_Cell.Decode(d); err != nil {
			return err
		}
	}

	// 3. resultsSSB: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.ResultsSSB = &struct {
				Best_Ssb_Index   SSB_Index
				Best_Ssb_Results MeasQuantityResults
				NumberOfGoodSSB  int64
			}{}
			c := ie.ResultsSSB
			{
				if err := c.Best_Ssb_Index.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.Best_Ssb_Results.Decode(d); err != nil {
					return err
				}
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, common.MaxNrofSSBs_r16))
				if err != nil {
					return err
				}
				c.NumberOfGoodSSB = v
			}
		}
	}

	return nil
}
