package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_nothing uint64 = iota
	CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Two_one
	CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Two_two
	CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Four_one
	CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Three_two
	CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Six_one
	CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Four_two
	CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Eight_one
	CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Four_three
	CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Six_two
	CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Twelve_one
	CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Four_four
	CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Eight_two
	CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Sixteen_one
)

type CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction struct {
	Choice      uint64
	Two_one     uper.BitString `lb:16,ub:16,madatory`
	Two_two     uper.BitString `lb:43,ub:43,madatory`
	Four_one    uper.BitString `lb:32,ub:32,madatory`
	Three_two   uper.BitString `lb:59,ub:59,madatory`
	Six_one     uper.BitString `lb:48,ub:48,madatory`
	Four_two    uper.BitString `lb:75,ub:75,madatory`
	Eight_one   uper.BitString `lb:64,ub:64,madatory`
	Four_three  uper.BitString `lb:107,ub:107,madatory`
	Six_two     uper.BitString `lb:107,ub:107,madatory`
	Twelve_one  uper.BitString `lb:96,ub:96,madatory`
	Four_four   uper.BitString `lb:139,ub:139,madatory`
	Eight_two   uper.BitString `lb:139,ub:139,madatory`
	Sixteen_one uper.BitString `lb:128,ub:128,madatory`
}

