package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_codebookType_type2_subType_typeII_PortSelection struct {
	PortSelectionSamplingSize          *CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize `optional`
	TypeII_PortSelectionRI_Restriction uper.BitString                                                                            `lb:2,ub:2,madatory`
}

func (ie *CodebookConfig_codebookType_type2_subType_typeII_PortSelection) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.PortSelectionSamplingSize != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.PortSelectionSamplingSize != nil {
		if err = ie.PortSelectionSamplingSize.Encode(w); err != nil {
			return utils.WrapError("Encode PortSelectionSamplingSize", err)
		}
	}
	if err = w.WriteBitString(ie.TypeII_PortSelectionRI_Restriction.Bytes, uint(ie.TypeII_PortSelectionRI_Restriction.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteBitString TypeII_PortSelectionRI_Restriction", err)
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type2_subType_typeII_PortSelection) Decode(r *uper.UperReader) error {
	var err error
	var PortSelectionSamplingSizePresent bool
	if PortSelectionSamplingSizePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if PortSelectionSamplingSizePresent {
		ie.PortSelectionSamplingSize = new(CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize)
		if err = ie.PortSelectionSamplingSize.Decode(r); err != nil {
			return utils.WrapError("Decode PortSelectionSamplingSize", err)
		}
	}
	var tmp_bs_TypeII_PortSelectionRI_Restriction []byte
	var n_TypeII_PortSelectionRI_Restriction uint
	if tmp_bs_TypeII_PortSelectionRI_Restriction, n_TypeII_PortSelectionRI_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadBitString TypeII_PortSelectionRI_Restriction", err)
	}
	ie.TypeII_PortSelectionRI_Restriction = uper.BitString{
		Bytes:   tmp_bs_TypeII_PortSelectionRI_Restriction,
		NumBits: uint64(n_TypeII_PortSelectionRI_Restriction),
	}
	return nil
}
