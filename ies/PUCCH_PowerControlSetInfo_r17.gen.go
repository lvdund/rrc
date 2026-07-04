// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-PowerControlSetInfo-r17 (line 12275).

var pUCCHPowerControlSetInfoR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pucch-PowerControlSetInfoId-r17"},
		{Name: "p0-PUCCH-Id-r17"},
		{Name: "pucch-ClosedLoopIndex-r17"},
		{Name: "pucch-PathlossReferenceRS-Id-r17"},
	},
}

const (
	PUCCH_PowerControlSetInfo_r17_Pucch_ClosedLoopIndex_r17_I0 = 0
	PUCCH_PowerControlSetInfo_r17_Pucch_ClosedLoopIndex_r17_I1 = 1
)

var pUCCHPowerControlSetInfoR17PucchClosedLoopIndexR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_PowerControlSetInfo_r17_Pucch_ClosedLoopIndex_r17_I0, PUCCH_PowerControlSetInfo_r17_Pucch_ClosedLoopIndex_r17_I1},
}

type PUCCH_PowerControlSetInfo_r17 struct {
	Pucch_PowerControlSetInfoId_r17  PUCCH_PowerControlSetInfoId_r17
	P0_PUCCH_Id_r17                  P0_PUCCH_Id
	Pucch_ClosedLoopIndex_r17        int64
	Pucch_PathlossReferenceRS_Id_r17 PUCCH_PathlossReferenceRS_Id_r17
}

func (ie *PUCCH_PowerControlSetInfo_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHPowerControlSetInfoR17Constraints)
	_ = seq

	// 1. pucch-PowerControlSetInfoId-r17: ref
	{
		if err := ie.Pucch_PowerControlSetInfoId_r17.Encode(e); err != nil {
			return err
		}
	}

	// 2. p0-PUCCH-Id-r17: ref
	{
		if err := ie.P0_PUCCH_Id_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. pucch-ClosedLoopIndex-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Pucch_ClosedLoopIndex_r17, pUCCHPowerControlSetInfoR17PucchClosedLoopIndexR17Constraints); err != nil {
			return err
		}
	}

	// 4. pucch-PathlossReferenceRS-Id-r17: ref
	{
		if err := ie.Pucch_PathlossReferenceRS_Id_r17.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PUCCH_PowerControlSetInfo_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHPowerControlSetInfoR17Constraints)
	_ = seq

	// 1. pucch-PowerControlSetInfoId-r17: ref
	{
		if err := ie.Pucch_PowerControlSetInfoId_r17.Decode(d); err != nil {
			return err
		}
	}

	// 2. p0-PUCCH-Id-r17: ref
	{
		if err := ie.P0_PUCCH_Id_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. pucch-ClosedLoopIndex-r17: enumerated
	{
		v2, err := d.DecodeEnumerated(pUCCHPowerControlSetInfoR17PucchClosedLoopIndexR17Constraints)
		if err != nil {
			return err
		}
		ie.Pucch_ClosedLoopIndex_r17 = v2
	}

	// 4. pucch-PathlossReferenceRS-Id-r17: ref
	{
		if err := ie.Pucch_PathlossReferenceRS_Id_r17.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
