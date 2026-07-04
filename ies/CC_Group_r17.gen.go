// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CC-Group-r17 (line 16392).

var cCGroupR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "servCellIndexLower-r17"},
		{Name: "servCellIndexHigher-r17", Optional: true},
		{Name: "defaultDC-Location-r17"},
		{Name: "offsetToDefault-r17", Optional: true},
	},
}

var cC_Group_r17OffsetToDefaultR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "offsetValue"},
		{Name: "offsetlist"},
	},
}

const (
	CC_Group_r17_OffsetToDefault_r17_OffsetValue = 0
	CC_Group_r17_OffsetToDefault_r17_Offsetlist  = 1
)

var cCGroupR17OffsetToDefaultR17OffsetlistConstraints = per.SizeRange(1, common.MaxNrofReqComDC_Location_r17)

type CC_Group_r17 struct {
	ServCellIndexLower_r17  ServCellIndex
	ServCellIndexHigher_r17 *ServCellIndex
	DefaultDC_Location_r17  DefaultDC_Location_r17
	OffsetToDefault_r17     *struct {
		Choice      int
		OffsetValue *OffsetValue_r17
		Offsetlist  *[]OffsetValue_r17
	}
}

func (ie *CC_Group_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cCGroupR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ServCellIndexHigher_r17 != nil, ie.OffsetToDefault_r17 != nil}); err != nil {
		return err
	}

	// 2. servCellIndexLower-r17: ref
	{
		if err := ie.ServCellIndexLower_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. servCellIndexHigher-r17: ref
	{
		if ie.ServCellIndexHigher_r17 != nil {
			if err := ie.ServCellIndexHigher_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. defaultDC-Location-r17: ref
	{
		if err := ie.DefaultDC_Location_r17.Encode(e); err != nil {
			return err
		}
	}

	// 5. offsetToDefault-r17: choice
	{
		if ie.OffsetToDefault_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(cC_Group_r17OffsetToDefaultR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.OffsetToDefault_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.OffsetToDefault_r17).Choice {
			case CC_Group_r17_OffsetToDefault_r17_OffsetValue:
				if err := (*ie.OffsetToDefault_r17).OffsetValue.Encode(e); err != nil {
					return err
				}
			case CC_Group_r17_OffsetToDefault_r17_Offsetlist:
				seqOf := e.NewSequenceOfEncoder(cCGroupR17OffsetToDefaultR17OffsetlistConstraints)
				if err := seqOf.EncodeLength(int64(len((*(*ie.OffsetToDefault_r17).Offsetlist)))); err != nil {
					return err
				}
				for i := range *(*ie.OffsetToDefault_r17).Offsetlist {
					if err := (*(*ie.OffsetToDefault_r17).Offsetlist)[i].Encode(e); err != nil {
						return err
					}
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.OffsetToDefault_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *CC_Group_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cCGroupR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. servCellIndexLower-r17: ref
	{
		if err := ie.ServCellIndexLower_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. servCellIndexHigher-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.ServCellIndexHigher_r17 = new(ServCellIndex)
			if err := ie.ServCellIndexHigher_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. defaultDC-Location-r17: ref
	{
		if err := ie.DefaultDC_Location_r17.Decode(d); err != nil {
			return err
		}
	}

	// 5. offsetToDefault-r17: choice
	{
		if seq.IsComponentPresent(3) {
			ie.OffsetToDefault_r17 = &struct {
				Choice      int
				OffsetValue *OffsetValue_r17
				Offsetlist  *[]OffsetValue_r17
			}{}
			choiceDec := d.NewChoiceDecoder(cC_Group_r17OffsetToDefaultR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.OffsetToDefault_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case CC_Group_r17_OffsetToDefault_r17_OffsetValue:
				(*ie.OffsetToDefault_r17).OffsetValue = new(OffsetValue_r17)
				if err := (*ie.OffsetToDefault_r17).OffsetValue.Decode(d); err != nil {
					return err
				}
			case CC_Group_r17_OffsetToDefault_r17_Offsetlist:
				(*ie.OffsetToDefault_r17).Offsetlist = new([]OffsetValue_r17)
				seqOf := d.NewSequenceOfDecoder(cCGroupR17OffsetToDefaultR17OffsetlistConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.OffsetToDefault_r17).Offsetlist) = make([]OffsetValue_r17, n)
				for i := int64(0); i < n; i++ {
					if err := (*(*ie.OffsetToDefault_r17).Offsetlist)[i].Decode(d); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
