// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DummyA (line 19809).

var dummyAConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberNZP-CSI-RS-PerCC"},
		{Name: "maxNumberPortsAcrossNZP-CSI-RS-PerCC"},
		{Name: "maxNumberCS-IM-PerCC"},
		{Name: "maxNumberSimultaneousCSI-RS-ActBWP-AllCC"},
		{Name: "totalNumberPortsSimultaneousCSI-RS-ActBWP-AllCC"},
	},
}

var dummyAMaxNumberNZPCSIRSPerCCConstraints = per.Constrained(1, 32)

const (
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P2   = 0
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P4   = 1
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P8   = 2
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P12  = 3
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P16  = 4
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P24  = 5
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P32  = 6
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P40  = 7
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P48  = 8
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P56  = 9
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P64  = 10
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P72  = 11
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P80  = 12
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P88  = 13
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P96  = 14
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P104 = 15
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P112 = 16
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P120 = 17
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P128 = 18
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P136 = 19
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P144 = 20
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P152 = 21
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P160 = 22
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P168 = 23
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P176 = 24
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P184 = 25
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P192 = 26
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P200 = 27
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P208 = 28
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P216 = 29
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P224 = 30
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P232 = 31
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P240 = 32
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P248 = 33
	DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P256 = 34
)

var dummyAMaxNumberPortsAcrossNZPCSIRSPerCCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P2, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P4, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P8, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P12, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P16, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P24, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P32, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P40, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P48, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P56, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P64, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P72, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P80, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P88, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P96, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P104, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P112, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P120, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P128, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P136, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P144, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P152, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P160, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P168, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P176, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P184, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P192, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P200, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P208, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P216, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P224, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P232, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P240, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P248, DummyA_MaxNumberPortsAcrossNZP_CSI_RS_PerCC_P256},
}

const (
	DummyA_MaxNumberCS_IM_PerCC_N1  = 0
	DummyA_MaxNumberCS_IM_PerCC_N2  = 1
	DummyA_MaxNumberCS_IM_PerCC_N4  = 2
	DummyA_MaxNumberCS_IM_PerCC_N8  = 3
	DummyA_MaxNumberCS_IM_PerCC_N16 = 4
	DummyA_MaxNumberCS_IM_PerCC_N32 = 5
)

var dummyAMaxNumberCSIMPerCCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DummyA_MaxNumberCS_IM_PerCC_N1, DummyA_MaxNumberCS_IM_PerCC_N2, DummyA_MaxNumberCS_IM_PerCC_N4, DummyA_MaxNumberCS_IM_PerCC_N8, DummyA_MaxNumberCS_IM_PerCC_N16, DummyA_MaxNumberCS_IM_PerCC_N32},
}

const (
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N5  = 0
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N6  = 1
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N7  = 2
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N8  = 3
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N9  = 4
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N10 = 5
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N12 = 6
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N14 = 7
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N16 = 8
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N18 = 9
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N20 = 10
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N22 = 11
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N24 = 12
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N26 = 13
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N28 = 14
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N30 = 15
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N32 = 16
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N34 = 17
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N36 = 18
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N38 = 19
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N40 = 20
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N42 = 21
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N44 = 22
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N46 = 23
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N48 = 24
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N50 = 25
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N52 = 26
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N54 = 27
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N56 = 28
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N58 = 29
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N60 = 30
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N62 = 31
	DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N64 = 32
)

var dummyAMaxNumberSimultaneousCSIRSActBWPAllCCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N5, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N6, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N7, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N8, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N9, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N10, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N12, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N14, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N16, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N18, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N20, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N22, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N24, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N26, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N28, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N30, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N32, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N34, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N36, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N38, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N40, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N42, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N44, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N46, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N48, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N50, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N52, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N54, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N56, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N58, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N60, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N62, DummyA_MaxNumberSimultaneousCSI_RS_ActBWP_AllCC_N64},
}

