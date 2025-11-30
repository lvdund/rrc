package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_r16_codebookType_type2_subType_typeII_PortSelection_r16 struct {
	PortSelectionSamplingSize_r16          CodebookConfig_r16_codebookType_type2_subType_typeII_PortSelection_r16_portSelectionSamplingSize_r16 `madatory`
	TypeII_PortSelectionRI_Restriction_r16 aper.BitString                                                                                       `lb:4,ub:4,madatory`
}

func (ie *CodebookConfig_r16_codebookType_type2_subType_typeII_PortSelection_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.PortSelectionSamplingSize_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PortSelectionSamplingSize_r16", err)
	}
	if err = w.WriteBitString(ie.TypeII_PortSelectionRI_Restriction_r16.Bytes, uint(ie.TypeII_PortSelectionRI_Restriction_r16.NumBits), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteBitString TypeII_PortSelectionRI_Restriction_r16", err)
	}
	return nil
}

func (ie *CodebookConfig_r16_codebookType_type2_subType_typeII_PortSelection_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.PortSelectionSamplingSize_r16.Decode(r); err != nil {
		return utils.WrapError("Decode PortSelectionSamplingSize_r16", err)
	}
	var tmp_bs_TypeII_PortSelectionRI_Restriction_r16 []byte
	var n_TypeII_PortSelectionRI_Restriction_r16 uint
	if tmp_bs_TypeII_PortSelectionRI_Restriction_r16, n_TypeII_PortSelectionRI_Restriction_r16, err = r.ReadBitString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadBitString TypeII_PortSelectionRI_Restriction_r16", err)
	}
	ie.TypeII_PortSelectionRI_Restriction_r16 = aper.BitString{
		Bytes:   tmp_bs_TypeII_PortSelectionRI_Restriction_r16,
		NumBits: uint64(n_TypeII_PortSelectionRI_Restriction_r16),
	}
	return nil
}
