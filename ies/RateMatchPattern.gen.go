// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RateMatchPattern (line 13258).

var rateMatchPatternConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rateMatchPatternId"},
		{Name: "patternType"},
		{Name: "subcarrierSpacing", Optional: true},
		{Name: "dummy"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var rateMatchPatternPatternTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "bitmaps"},
		{Name: "controlResourceSet"},
	},
}

const (
	RateMatchPattern_PatternType_Bitmaps            = 0
	RateMatchPattern_PatternType_ControlResourceSet = 1
)

const (
	RateMatchPattern_Dummy_Dynamic    = 0
	RateMatchPattern_Dummy_SemiStatic = 1
)

var rateMatchPatternDummyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RateMatchPattern_Dummy_Dynamic, RateMatchPattern_Dummy_SemiStatic},
}

var rateMatchPatternPatternTypeBitmapsConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "resourceBlocks"},
		{Name: "symbolsInResourceBlock"},
		{Name: "periodicityAndPattern", Optional: true},
	},
}

var rateMatchPatternPatternTypeBitmapsSymbolsInResourceBlockConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "oneSlot"},
		{Name: "twoSlots"},
	},
}

const (
	RateMatchPattern_PatternType_Bitmaps_SymbolsInResourceBlock_OneSlot  = 0
	RateMatchPattern_PatternType_Bitmaps_SymbolsInResourceBlock_TwoSlots = 1
)

var rateMatchPatternPatternTypeBitmapsPeriodicityAndPatternConstraints = per.ChoiceConstraints{
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
	RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N2  = 0
	RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N4  = 1
	RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N5  = 2
	RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N8  = 3
	RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N10 = 4
	RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N20 = 5
	RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N40 = 6
)

