package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_codebookType_type1_subType_Choice_nothing uint64 = iota
	CodebookConfig_codebookType_type1_subType_Choice_TypeI_SinglePanel
	CodebookConfig_codebookType_type1_subType_Choice_TypeI_MultiPanel
)

type CodebookConfig_codebookType_type1_subType struct {
	Choice            uint64
	TypeI_SinglePanel *CodebookConfig_codebookType_type1_subType_typeI_SinglePanel
	TypeI_MultiPanel  *CodebookConfig_codebookType_type1_subType_typeI_MultiPanel
}

func (ie *CodebookConfig_codebookType_type1_subType) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_codebookType_type1_subType_Choice_TypeI_SinglePanel:
		if err = ie.TypeI_SinglePanel.Encode(w); err != nil {
			err = utils.WrapError("Encode TypeI_SinglePanel", err)
		}
	case CodebookConfig_codebookType_type1_subType_Choice_TypeI_MultiPanel:
		if err = ie.TypeI_MultiPanel.Encode(w); err != nil {
			err = utils.WrapError("Encode TypeI_MultiPanel", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CodebookConfig_codebookType_type1_subType) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_codebookType_type1_subType_Choice_TypeI_SinglePanel:
		ie.TypeI_SinglePanel = new(CodebookConfig_codebookType_type1_subType_typeI_SinglePanel)
		if err = ie.TypeI_SinglePanel.Decode(r); err != nil {
			return utils.WrapError("Decode TypeI_SinglePanel", err)
		}
	case CodebookConfig_codebookType_type1_subType_Choice_TypeI_MultiPanel:
		ie.TypeI_MultiPanel = new(CodebookConfig_codebookType_type1_subType_typeI_MultiPanel)
		if err = ie.TypeI_MultiPanel.Decode(r); err != nil {
			return utils.WrapError("Decode TypeI_MultiPanel", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
