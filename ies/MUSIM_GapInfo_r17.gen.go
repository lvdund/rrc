// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MUSIM-GapInfo-r17 (line 10221).

var mUSIMGapInfoR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "musim-Starting-SFN-AndSubframe-r17", Optional: true},
		{Name: "musim-GapLength-r17", Optional: true},
		{Name: "musim-GapRepetitionAndOffset-r17", Optional: true},
	},
}

const (
	MUSIM_GapInfo_r17_Musim_GapLength_r17_Ms3  = 0
	MUSIM_GapInfo_r17_Musim_GapLength_r17_Ms4  = 1
	MUSIM_GapInfo_r17_Musim_GapLength_r17_Ms6  = 2
	MUSIM_GapInfo_r17_Musim_GapLength_r17_Ms10 = 3
	MUSIM_GapInfo_r17_Musim_GapLength_r17_Ms20 = 4
)

var mUSIMGapInfoR17MusimGapLengthR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MUSIM_GapInfo_r17_Musim_GapLength_r17_Ms3, MUSIM_GapInfo_r17_Musim_GapLength_r17_Ms4, MUSIM_GapInfo_r17_Musim_GapLength_r17_Ms6, MUSIM_GapInfo_r17_Musim_GapLength_r17_Ms10, MUSIM_GapInfo_r17_Musim_GapLength_r17_Ms20},
}

var mUSIM_GapInfo_r17MusimGapRepetitionAndOffsetR17Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ms20-r17"},
		{Name: "ms40-r17"},
		{Name: "ms80-r17"},
		{Name: "ms160-r17"},
		{Name: "ms320-r17"},
		{Name: "ms640-r17"},
		{Name: "ms1280-r17"},
		{Name: "ms2560-r17"},
		{Name: "ms5120-r17"},
	},
}

const (
	MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms20_r17   = 0
	MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms40_r17   = 1
	MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms80_r17   = 2
	MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms160_r17  = 3
	MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms320_r17  = 4
	MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms640_r17  = 5
	MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms1280_r17 = 6
	MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms2560_r17 = 7
	MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms5120_r17 = 8
)

type MUSIM_GapInfo_r17 struct {
	Musim_Starting_SFN_AndSubframe_r17 *MUSIM_Starting_SFN_AndSubframe_r17
	Musim_GapLength_r17                *int64
	Musim_GapRepetitionAndOffset_r17   *struct {
		Choice     int
		Ms20_r17   *int64
		Ms40_r17   *int64
		Ms80_r17   *int64
		Ms160_r17  *int64
		Ms320_r17  *int64
		Ms640_r17  *int64
		Ms1280_r17 *int64
		Ms2560_r17 *int64
		Ms5120_r17 *int64
	}
}

