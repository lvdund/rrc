// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: GapConfig (line 9162).

var gapConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "gapOffset"},
		{Name: "mgl"},
		{Name: "mgrp"},
		{Name: "mgta"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var gapConfigGapOffsetConstraints = per.Constrained(0, 159)

const (
	GapConfig_Mgl_Ms1dot5 = 0
	GapConfig_Mgl_Ms3     = 1
	GapConfig_Mgl_Ms3dot5 = 2
	GapConfig_Mgl_Ms4     = 3
	GapConfig_Mgl_Ms5dot5 = 4
	GapConfig_Mgl_Ms6     = 5
)

var gapConfigMglConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GapConfig_Mgl_Ms1dot5, GapConfig_Mgl_Ms3, GapConfig_Mgl_Ms3dot5, GapConfig_Mgl_Ms4, GapConfig_Mgl_Ms5dot5, GapConfig_Mgl_Ms6},
}

const (
	GapConfig_Mgrp_Ms20  = 0
	GapConfig_Mgrp_Ms40  = 1
	GapConfig_Mgrp_Ms80  = 2
	GapConfig_Mgrp_Ms160 = 3
)

var gapConfigMgrpConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GapConfig_Mgrp_Ms20, GapConfig_Mgrp_Ms40, GapConfig_Mgrp_Ms80, GapConfig_Mgrp_Ms160},
}

const (
	GapConfig_Mgta_Ms0      = 0
	GapConfig_Mgta_Ms0dot25 = 1
	GapConfig_Mgta_Ms0dot5  = 2
)

var gapConfigMgtaConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GapConfig_Mgta_Ms0, GapConfig_Mgta_Ms0dot25, GapConfig_Mgta_Ms0dot5},
}

const (
	GapConfig_Ext_RefServCellIndicator_PCell   = 0
	GapConfig_Ext_RefServCellIndicator_PSCell  = 1
	GapConfig_Ext_RefServCellIndicator_Mcg_FR2 = 2
)

var gapConfigExtRefServCellIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GapConfig_Ext_RefServCellIndicator_PCell, GapConfig_Ext_RefServCellIndicator_PSCell, GapConfig_Ext_RefServCellIndicator_Mcg_FR2},
}

const (
	GapConfig_Ext_Mgl_r16_Ms10 = 0
	GapConfig_Ext_Mgl_r16_Ms20 = 1
)

var gapConfigExtMglR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GapConfig_Ext_Mgl_r16_Ms10, GapConfig_Ext_Mgl_r16_Ms20},
}

type GapConfig struct {
	GapOffset                 int64
	Mgl                       int64
	Mgrp                      int64
	Mgta                      int64
	RefServCellIndicator      *int64
	RefFR2ServCellAsyncCA_r16 *ServCellIndex
	Mgl_r16                   *int64
}

func (ie *GapConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(gapConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.RefServCellIndicator != nil
	hasExtGroup1 := ie.RefFR2ServCellAsyncCA_r16 != nil || ie.Mgl_r16 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. gapOffset: integer
	{
		if err := e.EncodeInteger(ie.GapOffset, gapConfigGapOffsetConstraints); err != nil {
			return err
		}
	}

	// 3. mgl: enumerated
	{
		if err := e.EncodeEnumerated(ie.Mgl, gapConfigMglConstraints); err != nil {
			return err
		}
	}

	// 4. mgrp: enumerated
	{
		if err := e.EncodeEnumerated(ie.Mgrp, gapConfigMgrpConstraints); err != nil {
			return err
		}
	}

	// 5. mgta: enumerated
	{
		if err := e.EncodeEnumerated(ie.Mgta, gapConfigMgtaConstraints); err != nil {
			return err
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "refServCellIndicator", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.RefServCellIndicator != nil}); err != nil {
				return err
			}

			if ie.RefServCellIndicator != nil {
				if err := ex.EncodeEnumerated(*ie.RefServCellIndicator, gapConfigExtRefServCellIndicatorConstraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "refFR2ServCellAsyncCA-r16", Optional: true},
					{Name: "mgl-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.RefFR2ServCellAsyncCA_r16 != nil, ie.Mgl_r16 != nil}); err != nil {
				return err
			}

			if ie.RefFR2ServCellAsyncCA_r16 != nil {
				if err := ie.RefFR2ServCellAsyncCA_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Mgl_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Mgl_r16, gapConfigExtMglR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *GapConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(gapConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. gapOffset: integer
	{
		v0, err := d.DecodeInteger(gapConfigGapOffsetConstraints)
		if err != nil {
			return err
		}
		ie.GapOffset = v0
	}

	// 3. mgl: enumerated
	{
		v1, err := d.DecodeEnumerated(gapConfigMglConstraints)
		if err != nil {
			return err
		}
		ie.Mgl = v1
	}

	// 4. mgrp: enumerated
	{
		v2, err := d.DecodeEnumerated(gapConfigMgrpConstraints)
		if err != nil {
			return err
		}
		ie.Mgrp = v2
	}

	// 5. mgta: enumerated
	{
		v3, err := d.DecodeEnumerated(gapConfigMgtaConstraints)
		if err != nil {
			return err
		}
		ie.Mgta = v3
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "refServCellIndicator", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(gapConfigExtRefServCellIndicatorConstraints)
			if err != nil {
				return err
			}
			ie.RefServCellIndicator = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "refFR2ServCellAsyncCA-r16", Optional: true},
				{Name: "mgl-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.RefFR2ServCellAsyncCA_r16 = new(ServCellIndex)
			if err := ie.RefFR2ServCellAsyncCA_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(gapConfigExtMglR16Constraints)
			if err != nil {
				return err
			}
			ie.Mgl_r16 = &v
		}
	}

	return nil
}
