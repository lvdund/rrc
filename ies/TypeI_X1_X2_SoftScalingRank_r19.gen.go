// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TypeI-X1-X2-SoftScalingRank-r19 (line 6416).

var typeIX1X2SoftScalingRankR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "two-one-r19"},
		{Name: "two-two-r19"},
		{Name: "four-one-r19"},
		{Name: "four-two-r19"},
		{Name: "four-four-r19"},
		{Name: "eight-one-r19"},
	},
}

const (
	TypeI_X1_X2_SoftScalingRank_r19_Two_One_r19   = 0
	TypeI_X1_X2_SoftScalingRank_r19_Two_Two_r19   = 1
	TypeI_X1_X2_SoftScalingRank_r19_Four_One_r19  = 2
	TypeI_X1_X2_SoftScalingRank_r19_Four_Two_r19  = 3
	TypeI_X1_X2_SoftScalingRank_r19_Four_Four_r19 = 4
	TypeI_X1_X2_SoftScalingRank_r19_Eight_One_r19 = 5
)

var typeIX1X2SoftScalingRankR19TwoOneR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "twentyFour-r19"},
		{Name: "thirtyTwo-r19"},
		{Name: "sixtyFour-r19"},
	},
}

const (
	TypeI_X1_X2_SoftScalingRank_r19_Two_One_r19_TwentyFour_r19 = 0
	TypeI_X1_X2_SoftScalingRank_r19_Two_One_r19_ThirtyTwo_r19  = 1
	TypeI_X1_X2_SoftScalingRank_r19_Two_One_r19_SixtyFour_r19  = 2
)

var typeIX1X2SoftScalingRankR19TwoTwoR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "twentyFour-r19"},
		{Name: "thirtyTwo-r19"},
		{Name: "sixtyFour-r19"},
	},
}

const (
	TypeI_X1_X2_SoftScalingRank_r19_Two_Two_r19_TwentyFour_r19 = 0
	TypeI_X1_X2_SoftScalingRank_r19_Two_Two_r19_ThirtyTwo_r19  = 1
	TypeI_X1_X2_SoftScalingRank_r19_Two_Two_r19_SixtyFour_r19  = 2
)

var typeIX1X2SoftScalingRankR19FourOneR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "twentyFour-r19"},
		{Name: "thirtyTwo-r19"},
		{Name: "sixtyFour-r19"},
	},
}

const (
	TypeI_X1_X2_SoftScalingRank_r19_Four_One_r19_TwentyFour_r19 = 0
	TypeI_X1_X2_SoftScalingRank_r19_Four_One_r19_ThirtyTwo_r19  = 1
	TypeI_X1_X2_SoftScalingRank_r19_Four_One_r19_SixtyFour_r19  = 2
)

var typeIX1X2SoftScalingRankR19FourTwoR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "twentyFour-r19"},
		{Name: "thirtyTwo-r19"},
		{Name: "sixtyFour-r19"},
	},
}

const (
	TypeI_X1_X2_SoftScalingRank_r19_Four_Two_r19_TwentyFour_r19 = 0
	TypeI_X1_X2_SoftScalingRank_r19_Four_Two_r19_ThirtyTwo_r19  = 1
	TypeI_X1_X2_SoftScalingRank_r19_Four_Two_r19_SixtyFour_r19  = 2
)

var typeIX1X2SoftScalingRankR19FourFourR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "twentyFour-r19"},
		{Name: "thirtyTwo-r19"},
		{Name: "sixtyFour-r19"},
	},
}

const (
	TypeI_X1_X2_SoftScalingRank_r19_Four_Four_r19_TwentyFour_r19 = 0
	TypeI_X1_X2_SoftScalingRank_r19_Four_Four_r19_ThirtyTwo_r19  = 1
	TypeI_X1_X2_SoftScalingRank_r19_Four_Four_r19_SixtyFour_r19  = 2
)

var typeIX1X2SoftScalingRankR19EightOneR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "twentyFour-r19"},
		{Name: "thirtyTwo-r19"},
		{Name: "sixtyFour-r19"},
	},
}

