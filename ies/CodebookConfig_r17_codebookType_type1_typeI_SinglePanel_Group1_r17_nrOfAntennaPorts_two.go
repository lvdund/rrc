package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_two struct {
	TwoTX_CodebookSubsetRestriction1_r17 uper.BitString `lb:6,ub:6,madatory`
}

func (ie *CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_two) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.TwoTX_CodebookSubsetRestriction1_r17.Bytes, uint(ie.TwoTX_CodebookSubsetRestriction1_r17.NumBits), &uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
		return utils.WrapError("WriteBitString TwoTX_CodebookSubsetRestriction1_r17", err)
	}
	return nil
}

func (ie *CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_two) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bs_TwoTX_CodebookSubsetRestriction1_r17 []byte
	var n_TwoTX_CodebookSubsetRestriction1_r17 uint
	if tmp_bs_TwoTX_CodebookSubsetRestriction1_r17, n_TwoTX_CodebookSubsetRestriction1_r17, err = r.ReadBitString(&uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
		return utils.WrapError("ReadBitString TwoTX_CodebookSubsetRestriction1_r17", err)
	}
	ie.TwoTX_CodebookSubsetRestriction1_r17 = uper.BitString{
		Bytes:   tmp_bs_TwoTX_CodebookSubsetRestriction1_r17,
		NumBits: uint64(n_TwoTX_CodebookSubsetRestriction1_r17),
	}
	return nil
}
