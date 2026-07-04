// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PowerControl-r16 (line 28106).

var sLPowerControlR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-MaxTransPower-r16"},
		{Name: "sl-Alpha-PSSCH-PSCCH-r16", Optional: true},
		{Name: "dl-Alpha-PSSCH-PSCCH-r16", Optional: true},
		{Name: "sl-P0-PSSCH-PSCCH-r16", Optional: true},
		{Name: "dl-P0-PSSCH-PSCCH-r16", Optional: true},
		{Name: "dl-Alpha-PSFCH-r16", Optional: true},
		{Name: "dl-P0-PSFCH-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var sLPowerControlR16SlMaxTransPowerR16Constraints = per.Constrained(-30, 33)

const (
	SL_PowerControl_r16_Sl_Alpha_PSSCH_PSCCH_r16_Alpha0  = 0
	SL_PowerControl_r16_Sl_Alpha_PSSCH_PSCCH_r16_Alpha04 = 1
	SL_PowerControl_r16_Sl_Alpha_PSSCH_PSCCH_r16_Alpha05 = 2
	SL_PowerControl_r16_Sl_Alpha_PSSCH_PSCCH_r16_Alpha06 = 3
	SL_PowerControl_r16_Sl_Alpha_PSSCH_PSCCH_r16_Alpha07 = 4
	SL_PowerControl_r16_Sl_Alpha_PSSCH_PSCCH_r16_Alpha08 = 5
	SL_PowerControl_r16_Sl_Alpha_PSSCH_PSCCH_r16_Alpha09 = 6
	SL_PowerControl_r16_Sl_Alpha_PSSCH_PSCCH_r16_Alpha1  = 7
)

var sLPowerControlR16SlAlphaPSSCHPSCCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PowerControl_r16_Sl_Alpha_PSSCH_PSCCH_r16_Alpha0, SL_PowerControl_r16_Sl_Alpha_PSSCH_PSCCH_r16_Alpha04, SL_PowerControl_r16_Sl_Alpha_PSSCH_PSCCH_r16_Alpha05, SL_PowerControl_r16_Sl_Alpha_PSSCH_PSCCH_r16_Alpha06, SL_PowerControl_r16_Sl_Alpha_PSSCH_PSCCH_r16_Alpha07, SL_PowerControl_r16_Sl_Alpha_PSSCH_PSCCH_r16_Alpha08, SL_PowerControl_r16_Sl_Alpha_PSSCH_PSCCH_r16_Alpha09, SL_PowerControl_r16_Sl_Alpha_PSSCH_PSCCH_r16_Alpha1},
}

const (
	SL_PowerControl_r16_Dl_Alpha_PSSCH_PSCCH_r16_Alpha0  = 0
	SL_PowerControl_r16_Dl_Alpha_PSSCH_PSCCH_r16_Alpha04 = 1
	SL_PowerControl_r16_Dl_Alpha_PSSCH_PSCCH_r16_Alpha05 = 2
	SL_PowerControl_r16_Dl_Alpha_PSSCH_PSCCH_r16_Alpha06 = 3
	SL_PowerControl_r16_Dl_Alpha_PSSCH_PSCCH_r16_Alpha07 = 4
	SL_PowerControl_r16_Dl_Alpha_PSSCH_PSCCH_r16_Alpha08 = 5
	SL_PowerControl_r16_Dl_Alpha_PSSCH_PSCCH_r16_Alpha09 = 6
	SL_PowerControl_r16_Dl_Alpha_PSSCH_PSCCH_r16_Alpha1  = 7
)

var sLPowerControlR16DlAlphaPSSCHPSCCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PowerControl_r16_Dl_Alpha_PSSCH_PSCCH_r16_Alpha0, SL_PowerControl_r16_Dl_Alpha_PSSCH_PSCCH_r16_Alpha04, SL_PowerControl_r16_Dl_Alpha_PSSCH_PSCCH_r16_Alpha05, SL_PowerControl_r16_Dl_Alpha_PSSCH_PSCCH_r16_Alpha06, SL_PowerControl_r16_Dl_Alpha_PSSCH_PSCCH_r16_Alpha07, SL_PowerControl_r16_Dl_Alpha_PSSCH_PSCCH_r16_Alpha08, SL_PowerControl_r16_Dl_Alpha_PSSCH_PSCCH_r16_Alpha09, SL_PowerControl_r16_Dl_Alpha_PSSCH_PSCCH_r16_Alpha1},
}

var sLPowerControlR16SlP0PSSCHPSCCHR16Constraints = per.Constrained(-16, 15)

