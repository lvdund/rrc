// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PSBCH-Config-r16 (line 27704).

var sLPSBCHConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-P0-PSBCH-r16", Optional: true},
		{Name: "dl-Alpha-PSBCH-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var sLPSBCHConfigR16DlP0PSBCHR16Constraints = per.Constrained(-16, 15)

const (
	SL_PSBCH_Config_r16_Dl_Alpha_PSBCH_r16_Alpha0  = 0
	SL_PSBCH_Config_r16_Dl_Alpha_PSBCH_r16_Alpha04 = 1
	SL_PSBCH_Config_r16_Dl_Alpha_PSBCH_r16_Alpha05 = 2
	SL_PSBCH_Config_r16_Dl_Alpha_PSBCH_r16_Alpha06 = 3
	SL_PSBCH_Config_r16_Dl_Alpha_PSBCH_r16_Alpha07 = 4
	SL_PSBCH_Config_r16_Dl_Alpha_PSBCH_r16_Alpha08 = 5
	SL_PSBCH_Config_r16_Dl_Alpha_PSBCH_r16_Alpha09 = 6
	SL_PSBCH_Config_r16_Dl_Alpha_PSBCH_r16_Alpha1  = 7
)

var sLPSBCHConfigR16DlAlphaPSBCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PSBCH_Config_r16_Dl_Alpha_PSBCH_r16_Alpha0, SL_PSBCH_Config_r16_Dl_Alpha_PSBCH_r16_Alpha04, SL_PSBCH_Config_r16_Dl_Alpha_PSBCH_r16_Alpha05, SL_PSBCH_Config_r16_Dl_Alpha_PSBCH_r16_Alpha06, SL_PSBCH_Config_r16_Dl_Alpha_PSBCH_r16_Alpha07, SL_PSBCH_Config_r16_Dl_Alpha_PSBCH_r16_Alpha08, SL_PSBCH_Config_r16_Dl_Alpha_PSBCH_r16_Alpha09, SL_PSBCH_Config_r16_Dl_Alpha_PSBCH_r16_Alpha1},
}

var sLPSBCHConfigR16DlP0PSBCHR17Constraints = per.Constrained(-202, 24)

type SL_PSBCH_Config_r16 struct {
	Dl_P0_PSBCH_r16    *int64
	Dl_Alpha_PSBCH_r16 *int64
	Dl_P0_PSBCH_r17    *int64
}

func (ie *SL_PSBCH_Config_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPSBCHConfigR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Dl_P0_PSBCH_r17 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Dl_P0_PSBCH_r16 != nil, ie.Dl_Alpha_PSBCH_r16 != nil}); err != nil {
		return err
	}

	// 3. dl-P0-PSBCH-r16: integer
	{
		if ie.Dl_P0_PSBCH_r16 != nil {
			if err := e.EncodeInteger(*ie.Dl_P0_PSBCH_r16, sLPSBCHConfigR16DlP0PSBCHR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. dl-Alpha-PSBCH-r16: enumerated
	{
		if ie.Dl_Alpha_PSBCH_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Dl_Alpha_PSBCH_r16, sLPSBCHConfigR16DlAlphaPSBCHR16Constraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "dl-P0-PSBCH-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Dl_P0_PSBCH_r17 != nil}); err != nil {
				return err
			}

			if ie.Dl_P0_PSBCH_r17 != nil {
				if err := ex.EncodeInteger(*ie.Dl_P0_PSBCH_r17, sLPSBCHConfigR16DlP0PSBCHR17Constraints); err != nil {
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

func (ie *SL_PSBCH_Config_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPSBCHConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. dl-P0-PSBCH-r16: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(sLPSBCHConfigR16DlP0PSBCHR16Constraints)
			if err != nil {
				return err
			}
			ie.Dl_P0_PSBCH_r16 = &v
		}
	}

	// 4. dl-Alpha-PSBCH-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sLPSBCHConfigR16DlAlphaPSBCHR16Constraints)
			if err != nil {
				return err
			}
			ie.Dl_Alpha_PSBCH_r16 = &idx
		}
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
				{Name: "dl-P0-PSBCH-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(sLPSBCHConfigR16DlP0PSBCHR17Constraints)
			if err != nil {
				return err
			}
			ie.Dl_P0_PSBCH_r17 = &v
		}
	}

	return nil
}