const (
	TypeI_X1_X2_SoftScalingRank_r19_Eight_One_r19_TwentyFour_r19 = 0
	TypeI_X1_X2_SoftScalingRank_r19_Eight_One_r19_ThirtyTwo_r19  = 1
	TypeI_X1_X2_SoftScalingRank_r19_Eight_One_r19_SixtyFour_r19  = 2
)

type TypeI_X1_X2_SoftScalingRank_r19 struct {
	Choice      int
	Two_One_r19 *struct {
		Choice         int
		TwentyFour_r19 *per.BitString
		ThirtyTwo_r19  *per.BitString
		SixtyFour_r19  *per.BitString
	}
	Two_Two_r19 *struct {
		Choice         int
		TwentyFour_r19 *per.BitString
		ThirtyTwo_r19  *per.BitString
		SixtyFour_r19  *per.BitString
	}
	Four_One_r19 *struct {
		Choice         int
		TwentyFour_r19 *per.BitString
		ThirtyTwo_r19  *per.BitString
		SixtyFour_r19  *per.BitString
	}
	Four_Two_r19 *struct {
		Choice         int
		TwentyFour_r19 *per.BitString
		ThirtyTwo_r19  *per.BitString
		SixtyFour_r19  *per.BitString
	}
	Four_Four_r19 *struct {
		Choice         int
		TwentyFour_r19 *per.BitString
		ThirtyTwo_r19  *per.BitString
		SixtyFour_r19  *per.BitString
	}
	Eight_One_r19 *struct {
		Choice         int
		TwentyFour_r19 *per.BitString
		ThirtyTwo_r19  *per.BitString
		SixtyFour_r19  *per.BitString
	}
}

