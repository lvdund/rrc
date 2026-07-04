// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultsPerCarrierIdleNR-r16 (line 9950).

var measResultsPerCarrierIdleNRR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreq-r16"},
		{Name: "measResultsPerCellListIdleNR-r16"},
	},
}

var measResultsPerCarrierIdleNRR16MeasResultsPerCellListIdleNRR16Constraints = per.SizeRange(1, common.MaxCellMeasIdle_r16)

type MeasResultsPerCarrierIdleNR_r16 struct {
	CarrierFreq_r16                  ARFCN_ValueNR
	MeasResultsPerCellListIdleNR_r16 []MeasResultsPerCellIdleNR_r16
}

func (ie *MeasResultsPerCarrierIdleNR_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultsPerCarrierIdleNRR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. carrierFreq-r16: ref
	{
		if err := ie.CarrierFreq_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. measResultsPerCellListIdleNR-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(measResultsPerCarrierIdleNRR16MeasResultsPerCellListIdleNRR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.MeasResultsPerCellListIdleNR_r16))); err != nil {
			return err
		}
		for i := range ie.MeasResultsPerCellListIdleNR_r16 {
			if err := ie.MeasResultsPerCellListIdleNR_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasResultsPerCarrierIdleNR_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultsPerCarrierIdleNRR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. carrierFreq-r16: ref
	{
		if err := ie.CarrierFreq_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. measResultsPerCellListIdleNR-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(measResultsPerCarrierIdleNRR16MeasResultsPerCellListIdleNRR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.MeasResultsPerCellListIdleNR_r16 = make([]MeasResultsPerCellIdleNR_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.MeasResultsPerCellListIdleNR_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
