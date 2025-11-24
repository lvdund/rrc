package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_codebookType_type1_subType_typeI_SinglePanel struct {
	NrOfAntennaPorts                 *CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts `lb:6,ub:6,optional`
	TypeI_SinglePanel_ri_Restriction uper.BitString                                                                `lb:8,ub:8,madatory`
}

func (ie *CodebookConfig_codebookType_type1_subType_typeI_SinglePanel) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.NrOfAntennaPorts != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.NrOfAntennaPorts != nil {
		if err = ie.NrOfAntennaPorts.Encode(w); err != nil {
			return utils.WrapError("Encode NrOfAntennaPorts", err)
		}
	}
	if err = w.WriteBitString(ie.TypeI_SinglePanel_ri_Restriction.Bytes, uint(ie.TypeI_SinglePanel_ri_Restriction.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteBitString TypeI_SinglePanel_ri_Restriction", err)
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type1_subType_typeI_SinglePanel) Decode(r *uper.UperReader) error {
	var err error
	var NrOfAntennaPortsPresent bool
	if NrOfAntennaPortsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if NrOfAntennaPortsPresent {
		ie.NrOfAntennaPorts = new(CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts)
		if err = ie.NrOfAntennaPorts.Decode(r); err != nil {
			return utils.WrapError("Decode NrOfAntennaPorts", err)
		}
	}
	var tmp_bs_TypeI_SinglePanel_ri_Restriction []byte
	var n_TypeI_SinglePanel_ri_Restriction uint
	if tmp_bs_TypeI_SinglePanel_ri_Restriction, n_TypeI_SinglePanel_ri_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadBitString TypeI_SinglePanel_ri_Restriction", err)
	}
	ie.TypeI_SinglePanel_ri_Restriction = uper.BitString{
		Bytes:   tmp_bs_TypeI_SinglePanel_ri_Restriction,
		NumBits: uint64(n_TypeI_SinglePanel_ri_Restriction),
	}
	return nil
}
