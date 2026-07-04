// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultIdleNR-r16 (line 9940).

var measResultIdleNRR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "measResultServingCell-r16"},
		{Name: "measResultsPerCarrierListIdleNR-r16", Optional: true},
	},
}

var measResultIdleNRR16MeasResultsPerCarrierListIdleNRR16Constraints = per.SizeRange(1, common.MaxFreqIdle_r16)

var measResultIdleNRR16MeasResultServingCellR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rsrp-Result-r16", Optional: true},
		{Name: "rsrq-Result-r16", Optional: true},
		{Name: "resultsSSB-Indexes-r16", Optional: true},
	},
}

type MeasResultIdleNR_r16 struct {
	MeasResultServingCell_r16 struct {
		Rsrp_Result_r16        *RSRP_Range
		Rsrq_Result_r16        *RSRQ_Range
		ResultsSSB_Indexes_r16 *ResultsPerSSB_IndexList_r16
	}
	MeasResultsPerCarrierListIdleNR_r16 []MeasResultsPerCarrierIdleNR_r16
}

func (ie *MeasResultIdleNR_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultIdleNRR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasResultsPerCarrierListIdleNR_r16 != nil}); err != nil {
		return err
	}

	// 3. measResultServingCell-r16: sequence
	{
		{
			c := &ie.MeasResultServingCell_r16
			measResultIdleNRR16MeasResultServingCellR16Seq := e.NewSequenceEncoder(measResultIdleNRR16MeasResultServingCellR16Constraints)
			if err := measResultIdleNRR16MeasResultServingCellR16Seq.EncodePreamble([]bool{c.Rsrp_Result_r16 != nil, c.Rsrq_Result_r16 != nil, c.ResultsSSB_Indexes_r16 != nil}); err != nil {
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

	// 4. measResultsPerCarrierListIdleNR-r16: sequence-of
	{
		if ie.MeasResultsPerCarrierListIdleNR_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(measResultIdleNRR16MeasResultsPerCarrierListIdleNRR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.MeasResultsPerCarrierListIdleNR_r16))); err != nil {
				return err
			}
			for i := range ie.MeasResultsPerCarrierListIdleNR_r16 {
				if err := ie.MeasResultsPerCarrierListIdleNR_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *MeasResultIdleNR_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultIdleNRR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. measResultServingCell-r16: sequence
	{
		{
			c := &ie.MeasResultServingCell_r16
			measResultIdleNRR16MeasResultServingCellR16Seq := d.NewSequenceDecoder(measResultIdleNRR16MeasResultServingCellR16Constraints)
			if err := measResultIdleNRR16MeasResultServingCellR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if measResultIdleNRR16MeasResultServingCellR16Seq.IsComponentPresent(0) {
				c.Rsrp_Result_r16 = new(RSRP_Range)
				if err := (*c.Rsrp_Result_r16).Decode(d); err != nil {
					return err
				}
			}
			if measResultIdleNRR16MeasResultServingCellR16Seq.IsComponentPresent(1) {
				c.Rsrq_Result_r16 = new(RSRQ_Range)
				if err := (*c.Rsrq_Result_r16).Decode(d); err != nil {
					return err
				}
			}
			if measResultIdleNRR16MeasResultServingCellR16Seq.IsComponentPresent(2) {
				c.ResultsSSB_Indexes_r16 = new(ResultsPerSSB_IndexList_r16)
				if err := (*c.ResultsSSB_Indexes_r16).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. measResultsPerCarrierListIdleNR-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(measResultIdleNRR16MeasResultsPerCarrierListIdleNRR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.MeasResultsPerCarrierListIdleNR_r16 = make([]MeasResultsPerCarrierIdleNR_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.MeasResultsPerCarrierListIdleNR_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
