// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultsPerCarrierIdleEUTRA-r16 (line 9919).

var measResultsPerCarrierIdleEUTRAR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreqEUTRA-r16"},
		{Name: "measResultsPerCellListIdleEUTRA-r16"},
	},
}

var measResultsPerCarrierIdleEUTRAR16MeasResultsPerCellListIdleEUTRAR16Constraints = per.SizeRange(1, common.MaxCellMeasIdle_r16)

type MeasResultsPerCarrierIdleEUTRA_r16 struct {
	CarrierFreqEUTRA_r16                ARFCN_ValueEUTRA
	MeasResultsPerCellListIdleEUTRA_r16 []MeasResultsPerCellIdleEUTRA_r16
}

func (ie *MeasResultsPerCarrierIdleEUTRA_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultsPerCarrierIdleEUTRAR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. carrierFreqEUTRA-r16: ref
	{
		if err := ie.CarrierFreqEUTRA_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. measResultsPerCellListIdleEUTRA-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(measResultsPerCarrierIdleEUTRAR16MeasResultsPerCellListIdleEUTRAR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.MeasResultsPerCellListIdleEUTRA_r16))); err != nil {
			return err
		}
		for i := range ie.MeasResultsPerCellListIdleEUTRA_r16 {
			if err := ie.MeasResultsPerCellListIdleEUTRA_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasResultsPerCarrierIdleEUTRA_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultsPerCarrierIdleEUTRAR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. carrierFreqEUTRA-r16: ref
	{
		if err := ie.CarrierFreqEUTRA_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. measResultsPerCellListIdleEUTRA-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(measResultsPerCarrierIdleEUTRAR16MeasResultsPerCellListIdleEUTRAR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.MeasResultsPerCellListIdleEUTRA_r16 = make([]MeasResultsPerCellIdleEUTRA_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.MeasResultsPerCellListIdleEUTRA_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
