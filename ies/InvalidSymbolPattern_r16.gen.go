// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: InvalidSymbolPattern-r16 (line 8515).

var invalidSymbolPatternR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "symbols-r16"},
		{Name: "periodicityAndPattern-r16", Optional: true},
	},
}

var invalidSymbolPattern_r16SymbolsR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "oneSlot"},
		{Name: "twoSlots"},
	},
}

const (
	InvalidSymbolPattern_r16_Symbols_r16_OneSlot  = 0
	InvalidSymbolPattern_r16_Symbols_r16_TwoSlots = 1
)

var invalidSymbolPattern_r16PeriodicityAndPatternR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "n2"},
		{Name: "n4"},
		{Name: "n5"},
		{Name: "n8"},
		{Name: "n10"},
		{Name: "n20"},
		{Name: "n40"},
	},
}

const (
	InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N2  = 0
	InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N4  = 1
	InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N5  = 2
	InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N8  = 3
	InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N10 = 4
	InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N20 = 5
	InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N40 = 6
)

type InvalidSymbolPattern_r16 struct {
	Symbols_r16 struct {
		Choice   int
		OneSlot  *per.BitString
		TwoSlots *per.BitString
	}
	PeriodicityAndPattern_r16 *struct {
		Choice int
		N2     *per.BitString
		N4     *per.BitString
		N5     *per.BitString
		N8     *per.BitString
		N10    *per.BitString
		N20    *per.BitString
		N40    *per.BitString
	}
}

func (ie *InvalidSymbolPattern_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(invalidSymbolPatternR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PeriodicityAndPattern_r16 != nil}); err != nil {
		return err
	}

	// 3. symbols-r16: choice
	{
		choiceEnc := e.NewChoiceEncoder(invalidSymbolPattern_r16SymbolsR16Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Symbols_r16.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Symbols_r16.Choice {
		case InvalidSymbolPattern_r16_Symbols_r16_OneSlot:
			if err := e.EncodeBitString((*ie.Symbols_r16.OneSlot), per.FixedSize(14)); err != nil {
				return err
			}
		case InvalidSymbolPattern_r16_Symbols_r16_TwoSlots:
			if err := e.EncodeBitString((*ie.Symbols_r16.TwoSlots), per.FixedSize(28)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Symbols_r16.Choice), Constraint: "index out of range"}
		}
	}

	// 4. periodicityAndPattern-r16: choice
	{
		if ie.PeriodicityAndPattern_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(invalidSymbolPattern_r16PeriodicityAndPatternR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.PeriodicityAndPattern_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.PeriodicityAndPattern_r16).Choice {
			case InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N2:
				if err := e.EncodeBitString((*(*ie.PeriodicityAndPattern_r16).N2), per.FixedSize(2)); err != nil {
					return err
				}
			case InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N4:
				if err := e.EncodeBitString((*(*ie.PeriodicityAndPattern_r16).N4), per.FixedSize(4)); err != nil {
					return err
				}
			case InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N5:
				if err := e.EncodeBitString((*(*ie.PeriodicityAndPattern_r16).N5), per.FixedSize(5)); err != nil {
					return err
				}
			case InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N8:
				if err := e.EncodeBitString((*(*ie.PeriodicityAndPattern_r16).N8), per.FixedSize(8)); err != nil {
					return err
				}
			case InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N10:
				if err := e.EncodeBitString((*(*ie.PeriodicityAndPattern_r16).N10), per.FixedSize(10)); err != nil {
					return err
				}
			case InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N20:
				if err := e.EncodeBitString((*(*ie.PeriodicityAndPattern_r16).N20), per.FixedSize(20)); err != nil {
					return err
				}
			case InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N40:
				if err := e.EncodeBitString((*(*ie.PeriodicityAndPattern_r16).N40), per.FixedSize(40)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.PeriodicityAndPattern_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *InvalidSymbolPattern_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(invalidSymbolPatternR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. symbols-r16: choice
	{
		choiceDec := d.NewChoiceDecoder(invalidSymbolPattern_r16SymbolsR16Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Symbols_r16.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case InvalidSymbolPattern_r16_Symbols_r16_OneSlot:
			ie.Symbols_r16.OneSlot = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(14))
			if err != nil {
				return err
			}
			(*ie.Symbols_r16.OneSlot) = v
		case InvalidSymbolPattern_r16_Symbols_r16_TwoSlots:
			ie.Symbols_r16.TwoSlots = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(28))
			if err != nil {
				return err
			}
			(*ie.Symbols_r16.TwoSlots) = v
		}
	}

	// 4. periodicityAndPattern-r16: choice
	{
		if seq.IsComponentPresent(1) {
			ie.PeriodicityAndPattern_r16 = &struct {
				Choice int
				N2     *per.BitString
				N4     *per.BitString
				N5     *per.BitString
				N8     *per.BitString
				N10    *per.BitString
				N20    *per.BitString
				N40    *per.BitString
			}{}
			choiceDec := d.NewChoiceDecoder(invalidSymbolPattern_r16PeriodicityAndPatternR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndPattern_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N2:
				(*ie.PeriodicityAndPattern_r16).N2 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(2))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndPattern_r16).N2) = v
			case InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N4:
				(*ie.PeriodicityAndPattern_r16).N4 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(4))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndPattern_r16).N4) = v
			case InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N5:
				(*ie.PeriodicityAndPattern_r16).N5 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(5))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndPattern_r16).N5) = v
			case InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N8:
				(*ie.PeriodicityAndPattern_r16).N8 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(8))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndPattern_r16).N8) = v
			case InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N10:
				(*ie.PeriodicityAndPattern_r16).N10 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(10))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndPattern_r16).N10) = v
			case InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N20:
				(*ie.PeriodicityAndPattern_r16).N20 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(20))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndPattern_r16).N20) = v
			case InvalidSymbolPattern_r16_PeriodicityAndPattern_r16_N40:
				(*ie.PeriodicityAndPattern_r16).N40 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(40))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndPattern_r16).N40) = v
			}
		}
	}

	return nil
}
