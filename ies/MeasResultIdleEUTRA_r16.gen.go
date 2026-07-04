// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultIdleEUTRA-r16 (line 9914).

var measResultIdleEUTRAR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "measResultsPerCarrierListIdleEUTRA-r16"},
	},
}

var measResultIdleEUTRAR16MeasResultsPerCarrierListIdleEUTRAR16Constraints = per.SizeRange(1, common.MaxFreqIdle_r16)

type MeasResultIdleEUTRA_r16 struct {
	MeasResultsPerCarrierListIdleEUTRA_r16 []MeasResultsPerCarrierIdleEUTRA_r16
}

func (ie *MeasResultIdleEUTRA_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultIdleEUTRAR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. measResultsPerCarrierListIdleEUTRA-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(measResultIdleEUTRAR16MeasResultsPerCarrierListIdleEUTRAR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.MeasResultsPerCarrierListIdleEUTRA_r16))); err != nil {
			return err
		}
		for i := range ie.MeasResultsPerCarrierListIdleEUTRA_r16 {
			if err := ie.MeasResultsPerCarrierListIdleEUTRA_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasResultIdleEUTRA_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultIdleEUTRAR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. measResultsPerCarrierListIdleEUTRA-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(measResultIdleEUTRAR16MeasResultsPerCarrierListIdleEUTRAR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.MeasResultsPerCarrierListIdleEUTRA_r16 = make([]MeasResultsPerCarrierIdleEUTRA_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.MeasResultsPerCarrierListIdleEUTRA_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