type RateMatchPattern struct {
	RateMatchPatternId RateMatchPatternId
	PatternType        struct {
		Choice  int
		Bitmaps *struct {
			ResourceBlocks         per.BitString
			SymbolsInResourceBlock struct {
				Choice   int
				OneSlot  *per.BitString
				TwoSlots *per.BitString
			}
			PeriodicityAndPattern *struct {
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
		ControlResourceSet *ControlResourceSetId
	}
	SubcarrierSpacing      *SubcarrierSpacing
	Dummy                  int64
	ControlResourceSet_r16 *ControlResourceSetId_r16
}

func (ie *RateMatchPattern) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rateMatchPatternConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ControlResourceSet_r16 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SubcarrierSpacing != nil}); err != nil {
		return err
	}

	// 3. rateMatchPatternId: ref
	{
		if err := ie.RateMatchPatternId.Encode(e); err != nil {
			return err
		}
	}

	// 4. patternType: choice
	{
		choiceEnc := e.NewChoiceEncoder(rateMatchPatternPatternTypeConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.PatternType.Choice), false, nil); err != nil {
			return err
		}
		switch ie.PatternType.Choice {
		case RateMatchPattern_PatternType_Bitmaps:
			c := (*ie.PatternType.Bitmaps)
			rateMatchPatternPatternTypeBitmapsSeq := e.NewSequenceEncoder(rateMatchPatternPatternTypeBitmapsConstraints)
			if err := rateMatchPatternPatternTypeBitmapsSeq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := rateMatchPatternPatternTypeBitmapsSeq.EncodePreamble([]bool{c.PeriodicityAndPattern != nil}); err != nil {
				return err
			}
			if err := e.EncodeBitString(c.ResourceBlocks, per.FixedSize(275)); err != nil {
				return err
			}
			{
				choiceEnc := e.NewChoiceEncoder(rateMatchPatternPatternTypeBitmapsSymbolsInResourceBlockConstraints)
				if err := choiceEnc.EncodeChoice(int64(c.SymbolsInResourceBlock.Choice), false, nil); err != nil {
					return err
				}
				switch c.SymbolsInResourceBlock.Choice {
				case RateMatchPattern_PatternType_Bitmaps_SymbolsInResourceBlock_OneSlot:
					if err := e.EncodeBitString((*c.SymbolsInResourceBlock.OneSlot), per.FixedSize(14)); err != nil {
						return err
					}
				case RateMatchPattern_PatternType_Bitmaps_SymbolsInResourceBlock_TwoSlots:
					if err := e.EncodeBitString((*c.SymbolsInResourceBlock.TwoSlots), per.FixedSize(28)); err != nil {
						return err
					}
				}
			}
			if c.PeriodicityAndPattern != nil {
				choiceEnc := e.NewChoiceEncoder(rateMatchPatternPatternTypeBitmapsPeriodicityAndPatternConstraints)
				if err := choiceEnc.EncodeChoice(int64((*c.PeriodicityAndPattern).Choice), false, nil); err != nil {
					return err
				}
				switch (*c.PeriodicityAndPattern).Choice {
				case RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N2:
					if err := e.EncodeBitString((*(*c.PeriodicityAndPattern).N2), per.FixedSize(2)); err != nil {
						return err
					}
				case RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N4:
					if err := e.EncodeBitString((*(*c.PeriodicityAndPattern).N4), per.FixedSize(4)); err != nil {
						return err
					}
				case RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N5:
					if err := e.EncodeBitString((*(*c.PeriodicityAndPattern).N5), per.FixedSize(5)); err != nil {
						return err
					}
				case RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N8:
					if err := e.EncodeBitString((*(*c.PeriodicityAndPattern).N8), per.FixedSize(8)); err != nil {
						return err
					}
				case RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N10:
					if err := e.EncodeBitString((*(*c.PeriodicityAndPattern).N10), per.FixedSize(10)); err != nil {
						return err
					}
				case RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N20:
					if err := e.EncodeBitString((*(*c.PeriodicityAndPattern).N20), per.FixedSize(20)); err != nil {
						return err
					}
				case RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N40:
					if err := e.EncodeBitString((*(*c.PeriodicityAndPattern).N40), per.FixedSize(40)); err != nil {
						return err
					}
				}
			}
		case RateMatchPattern_PatternType_ControlResourceSet:
			if err := ie.PatternType.ControlResourceSet.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.PatternType.Choice), Constraint: "index out of range"}
		}
	}

	// 5. subcarrierSpacing: ref
	{
		if ie.SubcarrierSpacing != nil {
			if err := ie.SubcarrierSpacing.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. dummy: enumerated
	{
		if err := e.EncodeEnumerated(ie.Dummy, rateMatchPatternDummyConstraints); err != nil {
			return err
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
					{Name: "controlResourceSet-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ControlResourceSet_r16 != nil}); err != nil {
				return err
			}

			if ie.ControlResourceSet_r16 != nil {
				if err := ie.ControlResourceSet_r16.Encode(ex); err != nil {
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

func (ie *RateMatchPattern) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rateMatchPatternConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. rateMatchPatternId: ref
	{
		if err := ie.RateMatchPatternId.Decode(d); err != nil {
			return err
		}
	}

	// 4. patternType: choice
	{
		choiceDec := d.NewChoiceDecoder(rateMatchPatternPatternTypeConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.PatternType.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case RateMatchPattern_PatternType_Bitmaps:
			ie.PatternType.Bitmaps = &struct {
				ResourceBlocks         per.BitString
				SymbolsInResourceBlock struct {
					Choice   int
					OneSlot  *per.BitString
					TwoSlots *per.BitString
				}
				PeriodicityAndPattern *struct {
					Choice int
					N2     *per.BitString
					N4     *per.BitString
					N5     *per.BitString
					N8     *per.BitString
					N10    *per.BitString
					N20    *per.BitString
					N40    *per.BitString
				}
			}{}
			c := (*ie.PatternType.Bitmaps)
			rateMatchPatternPatternTypeBitmapsSeq := d.NewSequenceDecoder(rateMatchPatternPatternTypeBitmapsConstraints)
			if err := rateMatchPatternPatternTypeBitmapsSeq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := rateMatchPatternPatternTypeBitmapsSeq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeBitString(per.FixedSize(275))
				if err != nil {
					return err
				}
				c.ResourceBlocks = v
			}
			{
				choiceDec := d.NewChoiceDecoder(rateMatchPatternPatternTypeBitmapsSymbolsInResourceBlockConstraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.SymbolsInResourceBlock.Choice = int(index)
				switch index {
				case RateMatchPattern_PatternType_Bitmaps_SymbolsInResourceBlock_OneSlot:
					c.SymbolsInResourceBlock.OneSlot = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(14))
					if err != nil {
						return err
					}
					(*c.SymbolsInResourceBlock.OneSlot) = v
				case RateMatchPattern_PatternType_Bitmaps_SymbolsInResourceBlock_TwoSlots:
					c.SymbolsInResourceBlock.TwoSlots = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(28))
					if err != nil {
						return err
					}
					(*c.SymbolsInResourceBlock.TwoSlots) = v
				}
			}
			if rateMatchPatternPatternTypeBitmapsSeq.IsComponentPresent(2) {
				c.PeriodicityAndPattern = &struct {
					Choice int
					N2     *per.BitString
					N4     *per.BitString
					N5     *per.BitString
					N8     *per.BitString
					N10    *per.BitString
					N20    *per.BitString
					N40    *per.BitString
				}{}
				choiceDec := d.NewChoiceDecoder(rateMatchPatternPatternTypeBitmapsPeriodicityAndPatternConstraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*c.PeriodicityAndPattern).Choice = int(index)
				switch index {
				case RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N2:
					(*c.PeriodicityAndPattern).N2 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(2))
					if err != nil {
						return err
					}
					(*(*c.PeriodicityAndPattern).N2) = v
				case RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N4:
					(*c.PeriodicityAndPattern).N4 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(4))
					if err != nil {
						return err
					}
					(*(*c.PeriodicityAndPattern).N4) = v
				case RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N5:
					(*c.PeriodicityAndPattern).N5 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(5))
					if err != nil {
						return err
					}
					(*(*c.PeriodicityAndPattern).N5) = v
				case RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N8:
					(*c.PeriodicityAndPattern).N8 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(8))
					if err != nil {
						return err
					}
					(*(*c.PeriodicityAndPattern).N8) = v
				case RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N10:
					(*c.PeriodicityAndPattern).N10 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(10))
					if err != nil {
						return err
					}
					(*(*c.PeriodicityAndPattern).N10) = v
				case RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N20:
					(*c.PeriodicityAndPattern).N20 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(20))
					if err != nil {
						return err
					}
					(*(*c.PeriodicityAndPattern).N20) = v
				case RateMatchPattern_PatternType_Bitmaps_PeriodicityAndPattern_N40:
					(*c.PeriodicityAndPattern).N40 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(40))
					if err != nil {
						return err
					}
					(*(*c.PeriodicityAndPattern).N40) = v
				}
			}
		case RateMatchPattern_PatternType_ControlResourceSet:
			ie.PatternType.ControlResourceSet = new(ControlResourceSetId)
			if err := ie.PatternType.ControlResourceSet.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. subcarrierSpacing: ref
	{
		if seq.IsComponentPresent(2) {
			ie.SubcarrierSpacing = new(SubcarrierSpacing)
			if err := ie.SubcarrierSpacing.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. dummy: enumerated
	{
		v3, err := d.DecodeEnumerated(rateMatchPatternDummyConstraints)
		if err != nil {
			return err
		}
		ie.Dummy = v3
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
				{Name: "controlResourceSet-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.ControlResourceSet_r16 = new(ControlResourceSetId_r16)
			if err := ie.ControlResourceSet_r16.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