var sLPowerControlR16DlP0PSSCHPSCCHR16Constraints = per.Constrained(-16, 15)

const (
	SL_PowerControl_r16_Dl_Alpha_PSFCH_r16_Alpha0  = 0
	SL_PowerControl_r16_Dl_Alpha_PSFCH_r16_Alpha04 = 1
	SL_PowerControl_r16_Dl_Alpha_PSFCH_r16_Alpha05 = 2
	SL_PowerControl_r16_Dl_Alpha_PSFCH_r16_Alpha06 = 3
	SL_PowerControl_r16_Dl_Alpha_PSFCH_r16_Alpha07 = 4
	SL_PowerControl_r16_Dl_Alpha_PSFCH_r16_Alpha08 = 5
	SL_PowerControl_r16_Dl_Alpha_PSFCH_r16_Alpha09 = 6
	SL_PowerControl_r16_Dl_Alpha_PSFCH_r16_Alpha1  = 7
)

var sLPowerControlR16DlAlphaPSFCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PowerControl_r16_Dl_Alpha_PSFCH_r16_Alpha0, SL_PowerControl_r16_Dl_Alpha_PSFCH_r16_Alpha04, SL_PowerControl_r16_Dl_Alpha_PSFCH_r16_Alpha05, SL_PowerControl_r16_Dl_Alpha_PSFCH_r16_Alpha06, SL_PowerControl_r16_Dl_Alpha_PSFCH_r16_Alpha07, SL_PowerControl_r16_Dl_Alpha_PSFCH_r16_Alpha08, SL_PowerControl_r16_Dl_Alpha_PSFCH_r16_Alpha09, SL_PowerControl_r16_Dl_Alpha_PSFCH_r16_Alpha1},
}

var sLPowerControlR16DlP0PSFCHR16Constraints = per.Constrained(-16, 15)

var sLPowerControlR16DlP0PSSCHPSCCHR17Constraints = per.Constrained(-202, 24)

var sLPowerControlR16SlP0PSSCHPSCCHR17Constraints = per.Constrained(-202, 24)

var sLPowerControlR16DlP0PSFCHR17Constraints = per.Constrained(-202, 24)

type SL_PowerControl_r16 struct {
	Sl_MaxTransPower_r16     int64
	Sl_Alpha_PSSCH_PSCCH_r16 *int64
	Dl_Alpha_PSSCH_PSCCH_r16 *int64
	Sl_P0_PSSCH_PSCCH_r16    *int64
	Dl_P0_PSSCH_PSCCH_r16    *int64
	Dl_Alpha_PSFCH_r16       *int64
	Dl_P0_PSFCH_r16          *int64
	Dl_P0_PSSCH_PSCCH_r17    *int64
	Sl_P0_PSSCH_PSCCH_r17    *int64
	Dl_P0_PSFCH_r17          *int64
}

