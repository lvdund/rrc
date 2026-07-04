// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CO-DurationsPerCell-r16 (line 15187).

var cODurationsPerCellR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "servingCellId-r16"},
		{Name: "positionInDCI-r16"},
		{Name: "subcarrierSpacing-r16"},
		{Name: "co-DurationList-r16"},
	},
}

var cODurationsPerCellR16PositionInDCIR16Constraints = per.Constrained(0, common.MaxSFI_DCI_PayloadSize_1)

var cODurationsPerCellR16CoDurationListR16Constraints = per.SizeRange(1, 64)

type CO_DurationsPerCell_r16 struct {
	ServingCellId_r16     ServCellIndex
	PositionInDCI_r16     int64
	SubcarrierSpacing_r16 SubcarrierSpacing
	Co_DurationList_r16   []CO_Duration_r16
}

func (ie *CO_DurationsPerCell_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cODurationsPerCellR16Constraints)
	_ = seq

	// 1. servingCellId-r16: ref
	{
		if err := ie.ServingCellId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. positionInDCI-r16: integer
	{
		if err := e.EncodeInteger(ie.PositionInDCI_r16, cODurationsPerCellR16PositionInDCIR16Constraints); err != nil {
			return err
		}
	}

	// 3. subcarrierSpacing-r16: ref
	{
		if err := ie.SubcarrierSpacing_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. co-DurationList-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(cODurationsPerCellR16CoDurationListR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Co_DurationList_r16))); err != nil {
			return err
		}
		for i := range ie.Co_DurationList_r16 {
			if err := ie.Co_DurationList_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CO_DurationsPerCell_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cODurationsPerCellR16Constraints)
	_ = seq

	// 1. servingCellId-r16: ref
	{
		if err := ie.ServingCellId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. positionInDCI-r16: integer
	{
		v1, err := d.DecodeInteger(cODurationsPerCellR16PositionInDCIR16Constraints)
		if err != nil {
			return err
		}
		ie.PositionInDCI_r16 = v1
	}

	// 3. subcarrierSpacing-r16: ref
	{
		if err := ie.SubcarrierSpacing_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. co-DurationList-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(cODurationsPerCellR16CoDurationListR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Co_DurationList_r16 = make([]CO_Duration_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Co_DurationList_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
