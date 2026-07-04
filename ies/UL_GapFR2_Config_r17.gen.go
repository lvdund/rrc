// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UL-GapFR2-Config-r17 (line 16264).

var uLGapFR2ConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "gapOffset-r17"},
		{Name: "ugl-r17"},
		{Name: "ugrp-r17"},
		{Name: "refFR2-ServCellAsyncCA-r17", Optional: true},
	},
}

var uLGapFR2ConfigR17GapOffsetR17Constraints = per.Constrained(0, 159)

const (
	UL_GapFR2_Config_r17_Ugl_r17_Ms0dot125 = 0
	UL_GapFR2_Config_r17_Ugl_r17_Ms0dot25  = 1
	UL_GapFR2_Config_r17_Ugl_r17_Ms0dot5   = 2
	UL_GapFR2_Config_r17_Ugl_r17_Ms1       = 3
)

var uLGapFR2ConfigR17UglR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UL_GapFR2_Config_r17_Ugl_r17_Ms0dot125, UL_GapFR2_Config_r17_Ugl_r17_Ms0dot25, UL_GapFR2_Config_r17_Ugl_r17_Ms0dot5, UL_GapFR2_Config_r17_Ugl_r17_Ms1},
}

const (
	UL_GapFR2_Config_r17_Ugrp_r17_Ms5   = 0
	UL_GapFR2_Config_r17_Ugrp_r17_Ms20  = 1
	UL_GapFR2_Config_r17_Ugrp_r17_Ms40  = 2
	UL_GapFR2_Config_r17_Ugrp_r17_Ms160 = 3
)

var uLGapFR2ConfigR17UgrpR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UL_GapFR2_Config_r17_Ugrp_r17_Ms5, UL_GapFR2_Config_r17_Ugrp_r17_Ms20, UL_GapFR2_Config_r17_Ugrp_r17_Ms40, UL_GapFR2_Config_r17_Ugrp_r17_Ms160},
}

type UL_GapFR2_Config_r17 struct {
	GapOffset_r17              int64
	Ugl_r17                    int64
	Ugrp_r17                   int64
	RefFR2_ServCellAsyncCA_r17 *ServCellIndex
}

func (ie *UL_GapFR2_Config_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLGapFR2ConfigR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RefFR2_ServCellAsyncCA_r17 != nil}); err != nil {
		return err
	}

	// 2. gapOffset-r17: integer
	{
		if err := e.EncodeInteger(ie.GapOffset_r17, uLGapFR2ConfigR17GapOffsetR17Constraints); err != nil {
			return err
		}
	}

	// 3. ugl-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Ugl_r17, uLGapFR2ConfigR17UglR17Constraints); err != nil {
			return err
		}
	}

	// 4. ugrp-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Ugrp_r17, uLGapFR2ConfigR17UgrpR17Constraints); err != nil {
			return err
		}
	}

	// 5. refFR2-ServCellAsyncCA-r17: ref
	{
		if ie.RefFR2_ServCellAsyncCA_r17 != nil {
			if err := ie.RefFR2_ServCellAsyncCA_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UL_GapFR2_Config_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLGapFR2ConfigR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. gapOffset-r17: integer
	{
		v0, err := d.DecodeInteger(uLGapFR2ConfigR17GapOffsetR17Constraints)
		if err != nil {
			return err
		}
		ie.GapOffset_r17 = v0
	}

	// 3. ugl-r17: enumerated
	{
		v1, err := d.DecodeEnumerated(uLGapFR2ConfigR17UglR17Constraints)
		if err != nil {
			return err
		}
		ie.Ugl_r17 = v1
	}

	// 4. ugrp-r17: enumerated
	{
		v2, err := d.DecodeEnumerated(uLGapFR2ConfigR17UgrpR17Constraints)
		if err != nil {
			return err
		}
		ie.Ugrp_r17 = v2
	}

	// 5. refFR2-ServCellAsyncCA-r17: ref
	{
		if seq.IsComponentPresent(3) {
			ie.RefFR2_ServCellAsyncCA_r17 = new(ServCellIndex)
			if err := ie.RefFR2_ServCellAsyncCA_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