func (ie *SL_PowerControl_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPowerControlR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Dl_P0_PSSCH_PSCCH_r17 != nil || ie.Sl_P0_PSSCH_PSCCH_r17 != nil || ie.Dl_P0_PSFCH_r17 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_Alpha_PSSCH_PSCCH_r16 != nil, ie.Dl_Alpha_PSSCH_PSCCH_r16 != nil, ie.Sl_P0_PSSCH_PSCCH_r16 != nil, ie.Dl_P0_PSSCH_PSCCH_r16 != nil, ie.Dl_Alpha_PSFCH_r16 != nil, ie.Dl_P0_PSFCH_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-MaxTransPower-r16: integer
	{
		if err := e.EncodeInteger(ie.Sl_MaxTransPower_r16, sLPowerControlR16SlMaxTransPowerR16Constraints); err != nil {
			return err
		}
	}

	// 4. sl-Alpha-PSSCH-PSCCH-r16: enumerated
	{
		if ie.Sl_Alpha_PSSCH_PSCCH_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_Alpha_PSSCH_PSCCH_r16, sLPowerControlR16SlAlphaPSSCHPSCCHR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. dl-Alpha-PSSCH-PSCCH-r16: enumerated
	{
		if ie.Dl_Alpha_PSSCH_PSCCH_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Dl_Alpha_PSSCH_PSCCH_r16, sLPowerControlR16DlAlphaPSSCHPSCCHR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sl-P0-PSSCH-PSCCH-r16: integer
	{
		if ie.Sl_P0_PSSCH_PSCCH_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_P0_PSSCH_PSCCH_r16, sLPowerControlR16SlP0PSSCHPSCCHR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. dl-P0-PSSCH-PSCCH-r16: integer
	{
		if ie.Dl_P0_PSSCH_PSCCH_r16 != nil {
			if err := e.EncodeInteger(*ie.Dl_P0_PSSCH_PSCCH_r16, sLPowerControlR16DlP0PSSCHPSCCHR16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. dl-Alpha-PSFCH-r16: enumerated
	{
		if ie.Dl_Alpha_PSFCH_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Dl_Alpha_PSFCH_r16, sLPowerControlR16DlAlphaPSFCHR16Constraints); err != nil {
				return err
			}
		}
	}

	// 9. dl-P0-PSFCH-r16: integer
	{
		if ie.Dl_P0_PSFCH_r16 != nil {
			if err := e.EncodeInteger(*ie.Dl_P0_PSFCH_r16, sLPowerControlR16DlP0PSFCHR16Constraints); err != nil {
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
					{Name: "dl-P0-PSSCH-PSCCH-r17", Optional: true},
					{Name: "sl-P0-PSSCH-PSCCH-r17", Optional: true},
					{Name: "dl-P0-PSFCH-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Dl_P0_PSSCH_PSCCH_r17 != nil, ie.Sl_P0_PSSCH_PSCCH_r17 != nil, ie.Dl_P0_PSFCH_r17 != nil}); err != nil {
				return err
			}

			if ie.Dl_P0_PSSCH_PSCCH_r17 != nil {
				if err := ex.EncodeInteger(*ie.Dl_P0_PSSCH_PSCCH_r17, sLPowerControlR16DlP0PSSCHPSCCHR17Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_P0_PSSCH_PSCCH_r17 != nil {
				if err := ex.EncodeInteger(*ie.Sl_P0_PSSCH_PSCCH_r17, sLPowerControlR16SlP0PSSCHPSCCHR17Constraints); err != nil {
					return err
				}
			}

			if ie.Dl_P0_PSFCH_r17 != nil {
				if err := ex.EncodeInteger(*ie.Dl_P0_PSFCH_r17, sLPowerControlR16DlP0PSFCHR17Constraints); err != nil {
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

func (ie *SL_PowerControl_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPowerControlR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-MaxTransPower-r16: integer
	{
		v0, err := d.DecodeInteger(sLPowerControlR16SlMaxTransPowerR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_MaxTransPower_r16 = v0
	}

	// 4. sl-Alpha-PSSCH-PSCCH-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sLPowerControlR16SlAlphaPSSCHPSCCHR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_Alpha_PSSCH_PSCCH_r16 = &idx
		}
	}

	// 5. dl-Alpha-PSSCH-PSCCH-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sLPowerControlR16DlAlphaPSSCHPSCCHR16Constraints)
			if err != nil {
				return err
			}
			ie.Dl_Alpha_PSSCH_PSCCH_r16 = &idx
		}
	}

	// 6. sl-P0-PSSCH-PSCCH-r16: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(sLPowerControlR16SlP0PSSCHPSCCHR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_P0_PSSCH_PSCCH_r16 = &v
		}
	}

	// 7. dl-P0-PSSCH-PSCCH-r16: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(sLPowerControlR16DlP0PSSCHPSCCHR16Constraints)
			if err != nil {
				return err
			}
			ie.Dl_P0_PSSCH_PSCCH_r16 = &v
		}
	}

	// 8. dl-Alpha-PSFCH-r16: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(sLPowerControlR16DlAlphaPSFCHR16Constraints)
			if err != nil {
				return err
			}
			ie.Dl_Alpha_PSFCH_r16 = &idx
		}
	}

	// 9. dl-P0-PSFCH-r16: integer
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeInteger(sLPowerControlR16DlP0PSFCHR16Constraints)
			if err != nil {
				return err
			}
			ie.Dl_P0_PSFCH_r16 = &v
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
				{Name: "dl-P0-PSSCH-PSCCH-r17", Optional: true},
				{Name: "sl-P0-PSSCH-PSCCH-r17", Optional: true},
				{Name: "dl-P0-PSFCH-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(sLPowerControlR16DlP0PSSCHPSCCHR17Constraints)
			if err != nil {
				return err
			}
			ie.Dl_P0_PSSCH_PSCCH_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(sLPowerControlR16SlP0PSSCHPSCCHR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_P0_PSSCH_PSCCH_r17 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeInteger(sLPowerControlR16DlP0PSFCHR17Constraints)
			if err != nil {
				return err
			}
			ie.Dl_P0_PSFCH_r17 = &v
		}
	}

	return nil
}
