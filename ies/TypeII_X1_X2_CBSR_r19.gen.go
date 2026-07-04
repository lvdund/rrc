// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TypeII-X1-X2-CBSR-r19 (line 6431).

var typeIIX1X2CBSRR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "one-one-r19"},
		{Name: "two-one-r19"},
		{Name: "two-two-r19"},
		{Name: "four-one-r19"},
		{Name: "four-two-r19"},
	},
}

const (
	TypeII_X1_X2_CBSR_r19_One_One_r19  = 0
	TypeII_X1_X2_CBSR_r19_Two_One_r19  = 1
	TypeII_X1_X2_CBSR_r19_Two_Two_r19  = 2
	TypeII_X1_X2_CBSR_r19_Four_One_r19 = 3
	TypeII_X1_X2_CBSR_r19_Four_Two_r19 = 4
)

var typeIIX1X2CBSRR19OneOneR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "twentyFour-r19"},
		{Name: "thirtyTwo-r19"},
		{Name: "sixtyFour-r19"},
	},
}

const (
	TypeII_X1_X2_CBSR_r19_One_One_r19_TwentyFour_r19 = 0
	TypeII_X1_X2_CBSR_r19_One_One_r19_ThirtyTwo_r19  = 1
	TypeII_X1_X2_CBSR_r19_One_One_r19_SixtyFour_r19  = 2
)

var typeIIX1X2CBSRR19TwoOneR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "twentyFour-r19"},
		{Name: "thirtyTwo-r19"},
		{Name: "sixtyFour-r19"},
	},
}

const (
	TypeII_X1_X2_CBSR_r19_Two_One_r19_TwentyFour_r19 = 0
	TypeII_X1_X2_CBSR_r19_Two_One_r19_ThirtyTwo_r19  = 1
	TypeII_X1_X2_CBSR_r19_Two_One_r19_SixtyFour_r19  = 2
)

var typeIIX1X2CBSRR19TwoTwoR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "twentyFour-r19"},
		{Name: "thirtyTwo-r19"},
		{Name: "sixtyFour-r19"},
	},
}

const (
	TypeII_X1_X2_CBSR_r19_Two_Two_r19_TwentyFour_r19 = 0
	TypeII_X1_X2_CBSR_r19_Two_Two_r19_ThirtyTwo_r19  = 1
	TypeII_X1_X2_CBSR_r19_Two_Two_r19_SixtyFour_r19  = 2
)

var typeIIX1X2CBSRR19FourOneR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "twentyFour-r19"},
		{Name: "thirtyTwo-r19"},
		{Name: "sixtyFour-r19"},
	},
}

const (
	TypeII_X1_X2_CBSR_r19_Four_One_r19_TwentyFour_r19 = 0
	TypeII_X1_X2_CBSR_r19_Four_One_r19_ThirtyTwo_r19  = 1
	TypeII_X1_X2_CBSR_r19_Four_One_r19_SixtyFour_r19  = 2
)

var typeIIX1X2CBSRR19FourTwoR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "thirtyTwo-r19"},
		{Name: "sixtyFour-r19"},
	},
}

const (
	TypeII_X1_X2_CBSR_r19_Four_Two_r19_ThirtyTwo_r19 = 0
	TypeII_X1_X2_CBSR_r19_Four_Two_r19_SixtyFour_r19 = 1
)

type TypeII_X1_X2_CBSR_r19 struct {
	Choice      int
	One_One_r19 *struct {
		Choice         int
		TwentyFour_r19 *per.BitString
		ThirtyTwo_r19  *per.BitString
		SixtyFour_r19  *per.BitString
	}
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
		Choice        int
		ThirtyTwo_r19 *per.BitString
		SixtyFour_r19 *per.BitString
	}
}

