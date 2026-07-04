// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CO-DurationsPerCell-r17 (line 15194).

var cODurationsPerCellR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "servingCellId-r17"},
		{Name: "positionInDCI-r17"},
		{Name: "subcarrierSpacing-r17"},
		{Name: "co-DurationList-r17"},
	},
}

var cODurationsPerCellR17PositionInDCIR17Constraints = per.Constrained(0, common.MaxSFI_DCI_PayloadSize_1)

var cODurationsPerCellR17CoDurationListR17Constraints = per.SizeRange(1, 64)

type CO_DurationsPerCell_r17 struct {
	ServingCellId_r17     ServCellIndex
	PositionInDCI_r17     int64
	SubcarrierSpacing_r17 SubcarrierSpacing
	Co_DurationList_r17   []CO_Duration_r17
}

func (ie *CO_DurationsPerCell_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cODurationsPerCellR17Constraints)
	_ = seq

	// 1. servingCellId-r17: ref
	{
		if err := ie.ServingCellId_r17.Encode(e); err != nil {
			return err
		}
	}

	// 2. positionInDCI-r17: integer
	{
		if err := e.EncodeInteger(ie.PositionInDCI_r17, cODurationsPerCellR17PositionInDCIR17Constraints); err != nil {
			return err
		}
	}

	// 3. subcarrierSpacing-r17: ref
	{
		if err := ie.SubcarrierSpacing_r17.Encode(e); err != nil {
			return err
		}
	}

	// 4. co-DurationList-r17: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(cODurationsPerCellR17CoDurationListR17Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Co_DurationList_r17))); err != nil {
			return err
		}
		for i := range ie.Co_DurationList_r17 {
			if err := ie.Co_DurationList_r17[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CO_DurationsPerCell_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cODurationsPerCellR17Constraints)
	_ = seq

	// 1. servingCellId-r17: ref
	{
		if err := ie.ServingCellId_r17.Decode(d); err != nil {
			return err
		}
	}

	// 2. positionInDCI-r17: integer
	{
		v1, err := d.DecodeInteger(cODurationsPerCellR17PositionInDCIR17Constraints)
		if err != nil {
			return err
		}
		ie.PositionInDCI_r17 = v1
	}

	// 3. subcarrierSpacing-r17: ref
	{
		if err := ie.SubcarrierSpacing_r17.Decode(d); err != nil {
			return err
		}
	}

	// 4. co-DurationList-r17: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(cODurationsPerCellR17CoDurationListR17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Co_DurationList_r17 = make([]CO_Duration_r17, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Co_DurationList_r17[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
