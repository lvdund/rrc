// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-BWP-Generic-r16 (line 26761).

var sLBWPGenericR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-BWP-r16", Optional: true},
		{Name: "sl-LengthSymbols-r16", Optional: true},
		{Name: "sl-StartSymbol-r16", Optional: true},
		{Name: "sl-PSBCH-Config-r16", Optional: true},
		{Name: "sl-TxDirectCurrentLocation-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	SL_BWP_Generic_r16_Sl_LengthSymbols_r16_Sym7  = 0
	SL_BWP_Generic_r16_Sl_LengthSymbols_r16_Sym8  = 1
	SL_BWP_Generic_r16_Sl_LengthSymbols_r16_Sym9  = 2
	SL_BWP_Generic_r16_Sl_LengthSymbols_r16_Sym10 = 3
	SL_BWP_Generic_r16_Sl_LengthSymbols_r16_Sym11 = 4
	SL_BWP_Generic_r16_Sl_LengthSymbols_r16_Sym12 = 5
	SL_BWP_Generic_r16_Sl_LengthSymbols_r16_Sym13 = 6
	SL_BWP_Generic_r16_Sl_LengthSymbols_r16_Sym14 = 7
)

var sLBWPGenericR16SlLengthSymbolsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_BWP_Generic_r16_Sl_LengthSymbols_r16_Sym7, SL_BWP_Generic_r16_Sl_LengthSymbols_r16_Sym8, SL_BWP_Generic_r16_Sl_LengthSymbols_r16_Sym9, SL_BWP_Generic_r16_Sl_LengthSymbols_r16_Sym10, SL_BWP_Generic_r16_Sl_LengthSymbols_r16_Sym11, SL_BWP_Generic_r16_Sl_LengthSymbols_r16_Sym12, SL_BWP_Generic_r16_Sl_LengthSymbols_r16_Sym13, SL_BWP_Generic_r16_Sl_LengthSymbols_r16_Sym14},
}

const (
	SL_BWP_Generic_r16_Sl_StartSymbol_r16_Sym0 = 0
	SL_BWP_Generic_r16_Sl_StartSymbol_r16_Sym1 = 1
	SL_BWP_Generic_r16_Sl_StartSymbol_r16_Sym2 = 2
	SL_BWP_Generic_r16_Sl_StartSymbol_r16_Sym3 = 3
	SL_BWP_Generic_r16_Sl_StartSymbol_r16_Sym4 = 4
	SL_BWP_Generic_r16_Sl_StartSymbol_r16_Sym5 = 5
	SL_BWP_Generic_r16_Sl_StartSymbol_r16_Sym6 = 6
	SL_BWP_Generic_r16_Sl_StartSymbol_r16_Sym7 = 7
)

var sLBWPGenericR16SlStartSymbolR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_BWP_Generic_r16_Sl_StartSymbol_r16_Sym0, SL_BWP_Generic_r16_Sl_StartSymbol_r16_Sym1, SL_BWP_Generic_r16_Sl_StartSymbol_r16_Sym2, SL_BWP_Generic_r16_Sl_StartSymbol_r16_Sym3, SL_BWP_Generic_r16_Sl_StartSymbol_r16_Sym4, SL_BWP_Generic_r16_Sl_StartSymbol_r16_Sym5, SL_BWP_Generic_r16_Sl_StartSymbol_r16_Sym6, SL_BWP_Generic_r16_Sl_StartSymbol_r16_Sym7},
}

var sL_BWP_Generic_r16SlPSBCHConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_BWP_Generic_r16_Sl_PSBCH_Config_r16_Release = 0
	SL_BWP_Generic_r16_Sl_PSBCH_Config_r16_Setup   = 1
)

var sLBWPGenericR16SlTxDirectCurrentLocationR16Constraints = per.Constrained(0, 3301)

var sLBWPGenericR16ExtSlUnlicensedR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_BWP_Generic_r16_Ext_Sl_Unlicensed_r18_Release = 0
	SL_BWP_Generic_r16_Ext_Sl_Unlicensed_r18_Setup   = 1
)

type SL_BWP_Generic_r16 struct {
	Sl_BWP_r16           *BWP
	Sl_LengthSymbols_r16 *int64
	Sl_StartSymbol_r16   *int64
	Sl_PSBCH_Config_r16  *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_PSBCH_Config_r16
	}
	Sl_TxDirectCurrentLocation_r16 *int64
	Sl_Unlicensed_r18              *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_Unlicensed_r18
	}
}

