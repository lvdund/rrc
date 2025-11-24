package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_nothing uint64 = iota
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Two_one_TypeI_SinglePanel_Restriction2_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Two_two_TypeI_SinglePanel_Restriction2_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Four_one_TypeI_SinglePanel_Restriction2_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Three_two_TypeI_SinglePanel_Restriction2_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Six_one_TypeI_SinglePanel_Restriction2_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Four_two_TypeI_SinglePanel_Restriction2_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Eight_one_TypeI_SinglePanel_Restriction2_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Four_three_TypeI_SinglePanel_Restriction2_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Six_two_TypeI_SinglePanel_Restriction2_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Twelve_one_TypeI_SinglePanel_Restriction2_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Four_four_TypeI_SinglePanel_Restriction2_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Eight_two_TypeI_SinglePanel_Restriction2_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Sixteen_one_TypeI_SinglePanel_Restriction2_r17
)

type CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2 struct {
	Choice                                         uint64
	Two_one_TypeI_SinglePanel_Restriction2_r17     uper.BitString `lb:8,ub:8,madatory`
	Two_two_TypeI_SinglePanel_Restriction2_r17     uper.BitString `lb:64,ub:64,madatory`
	Four_one_TypeI_SinglePanel_Restriction2_r17    uper.BitString `lb:16,ub:16,madatory`
	Three_two_TypeI_SinglePanel_Restriction2_r17   uper.BitString `lb:96,ub:96,madatory`
	Six_one_TypeI_SinglePanel_Restriction2_r17     uper.BitString `lb:24,ub:24,madatory`
	Four_two_TypeI_SinglePanel_Restriction2_r17    uper.BitString `lb:128,ub:128,madatory`
	Eight_one_TypeI_SinglePanel_Restriction2_r17   uper.BitString `lb:32,ub:32,madatory`
	Four_three_TypeI_SinglePanel_Restriction2_r17  uper.BitString `lb:192,ub:192,madatory`
	Six_two_TypeI_SinglePanel_Restriction2_r17     uper.BitString `lb:192,ub:192,madatory`
	Twelve_one_TypeI_SinglePanel_Restriction2_r17  uper.BitString `lb:48,ub:48,madatory`
	Four_four_TypeI_SinglePanel_Restriction2_r17   uper.BitString `lb:256,ub:256,madatory`
	Eight_two_TypeI_SinglePanel_Restriction2_r17   uper.BitString `lb:256,ub:256,madatory`
	Sixteen_one_TypeI_SinglePanel_Restriction2_r17 uper.BitString `lb:64,ub:64,madatory`
}