func (ie *MUSIM_GapInfo_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mUSIMGapInfoR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Musim_Starting_SFN_AndSubframe_r17 != nil, ie.Musim_GapLength_r17 != nil, ie.Musim_GapRepetitionAndOffset_r17 != nil}); err != nil {
		return err
	}

	// 2. musim-Starting-SFN-AndSubframe-r17: ref
	{
		if ie.Musim_Starting_SFN_AndSubframe_r17 != nil {
			if err := ie.Musim_Starting_SFN_AndSubframe_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. musim-GapLength-r17: enumerated
	{
		if ie.Musim_GapLength_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Musim_GapLength_r17, mUSIMGapInfoR17MusimGapLengthR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. musim-GapRepetitionAndOffset-r17: choice
	{
		if ie.Musim_GapRepetitionAndOffset_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(mUSIM_GapInfo_r17MusimGapRepetitionAndOffsetR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Musim_GapRepetitionAndOffset_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Musim_GapRepetitionAndOffset_r17).Choice {
			case MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms20_r17:
				if err := e.EncodeInteger((*(*ie.Musim_GapRepetitionAndOffset_r17).Ms20_r17), per.Constrained(0, 19)); err != nil {
					return err
				}
			case MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms40_r17:
				if err := e.EncodeInteger((*(*ie.Musim_GapRepetitionAndOffset_r17).Ms40_r17), per.Constrained(0, 39)); err != nil {
					return err
				}
			case MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms80_r17:
				if err := e.EncodeInteger((*(*ie.Musim_GapRepetitionAndOffset_r17).Ms80_r17), per.Constrained(0, 79)); err != nil {
					return err
				}
			case MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms160_r17:
				if err := e.EncodeInteger((*(*ie.Musim_GapRepetitionAndOffset_r17).Ms160_r17), per.Constrained(0, 159)); err != nil {
					return err
				}
			case MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms320_r17:
				if err := e.EncodeInteger((*(*ie.Musim_GapRepetitionAndOffset_r17).Ms320_r17), per.Constrained(0, 319)); err != nil {
					return err
				}
			case MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms640_r17:
				if err := e.EncodeInteger((*(*ie.Musim_GapRepetitionAndOffset_r17).Ms640_r17), per.Constrained(0, 639)); err != nil {
					return err
				}
			case MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms1280_r17:
				if err := e.EncodeInteger((*(*ie.Musim_GapRepetitionAndOffset_r17).Ms1280_r17), per.Constrained(0, 1279)); err != nil {
					return err
				}
			case MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms2560_r17:
				if err := e.EncodeInteger((*(*ie.Musim_GapRepetitionAndOffset_r17).Ms2560_r17), per.Constrained(0, 2559)); err != nil {
					return err
				}
			case MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms5120_r17:
				if err := e.EncodeInteger((*(*ie.Musim_GapRepetitionAndOffset_r17).Ms5120_r17), per.Constrained(0, 5119)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Musim_GapRepetitionAndOffset_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *MUSIM_GapInfo_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mUSIMGapInfoR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. musim-Starting-SFN-AndSubframe-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Musim_Starting_SFN_AndSubframe_r17 = new(MUSIM_Starting_SFN_AndSubframe_r17)
			if err := ie.Musim_Starting_SFN_AndSubframe_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. musim-GapLength-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(mUSIMGapInfoR17MusimGapLengthR17Constraints)
			if err != nil {
				return err
			}
			ie.Musim_GapLength_r17 = &idx
		}
	}

	// 4. musim-GapRepetitionAndOffset-r17: choice
	{
		if seq.IsComponentPresent(2) {
			ie.Musim_GapRepetitionAndOffset_r17 = &struct {
				Choice     int
				Ms20_r17   *int64
				Ms40_r17   *int64
				Ms80_r17   *int64
				Ms160_r17  *int64
				Ms320_r17  *int64
				Ms640_r17  *int64
				Ms1280_r17 *int64
				Ms2560_r17 *int64
				Ms5120_r17 *int64
			}{}
			choiceDec := d.NewChoiceDecoder(mUSIM_GapInfo_r17MusimGapRepetitionAndOffsetR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Musim_GapRepetitionAndOffset_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms20_r17:
				(*ie.Musim_GapRepetitionAndOffset_r17).Ms20_r17 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 19))
				if err != nil {
					return err
				}
				(*(*ie.Musim_GapRepetitionAndOffset_r17).Ms20_r17) = v
			case MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms40_r17:
				(*ie.Musim_GapRepetitionAndOffset_r17).Ms40_r17 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 39))
				if err != nil {
					return err
				}
				(*(*ie.Musim_GapRepetitionAndOffset_r17).Ms40_r17) = v
			case MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms80_r17:
				(*ie.Musim_GapRepetitionAndOffset_r17).Ms80_r17 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 79))
				if err != nil {
					return err
				}
				(*(*ie.Musim_GapRepetitionAndOffset_r17).Ms80_r17) = v
			case MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms160_r17:
				(*ie.Musim_GapRepetitionAndOffset_r17).Ms160_r17 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 159))
				if err != nil {
					return err
				}
				(*(*ie.Musim_GapRepetitionAndOffset_r17).Ms160_r17) = v
			case MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms320_r17:
				(*ie.Musim_GapRepetitionAndOffset_r17).Ms320_r17 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 319))
				if err != nil {
					return err
				}
				(*(*ie.Musim_GapRepetitionAndOffset_r17).Ms320_r17) = v
			case MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms640_r17:
				(*ie.Musim_GapRepetitionAndOffset_r17).Ms640_r17 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 639))
				if err != nil {
					return err
				}
				(*(*ie.Musim_GapRepetitionAndOffset_r17).Ms640_r17) = v
			case MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms1280_r17:
				(*ie.Musim_GapRepetitionAndOffset_r17).Ms1280_r17 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 1279))
				if err != nil {
					return err
				}
				(*(*ie.Musim_GapRepetitionAndOffset_r17).Ms1280_r17) = v
			case MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms2560_r17:
				(*ie.Musim_GapRepetitionAndOffset_r17).Ms2560_r17 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 2559))
				if err != nil {
					return err
				}
				(*(*ie.Musim_GapRepetitionAndOffset_r17).Ms2560_r17) = v
			case MUSIM_GapInfo_r17_Musim_GapRepetitionAndOffset_r17_Ms5120_r17:
				(*ie.Musim_GapRepetitionAndOffset_r17).Ms5120_r17 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 5119))
				if err != nil {
					return err
				}
				(*(*ie.Musim_GapRepetitionAndOffset_r17).Ms5120_r17) = v
			}
		}
	}

	return nil
}
