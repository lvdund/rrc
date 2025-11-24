package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_nothing uint64 = iota
	CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Two_two_one_TypeI_MultiPanel_Restriction
	CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Two_four_one_TypeI_MultiPanel_Restriction
	CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Four_two_one_TypeI_MultiPanel_Restriction
	CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Two_two_two_TypeI_MultiPanel_Restriction
	CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Two_eight_one_TypeI_MultiPanel_Restriction
	CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Four_four_one_TypeI_MultiPanel_Restriction
	CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Two_four_two_TypeI_MultiPanel_Restriction
	CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Four_two_two_TypeI_MultiPanel_Restriction
)

type CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2 struct {
	Choice                                     uint64
	Two_two_one_TypeI_MultiPanel_Restriction   uper.BitString `lb:8,ub:8,madatory`
	Two_four_one_TypeI_MultiPanel_Restriction  uper.BitString `lb:16,ub:16,madatory`
	Four_two_one_TypeI_MultiPanel_Restriction  uper.BitString `lb:8,ub:8,madatory`
	Two_two_two_TypeI_MultiPanel_Restriction   uper.BitString `lb:64,ub:64,madatory`
	Two_eight_one_TypeI_MultiPanel_Restriction uper.BitString `lb:32,ub:32,madatory`
	Four_four_one_TypeI_MultiPanel_Restriction uper.BitString `lb:16,ub:16,madatory`
	Two_four_two_TypeI_MultiPanel_Restriction  uper.BitString `lb:128,ub:128,madatory`
	Four_two_two_TypeI_MultiPanel_Restriction  uper.BitString `lb:64,ub:64,madatory`
}