func (ie *CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 13, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Two_one_TypeI_SinglePanel_Restriction2_r17:
		if err = w.WriteBitString(ie.Two_one_TypeI_SinglePanel_Restriction2_r17.Bytes, uint(ie.Two_one_TypeI_SinglePanel_Restriction2_r17.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode Two_one_TypeI_SinglePanel_Restriction2_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Two_two_TypeI_SinglePanel_Restriction2_r17:
		if err = w.WriteBitString(ie.Two_two_TypeI_SinglePanel_Restriction2_r17.Bytes, uint(ie.Two_two_TypeI_SinglePanel_Restriction2_r17.NumBits), &uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			err = utils.WrapError("Encode Two_two_TypeI_SinglePanel_Restriction2_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Four_one_TypeI_SinglePanel_Restriction2_r17:
		if err = w.WriteBitString(ie.Four_one_TypeI_SinglePanel_Restriction2_r17.Bytes, uint(ie.Four_one_TypeI_SinglePanel_Restriction2_r17.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			err = utils.WrapError("Encode Four_one_TypeI_SinglePanel_Restriction2_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Three_two_TypeI_SinglePanel_Restriction2_r17:
		if err = w.WriteBitString(ie.Three_two_TypeI_SinglePanel_Restriction2_r17.Bytes, uint(ie.Three_two_TypeI_SinglePanel_Restriction2_r17.NumBits), &uper.Constraint{Lb: 96, Ub: 96}, false); err != nil {
			err = utils.WrapError("Encode Three_two_TypeI_SinglePanel_Restriction2_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Six_one_TypeI_SinglePanel_Restriction2_r17:
		if err = w.WriteBitString(ie.Six_one_TypeI_SinglePanel_Restriction2_r17.Bytes, uint(ie.Six_one_TypeI_SinglePanel_Restriction2_r17.NumBits), &uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
			err = utils.WrapError("Encode Six_one_TypeI_SinglePanel_Restriction2_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Four_two_TypeI_SinglePanel_Restriction2_r17:
		if err = w.WriteBitString(ie.Four_two_TypeI_SinglePanel_Restriction2_r17.Bytes, uint(ie.Four_two_TypeI_SinglePanel_Restriction2_r17.NumBits), &uper.Constraint{Lb: 128, Ub: 128}, false); err != nil {
			err = utils.WrapError("Encode Four_two_TypeI_SinglePanel_Restriction2_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Eight_one_TypeI_SinglePanel_Restriction2_r17:
		if err = w.WriteBitString(ie.Eight_one_TypeI_SinglePanel_Restriction2_r17.Bytes, uint(ie.Eight_one_TypeI_SinglePanel_Restriction2_r17.NumBits), &uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			err = utils.WrapError("Encode Eight_one_TypeI_SinglePanel_Restriction2_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Four_three_TypeI_SinglePanel_Restriction2_r17:
		if err = w.WriteBitString(ie.Four_three_TypeI_SinglePanel_Restriction2_r17.Bytes, uint(ie.Four_three_TypeI_SinglePanel_Restriction2_r17.NumBits), &uper.Constraint{Lb: 192, Ub: 192}, false); err != nil {
			err = utils.WrapError("Encode Four_three_TypeI_SinglePanel_Restriction2_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Six_two_TypeI_SinglePanel_Restriction2_r17:
		if err = w.WriteBitString(ie.Six_two_TypeI_SinglePanel_Restriction2_r17.Bytes, uint(ie.Six_two_TypeI_SinglePanel_Restriction2_r17.NumBits), &uper.Constraint{Lb: 192, Ub: 192}, false); err != nil {
			err = utils.WrapError("Encode Six_two_TypeI_SinglePanel_Restriction2_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Twelve_one_TypeI_SinglePanel_Restriction2_r17:
		if err = w.WriteBitString(ie.Twelve_one_TypeI_SinglePanel_Restriction2_r17.Bytes, uint(ie.Twelve_one_TypeI_SinglePanel_Restriction2_r17.NumBits), &uper.Constraint{Lb: 48, Ub: 48}, false); err != nil {
			err = utils.WrapError("Encode Twelve_one_TypeI_SinglePanel_Restriction2_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Four_four_TypeI_SinglePanel_Restriction2_r17:
		if err = w.WriteBitString(ie.Four_four_TypeI_SinglePanel_Restriction2_r17.Bytes, uint(ie.Four_four_TypeI_SinglePanel_Restriction2_r17.NumBits), &uper.Constraint{Lb: 256, Ub: 256}, false); err != nil {
			err = utils.WrapError("Encode Four_four_TypeI_SinglePanel_Restriction2_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Eight_two_TypeI_SinglePanel_Restriction2_r17:
		if err = w.WriteBitString(ie.Eight_two_TypeI_SinglePanel_Restriction2_r17.Bytes, uint(ie.Eight_two_TypeI_SinglePanel_Restriction2_r17.NumBits), &uper.Constraint{Lb: 256, Ub: 256}, false); err != nil {
			err = utils.WrapError("Encode Eight_two_TypeI_SinglePanel_Restriction2_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Sixteen_one_TypeI_SinglePanel_Restriction2_r17:
		if err = w.WriteBitString(ie.Sixteen_one_TypeI_SinglePanel_Restriction2_r17.Bytes, uint(ie.Sixteen_one_TypeI_SinglePanel_Restriction2_r17.NumBits), &uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			err = utils.WrapError("Encode Sixteen_one_TypeI_SinglePanel_Restriction2_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(13, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Two_one_TypeI_SinglePanel_Restriction2_r17:
		var tmp_bs_Two_one_TypeI_SinglePanel_Restriction2_r17 []byte
		var n_Two_one_TypeI_SinglePanel_Restriction2_r17 uint
		if tmp_bs_Two_one_TypeI_SinglePanel_Restriction2_r17, n_Two_one_TypeI_SinglePanel_Restriction2_r17, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Two_one_TypeI_SinglePanel_Restriction2_r17", err)
		}
		ie.Two_one_TypeI_SinglePanel_Restriction2_r17 = uper.BitString{
			Bytes:   tmp_bs_Two_one_TypeI_SinglePanel_Restriction2_r17,
			NumBits: uint64(n_Two_one_TypeI_SinglePanel_Restriction2_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Two_two_TypeI_SinglePanel_Restriction2_r17:
		var tmp_bs_Two_two_TypeI_SinglePanel_Restriction2_r17 []byte
		var n_Two_two_TypeI_SinglePanel_Restriction2_r17 uint
		if tmp_bs_Two_two_TypeI_SinglePanel_Restriction2_r17, n_Two_two_TypeI_SinglePanel_Restriction2_r17, err = r.ReadBitString(&uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode Two_two_TypeI_SinglePanel_Restriction2_r17", err)
		}
		ie.Two_two_TypeI_SinglePanel_Restriction2_r17 = uper.BitString{
			Bytes:   tmp_bs_Two_two_TypeI_SinglePanel_Restriction2_r17,
			NumBits: uint64(n_Two_two_TypeI_SinglePanel_Restriction2_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Four_one_TypeI_SinglePanel_Restriction2_r17:
		var tmp_bs_Four_one_TypeI_SinglePanel_Restriction2_r17 []byte
		var n_Four_one_TypeI_SinglePanel_Restriction2_r17 uint
		if tmp_bs_Four_one_TypeI_SinglePanel_Restriction2_r17, n_Four_one_TypeI_SinglePanel_Restriction2_r17, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Four_one_TypeI_SinglePanel_Restriction2_r17", err)
		}
		ie.Four_one_TypeI_SinglePanel_Restriction2_r17 = uper.BitString{
			Bytes:   tmp_bs_Four_one_TypeI_SinglePanel_Restriction2_r17,
			NumBits: uint64(n_Four_one_TypeI_SinglePanel_Restriction2_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Three_two_TypeI_SinglePanel_Restriction2_r17:
		var tmp_bs_Three_two_TypeI_SinglePanel_Restriction2_r17 []byte
		var n_Three_two_TypeI_SinglePanel_Restriction2_r17 uint
		if tmp_bs_Three_two_TypeI_SinglePanel_Restriction2_r17, n_Three_two_TypeI_SinglePanel_Restriction2_r17, err = r.ReadBitString(&uper.Constraint{Lb: 96, Ub: 96}, false); err != nil {
			return utils.WrapError("Decode Three_two_TypeI_SinglePanel_Restriction2_r17", err)
		}
		ie.Three_two_TypeI_SinglePanel_Restriction2_r17 = uper.BitString{
			Bytes:   tmp_bs_Three_two_TypeI_SinglePanel_Restriction2_r17,
			NumBits: uint64(n_Three_two_TypeI_SinglePanel_Restriction2_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Six_one_TypeI_SinglePanel_Restriction2_r17:
		var tmp_bs_Six_one_TypeI_SinglePanel_Restriction2_r17 []byte
		var n_Six_one_TypeI_SinglePanel_Restriction2_r17 uint
		if tmp_bs_Six_one_TypeI_SinglePanel_Restriction2_r17, n_Six_one_TypeI_SinglePanel_Restriction2_r17, err = r.ReadBitString(&uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
			return utils.WrapError("Decode Six_one_TypeI_SinglePanel_Restriction2_r17", err)
		}
		ie.Six_one_TypeI_SinglePanel_Restriction2_r17 = uper.BitString{
			Bytes:   tmp_bs_Six_one_TypeI_SinglePanel_Restriction2_r17,
			NumBits: uint64(n_Six_one_TypeI_SinglePanel_Restriction2_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Four_two_TypeI_SinglePanel_Restriction2_r17:
		var tmp_bs_Four_two_TypeI_SinglePanel_Restriction2_r17 []byte
		var n_Four_two_TypeI_SinglePanel_Restriction2_r17 uint
		if tmp_bs_Four_two_TypeI_SinglePanel_Restriction2_r17, n_Four_two_TypeI_SinglePanel_Restriction2_r17, err = r.ReadBitString(&uper.Constraint{Lb: 128, Ub: 128}, false); err != nil {
			return utils.WrapError("Decode Four_two_TypeI_SinglePanel_Restriction2_r17", err)
		}
		ie.Four_two_TypeI_SinglePanel_Restriction2_r17 = uper.BitString{
			Bytes:   tmp_bs_Four_two_TypeI_SinglePanel_Restriction2_r17,
			NumBits: uint64(n_Four_two_TypeI_SinglePanel_Restriction2_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Eight_one_TypeI_SinglePanel_Restriction2_r17:
		var tmp_bs_Eight_one_TypeI_SinglePanel_Restriction2_r17 []byte
		var n_Eight_one_TypeI_SinglePanel_Restriction2_r17 uint
		if tmp_bs_Eight_one_TypeI_SinglePanel_Restriction2_r17, n_Eight_one_TypeI_SinglePanel_Restriction2_r17, err = r.ReadBitString(&uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode Eight_one_TypeI_SinglePanel_Restriction2_r17", err)
		}
		ie.Eight_one_TypeI_SinglePanel_Restriction2_r17 = uper.BitString{
			Bytes:   tmp_bs_Eight_one_TypeI_SinglePanel_Restriction2_r17,
			NumBits: uint64(n_Eight_one_TypeI_SinglePanel_Restriction2_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Four_three_TypeI_SinglePanel_Restriction2_r17:
		var tmp_bs_Four_three_TypeI_SinglePanel_Restriction2_r17 []byte
		var n_Four_three_TypeI_SinglePanel_Restriction2_r17 uint
		if tmp_bs_Four_three_TypeI_SinglePanel_Restriction2_r17, n_Four_three_TypeI_SinglePanel_Restriction2_r17, err = r.ReadBitString(&uper.Constraint{Lb: 192, Ub: 192}, false); err != nil {
			return utils.WrapError("Decode Four_three_TypeI_SinglePanel_Restriction2_r17", err)
		}
		ie.Four_three_TypeI_SinglePanel_Restriction2_r17 = uper.BitString{
			Bytes:   tmp_bs_Four_three_TypeI_SinglePanel_Restriction2_r17,
			NumBits: uint64(n_Four_three_TypeI_SinglePanel_Restriction2_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Six_two_TypeI_SinglePanel_Restriction2_r17:
		var tmp_bs_Six_two_TypeI_SinglePanel_Restriction2_r17 []byte
		var n_Six_two_TypeI_SinglePanel_Restriction2_r17 uint
		if tmp_bs_Six_two_TypeI_SinglePanel_Restriction2_r17, n_Six_two_TypeI_SinglePanel_Restriction2_r17, err = r.ReadBitString(&uper.Constraint{Lb: 192, Ub: 192}, false); err != nil {
			return utils.WrapError("Decode Six_two_TypeI_SinglePanel_Restriction2_r17", err)
		}
		ie.Six_two_TypeI_SinglePanel_Restriction2_r17 = uper.BitString{
			Bytes:   tmp_bs_Six_two_TypeI_SinglePanel_Restriction2_r17,
			NumBits: uint64(n_Six_two_TypeI_SinglePanel_Restriction2_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Twelve_one_TypeI_SinglePanel_Restriction2_r17:
		var tmp_bs_Twelve_one_TypeI_SinglePanel_Restriction2_r17 []byte
		var n_Twelve_one_TypeI_SinglePanel_Restriction2_r17 uint
		if tmp_bs_Twelve_one_TypeI_SinglePanel_Restriction2_r17, n_Twelve_one_TypeI_SinglePanel_Restriction2_r17, err = r.ReadBitString(&uper.Constraint{Lb: 48, Ub: 48}, false); err != nil {
			return utils.WrapError("Decode Twelve_one_TypeI_SinglePanel_Restriction2_r17", err)
		}
		ie.Twelve_one_TypeI_SinglePanel_Restriction2_r17 = uper.BitString{
			Bytes:   tmp_bs_Twelve_one_TypeI_SinglePanel_Restriction2_r17,
			NumBits: uint64(n_Twelve_one_TypeI_SinglePanel_Restriction2_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Four_four_TypeI_SinglePanel_Restriction2_r17:
		var tmp_bs_Four_four_TypeI_SinglePanel_Restriction2_r17 []byte
		var n_Four_four_TypeI_SinglePanel_Restriction2_r17 uint
		if tmp_bs_Four_four_TypeI_SinglePanel_Restriction2_r17, n_Four_four_TypeI_SinglePanel_Restriction2_r17, err = r.ReadBitString(&uper.Constraint{Lb: 256, Ub: 256}, false); err != nil {
			return utils.WrapError("Decode Four_four_TypeI_SinglePanel_Restriction2_r17", err)
		}
		ie.Four_four_TypeI_SinglePanel_Restriction2_r17 = uper.BitString{
			Bytes:   tmp_bs_Four_four_TypeI_SinglePanel_Restriction2_r17,
			NumBits: uint64(n_Four_four_TypeI_SinglePanel_Restriction2_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Eight_two_TypeI_SinglePanel_Restriction2_r17:
		var tmp_bs_Eight_two_TypeI_SinglePanel_Restriction2_r17 []byte
		var n_Eight_two_TypeI_SinglePanel_Restriction2_r17 uint
		if tmp_bs_Eight_two_TypeI_SinglePanel_Restriction2_r17, n_Eight_two_TypeI_SinglePanel_Restriction2_r17, err = r.ReadBitString(&uper.Constraint{Lb: 256, Ub: 256}, false); err != nil {
			return utils.WrapError("Decode Eight_two_TypeI_SinglePanel_Restriction2_r17", err)
		}
		ie.Eight_two_TypeI_SinglePanel_Restriction2_r17 = uper.BitString{
			Bytes:   tmp_bs_Eight_two_TypeI_SinglePanel_Restriction2_r17,
			NumBits: uint64(n_Eight_two_TypeI_SinglePanel_Restriction2_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_Sixteen_one_TypeI_SinglePanel_Restriction2_r17:
		var tmp_bs_Sixteen_one_TypeI_SinglePanel_Restriction2_r17 []byte
		var n_Sixteen_one_TypeI_SinglePanel_Restriction2_r17 uint
		if tmp_bs_Sixteen_one_TypeI_SinglePanel_Restriction2_r17, n_Sixteen_one_TypeI_SinglePanel_Restriction2_r17, err = r.ReadBitString(&uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode Sixteen_one_TypeI_SinglePanel_Restriction2_r17", err)
		}
		ie.Sixteen_one_TypeI_SinglePanel_Restriction2_r17 = uper.BitString{
			Bytes:   tmp_bs_Sixteen_one_TypeI_SinglePanel_Restriction2_r17,
			NumBits: uint64(n_Sixteen_one_TypeI_SinglePanel_Restriction2_r17),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
