package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_codebookType_type1_subType_typeI_MultiPanel struct {
	Ng_n1_n2       CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2 `lb:8,ub:8,madatory`
	Ri_Restriction uper.BitString                                                      `lb:4,ub:4,madatory`
}

func (ie *CodebookConfig_codebookType_type1_subType_typeI_MultiPanel) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Ng_n1_n2.Encode(w); err != nil {
		return utils.WrapError("Encode Ng_n1_n2", err)
	}
	if err = w.WriteBitString(ie.Ri_Restriction.Bytes, uint(ie.Ri_Restriction.NumBits), &uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteBitString Ri_Restriction", err)
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type1_subType_typeI_MultiPanel) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Ng_n1_n2.Decode(r); err != nil {
		return utils.WrapError("Decode Ng_n1_n2", err)
	}
	var tmp_bs_Ri_Restriction []byte
	var n_Ri_Restriction uint
	if tmp_bs_Ri_Restriction, n_Ri_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadBitString Ri_Restriction", err)
	}
	ie.Ri_Restriction = uper.BitString{
		Bytes:   tmp_bs_Ri_Restriction,
		NumBits: uint64(n_Ri_Restriction),
	}
	return nil
}
