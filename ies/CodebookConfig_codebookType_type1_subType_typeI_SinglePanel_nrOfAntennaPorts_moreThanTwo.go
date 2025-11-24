package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_moreThanTwo struct {
	N1_n2                                          CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_moreThanTwo_n1_n2 `lb:8,ub:8,madatory`
	TypeI_SinglePanel_codebookSubsetRestriction_i2 *uper.BitString                                                                                `lb:16,ub:16,optional`
}

func (ie *CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_moreThanTwo) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.TypeI_SinglePanel_codebookSubsetRestriction_i2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.N1_n2.Encode(w); err != nil {
		return utils.WrapError("Encode N1_n2", err)
	}
	if ie.TypeI_SinglePanel_codebookSubsetRestriction_i2 != nil {
		if err = w.WriteBitString(ie.TypeI_SinglePanel_codebookSubsetRestriction_i2.Bytes, uint(ie.TypeI_SinglePanel_codebookSubsetRestriction_i2.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode TypeI_SinglePanel_codebookSubsetRestriction_i2", err)
		}
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_moreThanTwo) Decode(r *uper.UperReader) error {
	var err error
	var TypeI_SinglePanel_codebookSubsetRestriction_i2Present bool
	if TypeI_SinglePanel_codebookSubsetRestriction_i2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.N1_n2.Decode(r); err != nil {
		return utils.WrapError("Decode N1_n2", err)
	}
	if TypeI_SinglePanel_codebookSubsetRestriction_i2Present {
		var tmp_bs_TypeI_SinglePanel_codebookSubsetRestriction_i2 []byte
		var n_TypeI_SinglePanel_codebookSubsetRestriction_i2 uint
		if tmp_bs_TypeI_SinglePanel_codebookSubsetRestriction_i2, n_TypeI_SinglePanel_codebookSubsetRestriction_i2, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode TypeI_SinglePanel_codebookSubsetRestriction_i2", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_TypeI_SinglePanel_codebookSubsetRestriction_i2,
			NumBits: uint64(n_TypeI_SinglePanel_codebookSubsetRestriction_i2),
		}
		ie.TypeI_SinglePanel_codebookSubsetRestriction_i2 = &tmp_bitstring
	}
	return nil
}
