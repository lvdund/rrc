// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CellGrouping-r16 (line 25513).

var cellGroupingR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mcg-r16"},
		{Name: "scg-r16"},
		{Name: "mode-r16"},
	},
}

var cellGroupingR16McgR16Constraints = per.SizeRange(1, common.MaxBands)

var cellGroupingR16ScgR16Constraints = per.SizeRange(1, common.MaxBands)

const (
	CellGrouping_r16_Mode_r16_Sync  = 0
	CellGrouping_r16_Mode_r16_Async = 1
)

var cellGroupingR16ModeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellGrouping_r16_Mode_r16_Sync, CellGrouping_r16_Mode_r16_Async},
}

type CellGrouping_r16 struct {
	Mcg_r16  []FreqBandIndicatorNR
	Scg_r16  []FreqBandIndicatorNR
	Mode_r16 int64
}

func (ie *CellGrouping_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellGroupingR16Constraints)
	_ = seq

	// 1. mcg-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(cellGroupingR16McgR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Mcg_r16))); err != nil {
			return err
		}
		for i := range ie.Mcg_r16 {
			if err := ie.Mcg_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 2. scg-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(cellGroupingR16ScgR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Scg_r16))); err != nil {
			return err
		}
		for i := range ie.Scg_r16 {
			if err := ie.Scg_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. mode-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Mode_r16, cellGroupingR16ModeR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CellGrouping_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellGroupingR16Constraints)
	_ = seq

	// 1. mcg-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(cellGroupingR16McgR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Mcg_r16 = make([]FreqBandIndicatorNR, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Mcg_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 2. scg-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(cellGroupingR16ScgR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Scg_r16 = make([]FreqBandIndicatorNR, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Scg_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. mode-r16: enumerated
	{
		v2, err := d.DecodeEnumerated(cellGroupingR16ModeR16Constraints)
		if err != nil {
			return err
		}
		ie.Mode_r16 = v2
	}

	return nil
}
