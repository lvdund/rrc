// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MIB (line 798).

var mIBConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "systemFrameNumber"},
		{Name: "subCarrierSpacingCommon"},
		{Name: "ssb-SubcarrierOffset"},
		{Name: "dmrs-TypeA-Position"},
		{Name: "pdcch-ConfigSIB1"},
		{Name: "cellBarred"},
		{Name: "intraFreqReselection"},
		{Name: "spare"},
	},
}

var mIBSystemFrameNumberConstraints = per.FixedSize(6)

const (
	MIB_SubCarrierSpacingCommon_Scs15or60  = 0
	MIB_SubCarrierSpacingCommon_Scs30or120 = 1
)

var mIBSubCarrierSpacingCommonConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIB_SubCarrierSpacingCommon_Scs15or60, MIB_SubCarrierSpacingCommon_Scs30or120},
}

var mIBSsbSubcarrierOffsetConstraints = per.Constrained(0, 15)

const (
	MIB_Dmrs_TypeA_Position_Pos2 = 0
	MIB_Dmrs_TypeA_Position_Pos3 = 1
)

var mIBDmrsTypeAPositionConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIB_Dmrs_TypeA_Position_Pos2, MIB_Dmrs_TypeA_Position_Pos3},
}

const (
	MIB_CellBarred_Barred    = 0
	MIB_CellBarred_NotBarred = 1
)

var mIBCellBarredConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIB_CellBarred_Barred, MIB_CellBarred_NotBarred},
}

const (
	MIB_IntraFreqReselection_Allowed    = 0
	MIB_IntraFreqReselection_NotAllowed = 1
)

var mIBIntraFreqReselectionConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIB_IntraFreqReselection_Allowed, MIB_IntraFreqReselection_NotAllowed},
}

var mIBSpareConstraints = per.FixedSize(1)

type MIB struct {
	SystemFrameNumber       per.BitString
	SubCarrierSpacingCommon int64
	Ssb_SubcarrierOffset    int64
	Dmrs_TypeA_Position     int64
	Pdcch_ConfigSIB1        PDCCH_ConfigSIB1
	CellBarred              int64
	IntraFreqReselection    int64
	Spare                   per.BitString
}

func (ie *MIB) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mIBConstraints)
	_ = seq

	// 1. systemFrameNumber: bit-string
	{
		if err := e.EncodeBitString(ie.SystemFrameNumber, mIBSystemFrameNumberConstraints); err != nil {
			return err
		}
	}

	// 2. subCarrierSpacingCommon: enumerated
	{
		if err := e.EncodeEnumerated(ie.SubCarrierSpacingCommon, mIBSubCarrierSpacingCommonConstraints); err != nil {
			return err
		}
	}

	// 3. ssb-SubcarrierOffset: integer
	{
		if err := e.EncodeInteger(ie.Ssb_SubcarrierOffset, mIBSsbSubcarrierOffsetConstraints); err != nil {
			return err
		}
	}

	// 4. dmrs-TypeA-Position: enumerated
	{
		if err := e.EncodeEnumerated(ie.Dmrs_TypeA_Position, mIBDmrsTypeAPositionConstraints); err != nil {
			return err
		}
	}

	// 5. pdcch-ConfigSIB1: ref
	{
		if err := ie.Pdcch_ConfigSIB1.Encode(e); err != nil {
			return err
		}
	}

	// 6. cellBarred: enumerated
	{
		if err := e.EncodeEnumerated(ie.CellBarred, mIBCellBarredConstraints); err != nil {
			return err
		}
	}

	// 7. intraFreqReselection: enumerated
	{
		if err := e.EncodeEnumerated(ie.IntraFreqReselection, mIBIntraFreqReselectionConstraints); err != nil {
			return err
		}
	}

	// 8. spare: bit-string
	{
		if err := e.EncodeBitString(ie.Spare, mIBSpareConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MIB) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mIBConstraints)
	_ = seq

	// 1. systemFrameNumber: bit-string
	{
		v0, err := d.DecodeBitString(mIBSystemFrameNumberConstraints)
		if err != nil {
			return err
		}
		ie.SystemFrameNumber = v0
	}

	// 2. subCarrierSpacingCommon: enumerated
	{
		v1, err := d.DecodeEnumerated(mIBSubCarrierSpacingCommonConstraints)
		if err != nil {
			return err
		}
		ie.SubCarrierSpacingCommon = v1
	}

	// 3. ssb-SubcarrierOffset: integer
	{
		v2, err := d.DecodeInteger(mIBSsbSubcarrierOffsetConstraints)
		if err != nil {
			return err
		}
		ie.Ssb_SubcarrierOffset = v2
	}

	// 4. dmrs-TypeA-Position: enumerated
	{
		v3, err := d.DecodeEnumerated(mIBDmrsTypeAPositionConstraints)
		if err != nil {
			return err
		}
		ie.Dmrs_TypeA_Position = v3
	}

	// 5. pdcch-ConfigSIB1: ref
	{
		if err := ie.Pdcch_ConfigSIB1.Decode(d); err != nil {
			return err
		}
	}

	// 6. cellBarred: enumerated
	{
		v5, err := d.DecodeEnumerated(mIBCellBarredConstraints)
		if err != nil {
			return err
		}
		ie.CellBarred = v5
	}

	// 7. intraFreqReselection: enumerated
	{
		v6, err := d.DecodeEnumerated(mIBIntraFreqReselectionConstraints)
		if err != nil {
			return err
		}
		ie.IntraFreqReselection = v6
	}

	// 8. spare: bit-string
	{
		v7, err := d.DecodeBitString(mIBSpareConstraints)
		if err != nil {
			return err
		}
		ie.Spare = v7
	}

	return nil
}
