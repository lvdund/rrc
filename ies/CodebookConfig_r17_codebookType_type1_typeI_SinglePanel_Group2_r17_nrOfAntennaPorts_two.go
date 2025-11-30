package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_two struct {
	TwoTX_CodebookSubsetRestriction2_r17 aper.BitString `lb:6,ub:6,madatory`
}

func (ie *CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_two) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteBitString(ie.TwoTX_CodebookSubsetRestriction2_r17.Bytes, uint(ie.TwoTX_CodebookSubsetRestriction2_r17.NumBits), &aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
		return utils.WrapError("WriteBitString TwoTX_CodebookSubsetRestriction2_r17", err)
	}
	return nil
}

func (ie *CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_two) Decode(r *aper.AperReader) error {
	var err error
	var tmp_bs_TwoTX_CodebookSubsetRestriction2_r17 []byte
	var n_TwoTX_CodebookSubsetRestriction2_r17 uint
	if tmp_bs_TwoTX_CodebookSubsetRestriction2_r17, n_TwoTX_CodebookSubsetRestriction2_r17, err = r.ReadBitString(&aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
		return utils.WrapError("ReadBitString TwoTX_CodebookSubsetRestriction2_r17", err)
	}
	ie.TwoTX_CodebookSubsetRestriction2_r17 = aper.BitString{
		Bytes:   tmp_bs_TwoTX_CodebookSubsetRestriction2_r17,
		NumBits: uint64(n_TwoTX_CodebookSubsetRestriction2_r17),
	}
	return nil
}
