// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PRS-PowerControl-r18 (line 27672).

var sLPRSPowerControlR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-P0-SL-PRS-r18", Optional: true},
		{Name: "dl-Alpha-SL-PRS-r18", Optional: true},
		{Name: "sl-P0-SL-PRS-r18", Optional: true},
		{Name: "sl-Alpha-SL-PRS-r18", Optional: true},
	},
}

var sLPRSPowerControlR18DlP0SLPRSR18Constraints = per.Constrained(-202, 24)

const (
	SL_PRS_PowerControl_r18_Dl_Alpha_SL_PRS_r18_Alpha0  = 0
	SL_PRS_PowerControl_r18_Dl_Alpha_SL_PRS_r18_Alpha04 = 1
	SL_PRS_PowerControl_r18_Dl_Alpha_SL_PRS_r18_Alpha05 = 2
	SL_PRS_PowerControl_r18_Dl_Alpha_SL_PRS_r18_Alpha06 = 3
	SL_PRS_PowerControl_r18_Dl_Alpha_SL_PRS_r18_Alpha07 = 4
	SL_PRS_PowerControl_r18_Dl_Alpha_SL_PRS_r18_Alpha08 = 5
	SL_PRS_PowerControl_r18_Dl_Alpha_SL_PRS_r18_Alpha09 = 6
	SL_PRS_PowerControl_r18_Dl_Alpha_SL_PRS_r18_Alpha1  = 7
)

var sLPRSPowerControlR18DlAlphaSLPRSR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PRS_PowerControl_r18_Dl_Alpha_SL_PRS_r18_Alpha0, SL_PRS_PowerControl_r18_Dl_Alpha_SL_PRS_r18_Alpha04, SL_PRS_PowerControl_r18_Dl_Alpha_SL_PRS_r18_Alpha05, SL_PRS_PowerControl_r18_Dl_Alpha_SL_PRS_r18_Alpha06, SL_PRS_PowerControl_r18_Dl_Alpha_SL_PRS_r18_Alpha07, SL_PRS_PowerControl_r18_Dl_Alpha_SL_PRS_r18_Alpha08, SL_PRS_PowerControl_r18_Dl_Alpha_SL_PRS_r18_Alpha09, SL_PRS_PowerControl_r18_Dl_Alpha_SL_PRS_r18_Alpha1},
}

var sLPRSPowerControlR18SlP0SLPRSR18Constraints = per.Constrained(-202, 24)

const (
	SL_PRS_PowerControl_r18_Sl_Alpha_SL_PRS_r18_Alpha0  = 0
	SL_PRS_PowerControl_r18_Sl_Alpha_SL_PRS_r18_Alpha04 = 1
	SL_PRS_PowerControl_r18_Sl_Alpha_SL_PRS_r18_Alpha05 = 2
	SL_PRS_PowerControl_r18_Sl_Alpha_SL_PRS_r18_Alpha06 = 3
	SL_PRS_PowerControl_r18_Sl_Alpha_SL_PRS_r18_Alpha07 = 4
	SL_PRS_PowerControl_r18_Sl_Alpha_SL_PRS_r18_Alpha08 = 5
	SL_PRS_PowerControl_r18_Sl_Alpha_SL_PRS_r18_Alpha09 = 6
	SL_PRS_PowerControl_r18_Sl_Alpha_SL_PRS_r18_Alpha1  = 7
)

var sLPRSPowerControlR18SlAlphaSLPRSR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PRS_PowerControl_r18_Sl_Alpha_SL_PRS_r18_Alpha0, SL_PRS_PowerControl_r18_Sl_Alpha_SL_PRS_r18_Alpha04, SL_PRS_PowerControl_r18_Sl_Alpha_SL_PRS_r18_Alpha05, SL_PRS_PowerControl_r18_Sl_Alpha_SL_PRS_r18_Alpha06, SL_PRS_PowerControl_r18_Sl_Alpha_SL_PRS_r18_Alpha07, SL_PRS_PowerControl_r18_Sl_Alpha_SL_PRS_r18_Alpha08, SL_PRS_PowerControl_r18_Sl_Alpha_SL_PRS_r18_Alpha09, SL_PRS_PowerControl_r18_Sl_Alpha_SL_PRS_r18_Alpha1},
}

type SL_PRS_PowerControl_r18 struct {
	Dl_P0_SL_PRS_r18    *int64
	Dl_Alpha_SL_PRS_r18 *int64
	Sl_P0_SL_PRS_r18    *int64
	Sl_Alpha_SL_PRS_r18 *int64
}

func (ie *SL_PRS_PowerControl_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPRSPowerControlR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Dl_P0_SL_PRS_r18 != nil, ie.Dl_Alpha_SL_PRS_r18 != nil, ie.Sl_P0_SL_PRS_r18 != nil, ie.Sl_Alpha_SL_PRS_r18 != nil}); err != nil {
		return err
	}

	// 2. dl-P0-SL-PRS-r18: integer
	{
		if ie.Dl_P0_SL_PRS_r18 != nil {
			if err := e.EncodeInteger(*ie.Dl_P0_SL_PRS_r18, sLPRSPowerControlR18DlP0SLPRSR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. dl-Alpha-SL-PRS-r18: enumerated
	{
		if ie.Dl_Alpha_SL_PRS_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Dl_Alpha_SL_PRS_r18, sLPRSPowerControlR18DlAlphaSLPRSR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-P0-SL-PRS-r18: integer
	{
		if ie.Sl_P0_SL_PRS_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_P0_SL_PRS_r18, sLPRSPowerControlR18SlP0SLPRSR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-Alpha-SL-PRS-r18: enumerated
	{
		if ie.Sl_Alpha_SL_PRS_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_Alpha_SL_PRS_r18, sLPRSPowerControlR18SlAlphaSLPRSR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_PRS_PowerControl_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPRSPowerControlR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dl-P0-SL-PRS-r18: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(sLPRSPowerControlR18DlP0SLPRSR18Constraints)
			if err != nil {
				return err
			}
			ie.Dl_P0_SL_PRS_r18 = &v
		}
	}

	// 3. dl-Alpha-SL-PRS-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sLPRSPowerControlR18DlAlphaSLPRSR18Constraints)
			if err != nil {
				return err
			}
			ie.Dl_Alpha_SL_PRS_r18 = &idx
		}
	}

	// 4. sl-P0-SL-PRS-r18: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(sLPRSPowerControlR18SlP0SLPRSR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_P0_SL_PRS_r18 = &v
		}
	}

	// 5. sl-Alpha-SL-PRS-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(sLPRSPowerControlR18SlAlphaSLPRSR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_Alpha_SL_PRS_r18 = &idx
		}
	}

	return nil
}
