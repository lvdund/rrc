package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_two struct {
	TwoTX_CodebookSubsetRestriction uper.BitString `lb:6,ub:6,madatory`
}

func (ie *CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_two) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.TwoTX_CodebookSubsetRestriction.Bytes, uint(ie.TwoTX_CodebookSubsetRestriction.NumBits), &uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
		return utils.WrapError("WriteBitString TwoTX_CodebookSubsetRestriction", err)
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_two) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bs_TwoTX_CodebookSubsetRestriction []byte
	var n_TwoTX_CodebookSubsetRestriction uint
	if tmp_bs_TwoTX_CodebookSubsetRestriction, n_TwoTX_CodebookSubsetRestriction, err = r.ReadBitString(&uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
		return utils.WrapError("ReadBitString TwoTX_CodebookSubsetRestriction", err)
	}
	ie.TwoTX_CodebookSubsetRestriction = uper.BitString{
		Bytes:   tmp_bs_TwoTX_CodebookSubsetRestriction,
		NumBits: uint64(n_TwoTX_CodebookSubsetRestriction),
	}
	return nil
}