func (ie *CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Two_two_one_TypeI_MultiPanel_Restriction:
		if err = w.WriteBitString(ie.Two_two_one_TypeI_MultiPanel_Restriction.Bytes, uint(ie.Two_two_one_TypeI_MultiPanel_Restriction.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode Two_two_one_TypeI_MultiPanel_Restriction", err)
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Two_four_one_TypeI_MultiPanel_Restriction:
		if err = w.WriteBitString(ie.Two_four_one_TypeI_MultiPanel_Restriction.Bytes, uint(ie.Two_four_one_TypeI_MultiPanel_Restriction.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			err = utils.WrapError("Encode Two_four_one_TypeI_MultiPanel_Restriction", err)
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Four_two_one_TypeI_MultiPanel_Restriction:
		if err = w.WriteBitString(ie.Four_two_one_TypeI_MultiPanel_Restriction.Bytes, uint(ie.Four_two_one_TypeI_MultiPanel_Restriction.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode Four_two_one_TypeI_MultiPanel_Restriction", err)
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Two_two_two_TypeI_MultiPanel_Restriction:
		if err = w.WriteBitString(ie.Two_two_two_TypeI_MultiPanel_Restriction.Bytes, uint(ie.Two_two_two_TypeI_MultiPanel_Restriction.NumBits), &uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			err = utils.WrapError("Encode Two_two_two_TypeI_MultiPanel_Restriction", err)
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Two_eight_one_TypeI_MultiPanel_Restriction:
		if err = w.WriteBitString(ie.Two_eight_one_TypeI_MultiPanel_Restriction.Bytes, uint(ie.Two_eight_one_TypeI_MultiPanel_Restriction.NumBits), &uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			err = utils.WrapError("Encode Two_eight_one_TypeI_MultiPanel_Restriction", err)
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Four_four_one_TypeI_MultiPanel_Restriction:
		if err = w.WriteBitString(ie.Four_four_one_TypeI_MultiPanel_Restriction.Bytes, uint(ie.Four_four_one_TypeI_MultiPanel_Restriction.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			err = utils.WrapError("Encode Four_four_one_TypeI_MultiPanel_Restriction", err)
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Two_four_two_TypeI_MultiPanel_Restriction:
		if err = w.WriteBitString(ie.Two_four_two_TypeI_MultiPanel_Restriction.Bytes, uint(ie.Two_four_two_TypeI_MultiPanel_Restriction.NumBits), &uper.Constraint{Lb: 128, Ub: 128}, false); err != nil {
			err = utils.WrapError("Encode Two_four_two_TypeI_MultiPanel_Restriction", err)
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Four_two_two_TypeI_MultiPanel_Restriction:
		if err = w.WriteBitString(ie.Four_two_two_TypeI_MultiPanel_Restriction.Bytes, uint(ie.Four_two_two_TypeI_MultiPanel_Restriction.NumBits), &uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			err = utils.WrapError("Encode Four_two_two_TypeI_MultiPanel_Restriction", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Two_two_one_TypeI_MultiPanel_Restriction:
		var tmp_bs_Two_two_one_TypeI_MultiPanel_Restriction []byte
		var n_Two_two_one_TypeI_MultiPanel_Restriction uint
		if tmp_bs_Two_two_one_TypeI_MultiPanel_Restriction, n_Two_two_one_TypeI_MultiPanel_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Two_two_one_TypeI_MultiPanel_Restriction", err)
		}
		ie.Two_two_one_TypeI_MultiPanel_Restriction = uper.BitString{
			Bytes:   tmp_bs_Two_two_one_TypeI_MultiPanel_Restriction,
			NumBits: uint64(n_Two_two_one_TypeI_MultiPanel_Restriction),
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Two_four_one_TypeI_MultiPanel_Restriction:
		var tmp_bs_Two_four_one_TypeI_MultiPanel_Restriction []byte
		var n_Two_four_one_TypeI_MultiPanel_Restriction uint
		if tmp_bs_Two_four_one_TypeI_MultiPanel_Restriction, n_Two_four_one_TypeI_MultiPanel_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Two_four_one_TypeI_MultiPanel_Restriction", err)
		}
		ie.Two_four_one_TypeI_MultiPanel_Restriction = uper.BitString{
			Bytes:   tmp_bs_Two_four_one_TypeI_MultiPanel_Restriction,
			NumBits: uint64(n_Two_four_one_TypeI_MultiPanel_Restriction),
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Four_two_one_TypeI_MultiPanel_Restriction:
		var tmp_bs_Four_two_one_TypeI_MultiPanel_Restriction []byte
		var n_Four_two_one_TypeI_MultiPanel_Restriction uint
		if tmp_bs_Four_two_one_TypeI_MultiPanel_Restriction, n_Four_two_one_TypeI_MultiPanel_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Four_two_one_TypeI_MultiPanel_Restriction", err)
		}
		ie.Four_two_one_TypeI_MultiPanel_Restriction = uper.BitString{
			Bytes:   tmp_bs_Four_two_one_TypeI_MultiPanel_Restriction,
			NumBits: uint64(n_Four_two_one_TypeI_MultiPanel_Restriction),
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Two_two_two_TypeI_MultiPanel_Restriction:
		var tmp_bs_Two_two_two_TypeI_MultiPanel_Restriction []byte
		var n_Two_two_two_TypeI_MultiPanel_Restriction uint
		if tmp_bs_Two_two_two_TypeI_MultiPanel_Restriction, n_Two_two_two_TypeI_MultiPanel_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode Two_two_two_TypeI_MultiPanel_Restriction", err)
		}
		ie.Two_two_two_TypeI_MultiPanel_Restriction = uper.BitString{
			Bytes:   tmp_bs_Two_two_two_TypeI_MultiPanel_Restriction,
			NumBits: uint64(n_Two_two_two_TypeI_MultiPanel_Restriction),
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Two_eight_one_TypeI_MultiPanel_Restriction:
		var tmp_bs_Two_eight_one_TypeI_MultiPanel_Restriction []byte
		var n_Two_eight_one_TypeI_MultiPanel_Restriction uint
		if tmp_bs_Two_eight_one_TypeI_MultiPanel_Restriction, n_Two_eight_one_TypeI_MultiPanel_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode Two_eight_one_TypeI_MultiPanel_Restriction", err)
		}
		ie.Two_eight_one_TypeI_MultiPanel_Restriction = uper.BitString{
			Bytes:   tmp_bs_Two_eight_one_TypeI_MultiPanel_Restriction,
			NumBits: uint64(n_Two_eight_one_TypeI_MultiPanel_Restriction),
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Four_four_one_TypeI_MultiPanel_Restriction:
		var tmp_bs_Four_four_one_TypeI_MultiPanel_Restriction []byte
		var n_Four_four_one_TypeI_MultiPanel_Restriction uint
		if tmp_bs_Four_four_one_TypeI_MultiPanel_Restriction, n_Four_four_one_TypeI_MultiPanel_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Four_four_one_TypeI_MultiPanel_Restriction", err)
		}
		ie.Four_four_one_TypeI_MultiPanel_Restriction = uper.BitString{
			Bytes:   tmp_bs_Four_four_one_TypeI_MultiPanel_Restriction,
			NumBits: uint64(n_Four_four_one_TypeI_MultiPanel_Restriction),
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Two_four_two_TypeI_MultiPanel_Restriction:
		var tmp_bs_Two_four_two_TypeI_MultiPanel_Restriction []byte
		var n_Two_four_two_TypeI_MultiPanel_Restriction uint
		if tmp_bs_Two_four_two_TypeI_MultiPanel_Restriction, n_Two_four_two_TypeI_MultiPanel_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 128, Ub: 128}, false); err != nil {
			return utils.WrapError("Decode Two_four_two_TypeI_MultiPanel_Restriction", err)
		}
		ie.Two_four_two_TypeI_MultiPanel_Restriction = uper.BitString{
			Bytes:   tmp_bs_Two_four_two_TypeI_MultiPanel_Restriction,
			NumBits: uint64(n_Two_four_two_TypeI_MultiPanel_Restriction),
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_Four_two_two_TypeI_MultiPanel_Restriction:
		var tmp_bs_Four_two_two_TypeI_MultiPanel_Restriction []byte
		var n_Four_two_two_TypeI_MultiPanel_Restriction uint
		if tmp_bs_Four_two_two_TypeI_MultiPanel_Restriction, n_Four_two_two_TypeI_MultiPanel_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode Four_two_two_TypeI_MultiPanel_Restriction", err)
		}
		ie.Four_two_two_TypeI_MultiPanel_Restriction = uper.BitString{
			Bytes:   tmp_bs_Four_two_two_TypeI_MultiPanel_Restriction,
			NumBits: uint64(n_Four_two_two_TypeI_MultiPanel_Restriction),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