func (ie *SL_BWP_Generic_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLBWPGenericR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_Unlicensed_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_BWP_r16 != nil, ie.Sl_LengthSymbols_r16 != nil, ie.Sl_StartSymbol_r16 != nil, ie.Sl_PSBCH_Config_r16 != nil, ie.Sl_TxDirectCurrentLocation_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-BWP-r16: ref
	{
		if ie.Sl_BWP_r16 != nil {
			if err := ie.Sl_BWP_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-LengthSymbols-r16: enumerated
	{
		if ie.Sl_LengthSymbols_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_LengthSymbols_r16, sLBWPGenericR16SlLengthSymbolsR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-StartSymbol-r16: enumerated
	{
		if ie.Sl_StartSymbol_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_StartSymbol_r16, sLBWPGenericR16SlStartSymbolR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sl-PSBCH-Config-r16: choice
	{
		if ie.Sl_PSBCH_Config_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(sL_BWP_Generic_r16SlPSBCHConfigR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_PSBCH_Config_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_PSBCH_Config_r16).Choice {
			case SL_BWP_Generic_r16_Sl_PSBCH_Config_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SL_BWP_Generic_r16_Sl_PSBCH_Config_r16_Setup:
				if err := (*ie.Sl_PSBCH_Config_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_PSBCH_Config_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 7. sl-TxDirectCurrentLocation-r16: integer
	{
		if ie.Sl_TxDirectCurrentLocation_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_TxDirectCurrentLocation_r16, sLBWPGenericR16SlTxDirectCurrentLocationR16Constraints); err != nil {
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
					{Name: "sl-Unlicensed-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_Unlicensed_r18 != nil}); err != nil {
				return err
			}

			if ie.Sl_Unlicensed_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(sLBWPGenericR16ExtSlUnlicensedR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sl_Unlicensed_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sl_Unlicensed_r18).Choice {
				case SL_BWP_Generic_r16_Ext_Sl_Unlicensed_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SL_BWP_Generic_r16_Ext_Sl_Unlicensed_r18_Setup:
					if err := (*ie.Sl_Unlicensed_r18).Setup.Encode(ex); err != nil {
						return err
					}
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

func (ie *SL_BWP_Generic_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLBWPGenericR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-BWP-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_BWP_r16 = new(BWP)
			if err := ie.Sl_BWP_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-LengthSymbols-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sLBWPGenericR16SlLengthSymbolsR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_LengthSymbols_r16 = &idx
		}
	}

	// 5. sl-StartSymbol-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sLBWPGenericR16SlStartSymbolR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_StartSymbol_r16 = &idx
		}
	}

	// 6. sl-PSBCH-Config-r16: choice
	{
		if seq.IsComponentPresent(3) {
			ie.Sl_PSBCH_Config_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_PSBCH_Config_r16
			}{}
			choiceDec := d.NewChoiceDecoder(sL_BWP_Generic_r16SlPSBCHConfigR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_PSBCH_Config_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SL_BWP_Generic_r16_Sl_PSBCH_Config_r16_Release:
				(*ie.Sl_PSBCH_Config_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SL_BWP_Generic_r16_Sl_PSBCH_Config_r16_Setup:
				(*ie.Sl_PSBCH_Config_r16).Setup = new(SL_PSBCH_Config_r16)
				if err := (*ie.Sl_PSBCH_Config_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. sl-TxDirectCurrentLocation-r16: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(sLBWPGenericR16SlTxDirectCurrentLocationR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_TxDirectCurrentLocation_r16 = &v
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
				{Name: "sl-Unlicensed-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_Unlicensed_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_Unlicensed_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(sLBWPGenericR16ExtSlUnlicensedR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_Unlicensed_r18).Choice = int(index)
			switch index {
			case SL_BWP_Generic_r16_Ext_Sl_Unlicensed_r18_Release:
				(*ie.Sl_Unlicensed_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SL_BWP_Generic_r16_Ext_Sl_Unlicensed_r18_Setup:
				(*ie.Sl_Unlicensed_r18).Setup = new(SL_Unlicensed_r18)
				if err := (*ie.Sl_Unlicensed_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
