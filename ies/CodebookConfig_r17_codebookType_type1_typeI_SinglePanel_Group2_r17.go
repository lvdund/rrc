package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17 struct {
	NrOfAntennaPorts CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts `lb:6,ub:6,madatory`
}

func (ie *CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.NrOfAntennaPorts.Encode(w); err != nil {
		return utils.WrapError("Encode NrOfAntennaPorts", err)
	}
	return nil
}

func (ie *CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.NrOfAntennaPorts.Decode(r); err != nil {
		return utils.WrapError("Decode NrOfAntennaPorts", err)
	}
	return nil
}