func (ie *TypeI_X1_X2_SoftScalingRank_r19) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(typeIX1X2SoftScalingRankR19Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case TypeI_X1_X2_SoftScalingRank_r19_Two_One_r19:
		choiceEnc := e.NewChoiceEncoder(typeIX1X2SoftScalingRankR19TwoOneR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Two_One_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Two_One_r19).Choice {
		case TypeI_X1_X2_SoftScalingRank_r19_Two_One_r19_TwentyFour_r19:
			if err := e.EncodeBitString((*(*ie.Two_One_r19).TwentyFour_r19), per.FixedSize(576)); err != nil {
				return err
			}
		case TypeI_X1_X2_SoftScalingRank_r19_Two_One_r19_ThirtyTwo_r19:
			if err := e.EncodeBitString((*(*ie.Two_One_r19).ThirtyTwo_r19), per.FixedSize(768)); err != nil {
				return err
			}
		case TypeI_X1_X2_SoftScalingRank_r19_Two_One_r19_SixtyFour_r19:
			if err := e.EncodeBitString((*(*ie.Two_One_r19).SixtyFour_r19), per.FixedSize(1536)); err != nil {
				return err
			}
		}
	case TypeI_X1_X2_SoftScalingRank_r19_Two_Two_r19:
		choiceEnc := e.NewChoiceEncoder(typeIX1X2SoftScalingRankR19TwoTwoR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Two_Two_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Two_Two_r19).Choice {
		case TypeI_X1_X2_SoftScalingRank_r19_Two_Two_r19_TwentyFour_r19:
			if err := e.EncodeBitString((*(*ie.Two_Two_r19).TwentyFour_r19), per.FixedSize(288)); err != nil {
				return err
			}
		case TypeI_X1_X2_SoftScalingRank_r19_Two_Two_r19_ThirtyTwo_r19:
			if err := e.EncodeBitString((*(*ie.Two_Two_r19).ThirtyTwo_r19), per.FixedSize(384)); err != nil {
				return err
			}
		case TypeI_X1_X2_SoftScalingRank_r19_Two_Two_r19_SixtyFour_r19:
			if err := e.EncodeBitString((*(*ie.Two_Two_r19).SixtyFour_r19), per.FixedSize(768)); err != nil {
				return err
			}
		}
	case TypeI_X1_X2_SoftScalingRank_r19_Four_One_r19:
		choiceEnc := e.NewChoiceEncoder(typeIX1X2SoftScalingRankR19FourOneR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Four_One_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Four_One_r19).Choice {
		case TypeI_X1_X2_SoftScalingRank_r19_Four_One_r19_TwentyFour_r19:
			if err := e.EncodeBitString((*(*ie.Four_One_r19).TwentyFour_r19), per.FixedSize(288)); err != nil {
				return err
			}
		case TypeI_X1_X2_SoftScalingRank_r19_Four_One_r19_ThirtyTwo_r19:
			if err := e.EncodeBitString((*(*ie.Four_One_r19).ThirtyTwo_r19), per.FixedSize(384)); err != nil {
				return err
			}
		case TypeI_X1_X2_SoftScalingRank_r19_Four_One_r19_SixtyFour_r19:
			if err := e.EncodeBitString((*(*ie.Four_One_r19).SixtyFour_r19), per.FixedSize(768)); err != nil {
				return err
			}
		}
	case TypeI_X1_X2_SoftScalingRank_r19_Four_Two_r19:
		choiceEnc := e.NewChoiceEncoder(typeIX1X2SoftScalingRankR19FourTwoR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Four_Two_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Four_Two_r19).Choice {
		case TypeI_X1_X2_SoftScalingRank_r19_Four_Two_r19_TwentyFour_r19:
			if err := e.EncodeBitString((*(*ie.Four_Two_r19).TwentyFour_r19), per.FixedSize(144)); err != nil {
				return err
			}
		case TypeI_X1_X2_SoftScalingRank_r19_Four_Two_r19_ThirtyTwo_r19:
			if err := e.EncodeBitString((*(*ie.Four_Two_r19).ThirtyTwo_r19), per.FixedSize(192)); err != nil {
				return err
			}
		case TypeI_X1_X2_SoftScalingRank_r19_Four_Two_r19_SixtyFour_r19:
			if err := e.EncodeBitString((*(*ie.Four_Two_r19).SixtyFour_r19), per.FixedSize(384)); err != nil {
				return err
			}
		}
	case TypeI_X1_X2_SoftScalingRank_r19_Four_Four_r19:
		choiceEnc := e.NewChoiceEncoder(typeIX1X2SoftScalingRankR19FourFourR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Four_Four_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Four_Four_r19).Choice {
		case TypeI_X1_X2_SoftScalingRank_r19_Four_Four_r19_TwentyFour_r19:
			if err := e.EncodeBitString((*(*ie.Four_Four_r19).TwentyFour_r19), per.FixedSize(72)); err != nil {
				return err
			}
		case TypeI_X1_X2_SoftScalingRank_r19_Four_Four_r19_ThirtyTwo_r19:
			if err := e.EncodeBitString((*(*ie.Four_Four_r19).ThirtyTwo_r19), per.FixedSize(96)); err != nil {
				return err
			}
		case TypeI_X1_X2_SoftScalingRank_r19_Four_Four_r19_SixtyFour_r19:
			if err := e.EncodeBitString((*(*ie.Four_Four_r19).SixtyFour_r19), per.FixedSize(192)); err != nil {
				return err
			}
		}
	case TypeI_X1_X2_SoftScalingRank_r19_Eight_One_r19:
		choiceEnc := e.NewChoiceEncoder(typeIX1X2SoftScalingRankR19EightOneR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Eight_One_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Eight_One_r19).Choice {
		case TypeI_X1_X2_SoftScalingRank_r19_Eight_One_r19_TwentyFour_r19:
			if err := e.EncodeBitString((*(*ie.Eight_One_r19).TwentyFour_r19), per.FixedSize(144)); err != nil {
				return err
			}
		case TypeI_X1_X2_SoftScalingRank_r19_Eight_One_r19_ThirtyTwo_r19:
			if err := e.EncodeBitString((*(*ie.Eight_One_r19).ThirtyTwo_r19), per.FixedSize(192)); err != nil {
				return err
			}
		case TypeI_X1_X2_SoftScalingRank_r19_Eight_One_r19_SixtyFour_r19:
			if err := e.EncodeBitString((*(*ie.Eight_One_r19).SixtyFour_r19), per.FixedSize(384)); err != nil {
				return err
			}
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "TypeI-X1-X2-SoftScalingRank-r19",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *TypeI_X1_X2_SoftScalingRank_r19) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(typeIX1X2SoftScalingRankR19Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "TypeI-X1-X2-SoftScalingRank-r19", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case TypeI_X1_X2_SoftScalingRank_r19_Two_One_r19:
		ie.Two_One_r19 = &struct {
			Choice         int
			TwentyFour_r19 *per.BitString
			ThirtyTwo_r19  *per.BitString
			SixtyFour_r19  *per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(typeIX1X2SoftScalingRankR19TwoOneR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Two_One_r19).Choice = int(index)
		switch index {
		case TypeI_X1_X2_SoftScalingRank_r19_Two_One_r19_TwentyFour_r19:
			(*ie.Two_One_r19).TwentyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(576))
			if err != nil {
				return err
			}
			(*(*ie.Two_One_r19).TwentyFour_r19) = v
		case TypeI_X1_X2_SoftScalingRank_r19_Two_One_r19_ThirtyTwo_r19:
			(*ie.Two_One_r19).ThirtyTwo_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(768))
			if err != nil {
				return err
			}
			(*(*ie.Two_One_r19).ThirtyTwo_r19) = v
		case TypeI_X1_X2_SoftScalingRank_r19_Two_One_r19_SixtyFour_r19:
			(*ie.Two_One_r19).SixtyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(1536))
			if err != nil {
				return err
			}
			(*(*ie.Two_One_r19).SixtyFour_r19) = v
		}
	case TypeI_X1_X2_SoftScalingRank_r19_Two_Two_r19:
		ie.Two_Two_r19 = &struct {
			Choice         int
			TwentyFour_r19 *per.BitString
			ThirtyTwo_r19  *per.BitString
			SixtyFour_r19  *per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(typeIX1X2SoftScalingRankR19TwoTwoR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Two_Two_r19).Choice = int(index)
		switch index {
		case TypeI_X1_X2_SoftScalingRank_r19_Two_Two_r19_TwentyFour_r19:
			(*ie.Two_Two_r19).TwentyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(288))
			if err != nil {
				return err
			}
			(*(*ie.Two_Two_r19).TwentyFour_r19) = v
		case TypeI_X1_X2_SoftScalingRank_r19_Two_Two_r19_ThirtyTwo_r19:
			(*ie.Two_Two_r19).ThirtyTwo_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(384))
			if err != nil {
				return err
			}
			(*(*ie.Two_Two_r19).ThirtyTwo_r19) = v
		case TypeI_X1_X2_SoftScalingRank_r19_Two_Two_r19_SixtyFour_r19:
			(*ie.Two_Two_r19).SixtyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(768))
			if err != nil {
				return err
			}
			(*(*ie.Two_Two_r19).SixtyFour_r19) = v
		}
	case TypeI_X1_X2_SoftScalingRank_r19_Four_One_r19:
		ie.Four_One_r19 = &struct {
			Choice         int
			TwentyFour_r19 *per.BitString
			ThirtyTwo_r19  *per.BitString
			SixtyFour_r19  *per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(typeIX1X2SoftScalingRankR19FourOneR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Four_One_r19).Choice = int(index)
		switch index {
		case TypeI_X1_X2_SoftScalingRank_r19_Four_One_r19_TwentyFour_r19:
			(*ie.Four_One_r19).TwentyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(288))
			if err != nil {
				return err
			}
			(*(*ie.Four_One_r19).TwentyFour_r19) = v
		case TypeI_X1_X2_SoftScalingRank_r19_Four_One_r19_ThirtyTwo_r19:
			(*ie.Four_One_r19).ThirtyTwo_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(384))
			if err != nil {
				return err
			}
			(*(*ie.Four_One_r19).ThirtyTwo_r19) = v
		case TypeI_X1_X2_SoftScalingRank_r19_Four_One_r19_SixtyFour_r19:
			(*ie.Four_One_r19).SixtyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(768))
			if err != nil {
				return err
			}
			(*(*ie.Four_One_r19).SixtyFour_r19) = v
		}
	case TypeI_X1_X2_SoftScalingRank_r19_Four_Two_r19:
		ie.Four_Two_r19 = &struct {
			Choice         int
			TwentyFour_r19 *per.BitString
			ThirtyTwo_r19  *per.BitString
			SixtyFour_r19  *per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(typeIX1X2SoftScalingRankR19FourTwoR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Four_Two_r19).Choice = int(index)
		switch index {
		case TypeI_X1_X2_SoftScalingRank_r19_Four_Two_r19_TwentyFour_r19:
			(*ie.Four_Two_r19).TwentyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(144))
			if err != nil {
				return err
			}
			(*(*ie.Four_Two_r19).TwentyFour_r19) = v
		case TypeI_X1_X2_SoftScalingRank_r19_Four_Two_r19_ThirtyTwo_r19:
			(*ie.Four_Two_r19).ThirtyTwo_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(192))
			if err != nil {
				return err
			}
			(*(*ie.Four_Two_r19).ThirtyTwo_r19) = v
		case TypeI_X1_X2_SoftScalingRank_r19_Four_Two_r19_SixtyFour_r19:
			(*ie.Four_Two_r19).SixtyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(384))
			if err != nil {
				return err
			}
			(*(*ie.Four_Two_r19).SixtyFour_r19) = v
		}
	case TypeI_X1_X2_SoftScalingRank_r19_Four_Four_r19:
		ie.Four_Four_r19 = &struct {
			Choice         int
			TwentyFour_r19 *per.BitString
			ThirtyTwo_r19  *per.BitString
			SixtyFour_r19  *per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(typeIX1X2SoftScalingRankR19FourFourR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Four_Four_r19).Choice = int(index)
		switch index {
		case TypeI_X1_X2_SoftScalingRank_r19_Four_Four_r19_TwentyFour_r19:
			(*ie.Four_Four_r19).TwentyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(72))
			if err != nil {
				return err
			}
			(*(*ie.Four_Four_r19).TwentyFour_r19) = v
		case TypeI_X1_X2_SoftScalingRank_r19_Four_Four_r19_ThirtyTwo_r19:
			(*ie.Four_Four_r19).ThirtyTwo_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(96))
			if err != nil {
				return err
			}
			(*(*ie.Four_Four_r19).ThirtyTwo_r19) = v
		case TypeI_X1_X2_SoftScalingRank_r19_Four_Four_r19_SixtyFour_r19:
			(*ie.Four_Four_r19).SixtyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(192))
			if err != nil {
				return err
			}
			(*(*ie.Four_Four_r19).SixtyFour_r19) = v
		}
	case TypeI_X1_X2_SoftScalingRank_r19_Eight_One_r19:
		ie.Eight_One_r19 = &struct {
			Choice         int
			TwentyFour_r19 *per.BitString
			ThirtyTwo_r19  *per.BitString
			SixtyFour_r19  *per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(typeIX1X2SoftScalingRankR19EightOneR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Eight_One_r19).Choice = int(index)
		switch index {
		case TypeI_X1_X2_SoftScalingRank_r19_Eight_One_r19_TwentyFour_r19:
			(*ie.Eight_One_r19).TwentyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(144))
			if err != nil {
				return err
			}
			(*(*ie.Eight_One_r19).TwentyFour_r19) = v
		case TypeI_X1_X2_SoftScalingRank_r19_Eight_One_r19_ThirtyTwo_r19:
			(*ie.Eight_One_r19).ThirtyTwo_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(192))
			if err != nil {
				return err
			}
			(*(*ie.Eight_One_r19).ThirtyTwo_r19) = v
		case TypeI_X1_X2_SoftScalingRank_r19_Eight_One_r19_SixtyFour_r19:
			(*ie.Eight_One_r19).SixtyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(384))
			if err != nil {
				return err
			}
			(*(*ie.Eight_One_r19).SixtyFour_r19) = v
		}
	default:
		return &per.DecodeError{TypeName: "TypeI-X1-X2-SoftScalingRank-r19", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