const (
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P8   = 0
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P12  = 1
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P16  = 2
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P24  = 3
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P32  = 4
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P40  = 5
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P48  = 6
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P56  = 7
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P64  = 8
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P72  = 9
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P80  = 10
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P88  = 11
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P96  = 12
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P104 = 13
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P112 = 14
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P120 = 15
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P128 = 16
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P136 = 17
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P144 = 18
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P152 = 19
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P160 = 20
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P168 = 21
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P176 = 22
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P184 = 23
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P192 = 24
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P200 = 25
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P208 = 26
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P216 = 27
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P224 = 28
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P232 = 29
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P240 = 30
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P248 = 31
	DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P256 = 32
)

var dummyATotalNumberPortsSimultaneousCSIRSActBWPAllCCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P8, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P12, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P16, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P24, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P32, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P40, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P48, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P56, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P64, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P72, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P80, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P88, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P96, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P104, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P112, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P120, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P128, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P136, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P144, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P152, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P160, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P168, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P176, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P184, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P192, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P200, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P208, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P216, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P224, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P232, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P240, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P248, DummyA_TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_P256},
}

type DummyA struct {
	MaxNumberNZP_CSI_RS_PerCC                       int64
	MaxNumberPortsAcrossNZP_CSI_RS_PerCC            int64
	MaxNumberCS_IM_PerCC                            int64
	MaxNumberSimultaneousCSI_RS_ActBWP_AllCC        int64
	TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC int64
}

func (ie *DummyA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dummyAConstraints)
	_ = seq

	// 1. maxNumberNZP-CSI-RS-PerCC: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberNZP_CSI_RS_PerCC, dummyAMaxNumberNZPCSIRSPerCCConstraints); err != nil {
			return err
		}
	}

	// 2. maxNumberPortsAcrossNZP-CSI-RS-PerCC: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberPortsAcrossNZP_CSI_RS_PerCC, dummyAMaxNumberPortsAcrossNZPCSIRSPerCCConstraints); err != nil {
			return err
		}
	}

	// 3. maxNumberCS-IM-PerCC: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberCS_IM_PerCC, dummyAMaxNumberCSIMPerCCConstraints); err != nil {
			return err
		}
	}

	// 4. maxNumberSimultaneousCSI-RS-ActBWP-AllCC: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberSimultaneousCSI_RS_ActBWP_AllCC, dummyAMaxNumberSimultaneousCSIRSActBWPAllCCConstraints); err != nil {
			return err
		}
	}

	// 5. totalNumberPortsSimultaneousCSI-RS-ActBWP-AllCC: enumerated
	{
		if err := e.EncodeEnumerated(ie.TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC, dummyATotalNumberPortsSimultaneousCSIRSActBWPAllCCConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DummyA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dummyAConstraints)
	_ = seq

	// 1. maxNumberNZP-CSI-RS-PerCC: integer
	{
		v0, err := d.DecodeInteger(dummyAMaxNumberNZPCSIRSPerCCConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberNZP_CSI_RS_PerCC = v0
	}

	// 2. maxNumberPortsAcrossNZP-CSI-RS-PerCC: enumerated
	{
		v1, err := d.DecodeEnumerated(dummyAMaxNumberPortsAcrossNZPCSIRSPerCCConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberPortsAcrossNZP_CSI_RS_PerCC = v1
	}

	// 3. maxNumberCS-IM-PerCC: enumerated
	{
		v2, err := d.DecodeEnumerated(dummyAMaxNumberCSIMPerCCConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberCS_IM_PerCC = v2
	}

	// 4. maxNumberSimultaneousCSI-RS-ActBWP-AllCC: enumerated
	{
		v3, err := d.DecodeEnumerated(dummyAMaxNumberSimultaneousCSIRSActBWPAllCCConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberSimultaneousCSI_RS_ActBWP_AllCC = v3
	}

	// 5. totalNumberPortsSimultaneousCSI-RS-ActBWP-AllCC: enumerated
	{
		v4, err := d.DecodeEnumerated(dummyATotalNumberPortsSimultaneousCSIRSActBWPAllCCConstraints)
		if err != nil {
			return err
		}
		ie.TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC = v4
	}

	return nil
}
