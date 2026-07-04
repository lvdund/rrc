// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PosGapConfig-r17 (line 9198).

var posGapConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "measPosPreConfigGapId-r17"},
		{Name: "gapOffset-r17"},
		{Name: "mgl-r17"},
		{Name: "mgrp-r17"},
		{Name: "mgta-r17"},
		{Name: "gapType-r17"},
	},
}

var posGapConfigR17GapOffsetR17Constraints = per.Constrained(0, 159)

const (
	PosGapConfig_r17_Mgl_r17_Ms1dot5 = 0
	PosGapConfig_r17_Mgl_r17_Ms3     = 1
	PosGapConfig_r17_Mgl_r17_Ms3dot5 = 2
	PosGapConfig_r17_Mgl_r17_Ms4     = 3
	PosGapConfig_r17_Mgl_r17_Ms5dot5 = 4
	PosGapConfig_r17_Mgl_r17_Ms6     = 5
	PosGapConfig_r17_Mgl_r17_Ms10    = 6
	PosGapConfig_r17_Mgl_r17_Ms20    = 7
)

var posGapConfigR17MglR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosGapConfig_r17_Mgl_r17_Ms1dot5, PosGapConfig_r17_Mgl_r17_Ms3, PosGapConfig_r17_Mgl_r17_Ms3dot5, PosGapConfig_r17_Mgl_r17_Ms4, PosGapConfig_r17_Mgl_r17_Ms5dot5, PosGapConfig_r17_Mgl_r17_Ms6, PosGapConfig_r17_Mgl_r17_Ms10, PosGapConfig_r17_Mgl_r17_Ms20},
}

const (
	PosGapConfig_r17_Mgrp_r17_Ms20  = 0
	PosGapConfig_r17_Mgrp_r17_Ms40  = 1
	PosGapConfig_r17_Mgrp_r17_Ms80  = 2
	PosGapConfig_r17_Mgrp_r17_Ms160 = 3
)

var posGapConfigR17MgrpR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosGapConfig_r17_Mgrp_r17_Ms20, PosGapConfig_r17_Mgrp_r17_Ms40, PosGapConfig_r17_Mgrp_r17_Ms80, PosGapConfig_r17_Mgrp_r17_Ms160},
}

const (
	PosGapConfig_r17_Mgta_r17_Ms0      = 0
	PosGapConfig_r17_Mgta_r17_Ms0dot25 = 1
	PosGapConfig_r17_Mgta_r17_Ms0dot5  = 2
)

var posGapConfigR17MgtaR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosGapConfig_r17_Mgta_r17_Ms0, PosGapConfig_r17_Mgta_r17_Ms0dot25, PosGapConfig_r17_Mgta_r17_Ms0dot5},
}

const (
	PosGapConfig_r17_GapType_r17_PerUE  = 0
	PosGapConfig_r17_GapType_r17_PerFR1 = 1
	PosGapConfig_r17_GapType_r17_PerFR2 = 2
)

var posGapConfigR17GapTypeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosGapConfig_r17_GapType_r17_PerUE, PosGapConfig_r17_GapType_r17_PerFR1, PosGapConfig_r17_GapType_r17_PerFR2},
}

type PosGapConfig_r17 struct {
	MeasPosPreConfigGapId_r17 MeasPosPreConfigGapId_r17
	GapOffset_r17             int64
	Mgl_r17                   int64
	Mgrp_r17                  int64
	Mgta_r17                  int64
	GapType_r17               int64
}

func (ie *PosGapConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(posGapConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. measPosPreConfigGapId-r17: ref
	{
		if err := ie.MeasPosPreConfigGapId_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. gapOffset-r17: integer
	{
		if err := e.EncodeInteger(ie.GapOffset_r17, posGapConfigR17GapOffsetR17Constraints); err != nil {
			return err
		}
	}

	// 4. mgl-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Mgl_r17, posGapConfigR17MglR17Constraints); err != nil {
			return err
		}
	}

	// 5. mgrp-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Mgrp_r17, posGapConfigR17MgrpR17Constraints); err != nil {
			return err
		}
	}

	// 6. mgta-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Mgta_r17, posGapConfigR17MgtaR17Constraints); err != nil {
			return err
		}
	}

	// 7. gapType-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.GapType_r17, posGapConfigR17GapTypeR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PosGapConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(posGapConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. measPosPreConfigGapId-r17: ref
	{
		if err := ie.MeasPosPreConfigGapId_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. gapOffset-r17: integer
	{
		v1, err := d.DecodeInteger(posGapConfigR17GapOffsetR17Constraints)
		if err != nil {
			return err
		}
		ie.GapOffset_r17 = v1
	}

	// 4. mgl-r17: enumerated
	{
		v2, err := d.DecodeEnumerated(posGapConfigR17MglR17Constraints)
		if err != nil {
			return err
		}
		ie.Mgl_r17 = v2
	}

	// 5. mgrp-r17: enumerated
	{
		v3, err := d.DecodeEnumerated(posGapConfigR17MgrpR17Constraints)
		if err != nil {
			return err
		}
		ie.Mgrp_r17 = v3
	}

	// 6. mgta-r17: enumerated
	{
		v4, err := d.DecodeEnumerated(posGapConfigR17MgtaR17Constraints)
		if err != nil {
			return err
		}
		ie.Mgta_r17 = v4
	}

	// 7. gapType-r17: enumerated
	{
		v5, err := d.DecodeEnumerated(posGapConfigR17GapTypeR17Constraints)
		if err != nil {
			return err
		}
		ie.GapType_r17 = v5
	}

	return nil
}