func (ie *TypeII_X1_X2_CBSR_r19) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(typeIIX1X2CBSRR19Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case TypeII_X1_X2_CBSR_r19_One_One_r19:
		choiceEnc := e.NewChoiceEncoder(typeIIX1X2CBSRR19OneOneR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.One_One_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.One_One_r19).Choice {
		case TypeII_X1_X2_CBSR_r19_One_One_r19_TwentyFour_r19:
			if err := e.EncodeBitString((*(*ie.One_One_r19).TwentyFour_r19), per.FixedSize(24)); err != nil {
				return err
			}
		case TypeII_X1_X2_CBSR_r19_One_One_r19_ThirtyTwo_r19:
			if err := e.EncodeBitString((*(*ie.One_One_r19).ThirtyTwo_r19), per.FixedSize(32)); err != nil {
				return err
			}
		case TypeII_X1_X2_CBSR_r19_One_One_r19_SixtyFour_r19:
			if err := e.EncodeBitString((*(*ie.One_One_r19).SixtyFour_r19), per.FixedSize(64)); err != nil {
				return err
			}
		}
	case TypeII_X1_X2_CBSR_r19_Two_One_r19:
		choiceEnc := e.NewChoiceEncoder(typeIIX1X2CBSRR19TwoOneR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Two_One_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Two_One_r19).Choice {
		case TypeII_X1_X2_CBSR_r19_Two_One_r19_TwentyFour_r19:
			if err := e.EncodeBitString((*(*ie.Two_One_r19).TwentyFour_r19), per.FixedSize(12)); err != nil {
				return err
			}
		case TypeII_X1_X2_CBSR_r19_Two_One_r19_ThirtyTwo_r19:
			if err := e.EncodeBitString((*(*ie.Two_One_r19).ThirtyTwo_r19), per.FixedSize(16)); err != nil {
				return err
			}
		case TypeII_X1_X2_CBSR_r19_Two_One_r19_SixtyFour_r19:
			if err := e.EncodeBitString((*(*ie.Two_One_r19).SixtyFour_r19), per.FixedSize(32)); err != nil {
				return err
			}
		}
	case TypeII_X1_X2_CBSR_r19_Two_Two_r19:
		choiceEnc := e.NewChoiceEncoder(typeIIX1X2CBSRR19TwoTwoR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Two_Two_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Two_Two_r19).Choice {
		case TypeII_X1_X2_CBSR_r19_Two_Two_r19_TwentyFour_r19:
			if err := e.EncodeBitString((*(*ie.Two_Two_r19).TwentyFour_r19), per.FixedSize(6)); err != nil {
				return err
			}
		case TypeII_X1_X2_CBSR_r19_Two_Two_r19_ThirtyTwo_r19:
			if err := e.EncodeBitString((*(*ie.Two_Two_r19).ThirtyTwo_r19), per.FixedSize(8)); err != nil {
				return err
			}
		case TypeII_X1_X2_CBSR_r19_Two_Two_r19_SixtyFour_r19:
			if err := e.EncodeBitString((*(*ie.Two_Two_r19).SixtyFour_r19), per.FixedSize(16)); err != nil {
				return err
			}
		}
	case TypeII_X1_X2_CBSR_r19_Four_One_r19:
		choiceEnc := e.NewChoiceEncoder(typeIIX1X2CBSRR19FourOneR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Four_One_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Four_One_r19).Choice {
		case TypeII_X1_X2_CBSR_r19_Four_One_r19_TwentyFour_r19:
			if err := e.EncodeBitString((*(*ie.Four_One_r19).TwentyFour_r19), per.FixedSize(6)); err != nil {
				return err
			}
		case TypeII_X1_X2_CBSR_r19_Four_One_r19_ThirtyTwo_r19:
			if err := e.EncodeBitString((*(*ie.Four_One_r19).ThirtyTwo_r19), per.FixedSize(8)); err != nil {
				return err
			}
		case TypeII_X1_X2_CBSR_r19_Four_One_r19_SixtyFour_r19:
			if err := e.EncodeBitString((*(*ie.Four_One_r19).SixtyFour_r19), per.FixedSize(16)); err != nil {
				return err
			}
		}
	case TypeII_X1_X2_CBSR_r19_Four_Two_r19:
		choiceEnc := e.NewChoiceEncoder(typeIIX1X2CBSRR19FourTwoR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Four_Two_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Four_Two_r19).Choice {
		case TypeII_X1_X2_CBSR_r19_Four_Two_r19_ThirtyTwo_r19:
			if err := e.EncodeBitString((*(*ie.Four_Two_r19).ThirtyTwo_r19), per.FixedSize(4)); err != nil {
				return err
			}
		case TypeII_X1_X2_CBSR_r19_Four_Two_r19_SixtyFour_r19:
			if err := e.EncodeBitString((*(*ie.Four_Two_r19).SixtyFour_r19), per.FixedSize(8)); err != nil {
				return err
			}
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "TypeII-X1-X2-CBSR-r19",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *TypeII_X1_X2_CBSR_r19) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(typeIIX1X2CBSRR19Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "TypeII-X1-X2-CBSR-r19", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case TypeII_X1_X2_CBSR_r19_One_One_r19:
		ie.One_One_r19 = &struct {
			Choice         int
			TwentyFour_r19 *per.BitString
			ThirtyTwo_r19  *per.BitString
			SixtyFour_r19  *per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(typeIIX1X2CBSRR19OneOneR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.One_One_r19).Choice = int(index)
		switch index {
		case TypeII_X1_X2_CBSR_r19_One_One_r19_TwentyFour_r19:
			(*ie.One_One_r19).TwentyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(24))
			if err != nil {
				return err
			}
			(*(*ie.One_One_r19).TwentyFour_r19) = v
		case TypeII_X1_X2_CBSR_r19_One_One_r19_ThirtyTwo_r19:
			(*ie.One_One_r19).ThirtyTwo_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(32))
			if err != nil {
				return err
			}
			(*(*ie.One_One_r19).ThirtyTwo_r19) = v
		case TypeII_X1_X2_CBSR_r19_One_One_r19_SixtyFour_r19:
			(*ie.One_One_r19).SixtyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(64))
			if err != nil {
				return err
			}
			(*(*ie.One_One_r19).SixtyFour_r19) = v
		}
	case TypeII_X1_X2_CBSR_r19_Two_One_r19:
		ie.Two_One_r19 = &struct {
			Choice         int
			TwentyFour_r19 *per.BitString
			ThirtyTwo_r19  *per.BitString
			SixtyFour_r19  *per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(typeIIX1X2CBSRR19TwoOneR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Two_One_r19).Choice = int(index)
		switch index {
		case TypeII_X1_X2_CBSR_r19_Two_One_r19_TwentyFour_r19:
			(*ie.Two_One_r19).TwentyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(12))
			if err != nil {
				return err
			}
			(*(*ie.Two_One_r19).TwentyFour_r19) = v
		case TypeII_X1_X2_CBSR_r19_Two_One_r19_ThirtyTwo_r19:
			(*ie.Two_One_r19).ThirtyTwo_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(16))
			if err != nil {
				return err
			}
			(*(*ie.Two_One_r19).ThirtyTwo_r19) = v
		case TypeII_X1_X2_CBSR_r19_Two_One_r19_SixtyFour_r19:
			(*ie.Two_One_r19).SixtyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(32))
			if err != nil {
				return err
			}
			(*(*ie.Two_One_r19).SixtyFour_r19) = v
		}
	case TypeII_X1_X2_CBSR_r19_Two_Two_r19:
		ie.Two_Two_r19 = &struct {
			Choice         int
			TwentyFour_r19 *per.BitString
			ThirtyTwo_r19  *per.BitString
			SixtyFour_r19  *per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(typeIIX1X2CBSRR19TwoTwoR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Two_Two_r19).Choice = int(index)
		switch index {
		case TypeII_X1_X2_CBSR_r19_Two_Two_r19_TwentyFour_r19:
			(*ie.Two_Two_r19).TwentyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(6))
			if err != nil {
				return err
			}
			(*(*ie.Two_Two_r19).TwentyFour_r19) = v
		case TypeII_X1_X2_CBSR_r19_Two_Two_r19_ThirtyTwo_r19:
			(*ie.Two_Two_r19).ThirtyTwo_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(8))
			if err != nil {
				return err
			}
			(*(*ie.Two_Two_r19).ThirtyTwo_r19) = v
		case TypeII_X1_X2_CBSR_r19_Two_Two_r19_SixtyFour_r19:
			(*ie.Two_Two_r19).SixtyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(16))
			if err != nil {
				return err
			}
			(*(*ie.Two_Two_r19).SixtyFour_r19) = v
		}
	case TypeII_X1_X2_CBSR_r19_Four_One_r19:
		ie.Four_One_r19 = &struct {
			Choice         int
			TwentyFour_r19 *per.BitString
			ThirtyTwo_r19  *per.BitString
			SixtyFour_r19  *per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(typeIIX1X2CBSRR19FourOneR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Four_One_r19).Choice = int(index)
		switch index {
		case TypeII_X1_X2_CBSR_r19_Four_One_r19_TwentyFour_r19:
			(*ie.Four_One_r19).TwentyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(6))
			if err != nil {
				return err
			}
			(*(*ie.Four_One_r19).TwentyFour_r19) = v
		case TypeII_X1_X2_CBSR_r19_Four_One_r19_ThirtyTwo_r19:
			(*ie.Four_One_r19).ThirtyTwo_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(8))
			if err != nil {
				return err
			}
			(*(*ie.Four_One_r19).ThirtyTwo_r19) = v
		case TypeII_X1_X2_CBSR_r19_Four_One_r19_SixtyFour_r19:
			(*ie.Four_One_r19).SixtyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(16))
			if err != nil {
				return err
			}
			(*(*ie.Four_One_r19).SixtyFour_r19) = v
		}
	case TypeII_X1_X2_CBSR_r19_Four_Two_r19:
		ie.Four_Two_r19 = &struct {
			Choice        int
			ThirtyTwo_r19 *per.BitString
			SixtyFour_r19 *per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(typeIIX1X2CBSRR19FourTwoR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Four_Two_r19).Choice = int(index)
		switch index {
		case TypeII_X1_X2_CBSR_r19_Four_Two_r19_ThirtyTwo_r19:
			(*ie.Four_Two_r19).ThirtyTwo_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(4))
			if err != nil {
				return err
			}
			(*(*ie.Four_Two_r19).ThirtyTwo_r19) = v
		case TypeII_X1_X2_CBSR_r19_Four_Two_r19_SixtyFour_r19:
			(*ie.Four_Two_r19).SixtyFour_r19 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(8))
			if err != nil {
				return err
			}
			(*(*ie.Four_Two_r19).SixtyFour_r19) = v
		}
	default:
		return &per.DecodeError{TypeName: "TypeII-X1-X2-CBSR-r19", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
