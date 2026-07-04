// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UAC-BarringPerPLMN (line 16209).

var uACBarringPerPLMNConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-IdentityIndex"},
		{Name: "uac-ACBarringListType", Optional: true},
	},
}

var uACBarringPerPLMNPlmnIdentityIndexConstraints = per.Constrained(1, common.MaxPLMN)

var uAC_BarringPerPLMNUacACBarringListTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "uac-ImplicitACBarringList"},
		{Name: "uac-ExplicitACBarringList"},
	},
}

const (
	UAC_BarringPerPLMN_Uac_ACBarringListType_Uac_ImplicitACBarringList = 0
	UAC_BarringPerPLMN_Uac_ACBarringListType_Uac_ExplicitACBarringList = 1
)

var uACBarringPerPLMNUacACBarringListTypeUacImplicitACBarringListConstraints = per.FixedSize(common.MaxAccessCat_1)

type UAC_BarringPerPLMN struct {
	Plmn_IdentityIndex    int64
	Uac_ACBarringListType *struct {
		Choice                    int
		Uac_ImplicitACBarringList *[]UAC_BarringInfoSetIndex
		Uac_ExplicitACBarringList *UAC_BarringPerCatList
	}
}

func (ie *UAC_BarringPerPLMN) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uACBarringPerPLMNConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Uac_ACBarringListType != nil}); err != nil {
		return err
	}

	// 2. plmn-IdentityIndex: integer
	{
		if err := e.EncodeInteger(ie.Plmn_IdentityIndex, uACBarringPerPLMNPlmnIdentityIndexConstraints); err != nil {
			return err
		}
	}

	// 3. uac-ACBarringListType: choice
	{
		if ie.Uac_ACBarringListType != nil {
			choiceEnc := e.NewChoiceEncoder(uAC_BarringPerPLMNUacACBarringListTypeConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Uac_ACBarringListType).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Uac_ACBarringListType).Choice {
			case UAC_BarringPerPLMN_Uac_ACBarringListType_Uac_ImplicitACBarringList:
				seqOf := e.NewSequenceOfEncoder(uACBarringPerPLMNUacACBarringListTypeUacImplicitACBarringListConstraints)
				if err := seqOf.EncodeLength(int64(len((*(*ie.Uac_ACBarringListType).Uac_ImplicitACBarringList)))); err != nil {
					return err
				}
				for i := range *(*ie.Uac_ACBarringListType).Uac_ImplicitACBarringList {
					if err := (*(*ie.Uac_ACBarringListType).Uac_ImplicitACBarringList)[i].Encode(e); err != nil {
						return err
					}
				}
			case UAC_BarringPerPLMN_Uac_ACBarringListType_Uac_ExplicitACBarringList:
				if err := (*ie.Uac_ACBarringListType).Uac_ExplicitACBarringList.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Uac_ACBarringListType).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *UAC_BarringPerPLMN) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uACBarringPerPLMNConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. plmn-IdentityIndex: integer
	{
		v0, err := d.DecodeInteger(uACBarringPerPLMNPlmnIdentityIndexConstraints)
		if err != nil {
			return err
		}
		ie.Plmn_IdentityIndex = v0
	}

	// 3. uac-ACBarringListType: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Uac_ACBarringListType = &struct {
				Choice                    int
				Uac_ImplicitACBarringList *[]UAC_BarringInfoSetIndex
				Uac_ExplicitACBarringList *UAC_BarringPerCatList
			}{}
			choiceDec := d.NewChoiceDecoder(uAC_BarringPerPLMNUacACBarringListTypeConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Uac_ACBarringListType).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case UAC_BarringPerPLMN_Uac_ACBarringListType_Uac_ImplicitACBarringList:
				(*ie.Uac_ACBarringListType).Uac_ImplicitACBarringList = new([]UAC_BarringInfoSetIndex)
				seqOf := d.NewSequenceOfDecoder(uACBarringPerPLMNUacACBarringListTypeUacImplicitACBarringListConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.Uac_ACBarringListType).Uac_ImplicitACBarringList) = make([]UAC_BarringInfoSetIndex, n)
				for i := int64(0); i < n; i++ {
					if err := (*(*ie.Uac_ACBarringListType).Uac_ImplicitACBarringList)[i].Decode(d); err != nil {
						return err
					}
				}
			case UAC_BarringPerPLMN_Uac_ACBarringListType_Uac_ExplicitACBarringList:
				(*ie.Uac_ACBarringListType).Uac_ExplicitACBarringList = new(UAC_BarringPerCatList)
				if err := (*ie.Uac_ACBarringListType).Uac_ExplicitACBarringList.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