func (ie *CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 13, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Two_one:
		if err = w.WriteBitString(ie.Two_one.Bytes, uint(ie.Two_one.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			err = utils.WrapError("Encode Two_one", err)
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Two_two:
		if err = w.WriteBitString(ie.Two_two.Bytes, uint(ie.Two_two.NumBits), &uper.Constraint{Lb: 43, Ub: 43}, false); err != nil {
			err = utils.WrapError("Encode Two_two", err)
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Four_one:
		if err = w.WriteBitString(ie.Four_one.Bytes, uint(ie.Four_one.NumBits), &uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			err = utils.WrapError("Encode Four_one", err)
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Three_two:
		if err = w.WriteBitString(ie.Three_two.Bytes, uint(ie.Three_two.NumBits), &uper.Constraint{Lb: 59, Ub: 59}, false); err != nil {
			err = utils.WrapError("Encode Three_two", err)
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Six_one:
		if err = w.WriteBitString(ie.Six_one.Bytes, uint(ie.Six_one.NumBits), &uper.Constraint{Lb: 48, Ub: 48}, false); err != nil {
			err = utils.WrapError("Encode Six_one", err)
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Four_two:
		if err = w.WriteBitString(ie.Four_two.Bytes, uint(ie.Four_two.NumBits), &uper.Constraint{Lb: 75, Ub: 75}, false); err != nil {
			err = utils.WrapError("Encode Four_two", err)
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Eight_one:
		if err = w.WriteBitString(ie.Eight_one.Bytes, uint(ie.Eight_one.NumBits), &uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			err = utils.WrapError("Encode Eight_one", err)
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Four_three:
		if err = w.WriteBitString(ie.Four_three.Bytes, uint(ie.Four_three.NumBits), &uper.Constraint{Lb: 107, Ub: 107}, false); err != nil {
			err = utils.WrapError("Encode Four_three", err)
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Six_two:
		if err = w.WriteBitString(ie.Six_two.Bytes, uint(ie.Six_two.NumBits), &uper.Constraint{Lb: 107, Ub: 107}, false); err != nil {
			err = utils.WrapError("Encode Six_two", err)
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Twelve_one:
		if err = w.WriteBitString(ie.Twelve_one.Bytes, uint(ie.Twelve_one.NumBits), &uper.Constraint{Lb: 96, Ub: 96}, false); err != nil {
			err = utils.WrapError("Encode Twelve_one", err)
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Four_four:
		if err = w.WriteBitString(ie.Four_four.Bytes, uint(ie.Four_four.NumBits), &uper.Constraint{Lb: 139, Ub: 139}, false); err != nil {
			err = utils.WrapError("Encode Four_four", err)
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Eight_two:
		if err = w.WriteBitString(ie.Eight_two.Bytes, uint(ie.Eight_two.NumBits), &uper.Constraint{Lb: 139, Ub: 139}, false); err != nil {
			err = utils.WrapError("Encode Eight_two", err)
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Sixteen_one:
		if err = w.WriteBitString(ie.Sixteen_one.Bytes, uint(ie.Sixteen_one.NumBits), &uper.Constraint{Lb: 128, Ub: 128}, false); err != nil {
			err = utils.WrapError("Encode Sixteen_one", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(13, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Two_one:
		var tmp_bs_Two_one []byte
		var n_Two_one uint
		if tmp_bs_Two_one, n_Two_one, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Two_one", err)
		}
		ie.Two_one = uper.BitString{
			Bytes:   tmp_bs_Two_one,
			NumBits: uint64(n_Two_one),
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Two_two:
		var tmp_bs_Two_two []byte
		var n_Two_two uint
		if tmp_bs_Two_two, n_Two_two, err = r.ReadBitString(&uper.Constraint{Lb: 43, Ub: 43}, false); err != nil {
			return utils.WrapError("Decode Two_two", err)
		}
		ie.Two_two = uper.BitString{
			Bytes:   tmp_bs_Two_two,
			NumBits: uint64(n_Two_two),
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Four_one:
		var tmp_bs_Four_one []byte
		var n_Four_one uint
		if tmp_bs_Four_one, n_Four_one, err = r.ReadBitString(&uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode Four_one", err)
		}
		ie.Four_one = uper.BitString{
			Bytes:   tmp_bs_Four_one,
			NumBits: uint64(n_Four_one),
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Three_two:
		var tmp_bs_Three_two []byte
		var n_Three_two uint
		if tmp_bs_Three_two, n_Three_two, err = r.ReadBitString(&uper.Constraint{Lb: 59, Ub: 59}, false); err != nil {
			return utils.WrapError("Decode Three_two", err)
		}
		ie.Three_two = uper.BitString{
			Bytes:   tmp_bs_Three_two,
			NumBits: uint64(n_Three_two),
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Six_one:
		var tmp_bs_Six_one []byte
		var n_Six_one uint
		if tmp_bs_Six_one, n_Six_one, err = r.ReadBitString(&uper.Constraint{Lb: 48, Ub: 48}, false); err != nil {
			return utils.WrapError("Decode Six_one", err)
		}
		ie.Six_one = uper.BitString{
			Bytes:   tmp_bs_Six_one,
			NumBits: uint64(n_Six_one),
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Four_two:
		var tmp_bs_Four_two []byte
		var n_Four_two uint
		if tmp_bs_Four_two, n_Four_two, err = r.ReadBitString(&uper.Constraint{Lb: 75, Ub: 75}, false); err != nil {
			return utils.WrapError("Decode Four_two", err)
		}
		ie.Four_two = uper.BitString{
			Bytes:   tmp_bs_Four_two,
			NumBits: uint64(n_Four_two),
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Eight_one:
		var tmp_bs_Eight_one []byte
		var n_Eight_one uint
		if tmp_bs_Eight_one, n_Eight_one, err = r.ReadBitString(&uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode Eight_one", err)
		}
		ie.Eight_one = uper.BitString{
			Bytes:   tmp_bs_Eight_one,
			NumBits: uint64(n_Eight_one),
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Four_three:
		var tmp_bs_Four_three []byte
		var n_Four_three uint
		if tmp_bs_Four_three, n_Four_three, err = r.ReadBitString(&uper.Constraint{Lb: 107, Ub: 107}, false); err != nil {
			return utils.WrapError("Decode Four_three", err)
		}
		ie.Four_three = uper.BitString{
			Bytes:   tmp_bs_Four_three,
			NumBits: uint64(n_Four_three),
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Six_two:
		var tmp_bs_Six_two []byte
		var n_Six_two uint
		if tmp_bs_Six_two, n_Six_two, err = r.ReadBitString(&uper.Constraint{Lb: 107, Ub: 107}, false); err != nil {
			return utils.WrapError("Decode Six_two", err)
		}
		ie.Six_two = uper.BitString{
			Bytes:   tmp_bs_Six_two,
			NumBits: uint64(n_Six_two),
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Twelve_one:
		var tmp_bs_Twelve_one []byte
		var n_Twelve_one uint
		if tmp_bs_Twelve_one, n_Twelve_one, err = r.ReadBitString(&uper.Constraint{Lb: 96, Ub: 96}, false); err != nil {
			return utils.WrapError("Decode Twelve_one", err)
		}
		ie.Twelve_one = uper.BitString{
			Bytes:   tmp_bs_Twelve_one,
			NumBits: uint64(n_Twelve_one),
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Four_four:
		var tmp_bs_Four_four []byte
		var n_Four_four uint
		if tmp_bs_Four_four, n_Four_four, err = r.ReadBitString(&uper.Constraint{Lb: 139, Ub: 139}, false); err != nil {
			return utils.WrapError("Decode Four_four", err)
		}
		ie.Four_four = uper.BitString{
			Bytes:   tmp_bs_Four_four,
			NumBits: uint64(n_Four_four),
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Eight_two:
		var tmp_bs_Eight_two []byte
		var n_Eight_two uint
		if tmp_bs_Eight_two, n_Eight_two, err = r.ReadBitString(&uper.Constraint{Lb: 139, Ub: 139}, false); err != nil {
			return utils.WrapError("Decode Eight_two", err)
		}
		ie.Eight_two = uper.BitString{
			Bytes:   tmp_bs_Eight_two,
			NumBits: uint64(n_Eight_two),
		}
	case CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction_Choice_Sixteen_one:
		var tmp_bs_Sixteen_one []byte
		var n_Sixteen_one uint
		if tmp_bs_Sixteen_one, n_Sixteen_one, err = r.ReadBitString(&uper.Constraint{Lb: 128, Ub: 128}, false); err != nil {
			return utils.WrapError("Decode Sixteen_one", err)
		}
		ie.Sixteen_one = uper.BitString{
			Bytes:   tmp_bs_Sixteen_one,
			NumBits: uint64(n_Sixteen_one),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
