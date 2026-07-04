// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultsPerCellIdleNR-r16 (line 9956).

var measResultsPerCellIdleNRR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "physCellId-r16"},
		{Name: "measIdleResultNR-r16"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var measResultsPerCellIdleNRR16MeasIdleResultNRR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rsrp-Result-r16", Optional: true},
		{Name: "rsrq-Result-r16", Optional: true},
		{Name: "resultsSSB-Indexes-r16", Optional: true},
	},
}

type MeasResultsPerCellIdleNR_r16 struct {
	PhysCellId_r16       PhysCellId
	MeasIdleResultNR_r16 struct {
		Rsrp_Result_r16        *RSRP_Range
		Rsrq_Result_r16        *RSRQ_Range
		ResultsSSB_Indexes_r16 *ResultsPerSSB_IndexList_r16
	}
	ValidityStatus_r18 *MeasurementValidityDuration_r18
}

func (ie *MeasResultsPerCellIdleNR_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultsPerCellIdleNRR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ValidityStatus_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. physCellId-r16: ref
	{
		if err := ie.PhysCellId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. measIdleResultNR-r16: sequence
	{
		{
			c := &ie.MeasIdleResultNR_r16
			measResultsPerCellIdleNRR16MeasIdleResultNRR16Seq := e.NewSequenceEncoder(measResultsPerCellIdleNRR16MeasIdleResultNRR16Constraints)
			if err := measResultsPerCellIdleNRR16MeasIdleResultNRR16Seq.EncodePreamble([]bool{c.Rsrp_Result_r16 != nil, c.Rsrq_Result_r16 != nil, c.ResultsSSB_Indexes_r16 != nil}); err != nil {
				return err
			}
			if c.Rsrp_Result_r16 != nil {
				if err := c.Rsrp_Result_r16.Encode(e); err != nil {
					return err
				}
			}
			if c.Rsrq_Result_r16 != nil {
				if err := c.Rsrq_Result_r16.Encode(e); err != nil {
					return err
				}
			}
			if c.ResultsSSB_Indexes_r16 != nil {
				if err := c.ResultsSSB_Indexes_r16.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "validityStatus-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ValidityStatus_r18 != nil}); err != nil {
				return err
			}

			if ie.ValidityStatus_r18 != nil {
				if err := ie.ValidityStatus_r18.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasResultsPerCellIdleNR_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultsPerCellIdleNRR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. physCellId-r16: ref
	{
		if err := ie.PhysCellId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. measIdleResultNR-r16: sequence
	{
		{
			c := &ie.MeasIdleResultNR_r16
			measResultsPerCellIdleNRR16MeasIdleResultNRR16Seq := d.NewSequenceDecoder(measResultsPerCellIdleNRR16MeasIdleResultNRR16Constraints)
			if err := measResultsPerCellIdleNRR16MeasIdleResultNRR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if measResultsPerCellIdleNRR16MeasIdleResultNRR16Seq.IsComponentPresent(0) {
				c.Rsrp_Result_r16 = new(RSRP_Range)
				if err := (*c.Rsrp_Result_r16).Decode(d); err != nil {
					return err
				}
			}
			if measResultsPerCellIdleNRR16MeasIdleResultNRR16Seq.IsComponentPresent(1) {
				c.Rsrq_Result_r16 = new(RSRQ_Range)
				if err := (*c.Rsrq_Result_r16).Decode(d); err != nil {
					return err
				}
			}
			if measResultsPerCellIdleNRR16MeasIdleResultNRR16Seq.IsComponentPresent(2) {
				c.ResultsSSB_Indexes_r16 = new(ResultsPerSSB_IndexList_r16)
				if err := (*c.ResultsSSB_Indexes_r16).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "validityStatus-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.ValidityStatus_r18 = new(MeasurementValidityDuration_r18)
			if err := ie.ValidityStatus_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
