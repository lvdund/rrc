package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_codebookType_type2_subType_typeII struct {
	N1_n2_codebookSubsetRestriction CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction `lb:16,ub:16,madatory`
	TypeII_RI_Restriction           aper.BitString                                                                   `lb:2,ub:2,madatory`
}

func (ie *CodebookConfig_codebookType_type2_subType_typeII) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.N1_n2_codebookSubsetRestriction.Encode(w); err != nil {
		return utils.WrapError("Encode N1_n2_codebookSubsetRestriction", err)
	}
	if err = w.WriteBitString(ie.TypeII_RI_Restriction.Bytes, uint(ie.TypeII_RI_Restriction.NumBits), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteBitString TypeII_RI_Restriction", err)
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type2_subType_typeII) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.N1_n2_codebookSubsetRestriction.Decode(r); err != nil {
		return utils.WrapError("Decode N1_n2_codebookSubsetRestriction", err)
	}
	var tmp_bs_TypeII_RI_Restriction []byte
	var n_TypeII_RI_Restriction uint
	if tmp_bs_TypeII_RI_Restriction, n_TypeII_RI_Restriction, err = r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadBitString TypeII_RI_Restriction", err)
	}
	ie.TypeII_RI_Restriction = aper.BitString{
		Bytes:   tmp_bs_TypeII_RI_Restriction,
		NumBits: uint64(n_TypeII_RI_Restriction),
	}
	return nil
}
