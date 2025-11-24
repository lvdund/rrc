package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_codebookType_type2_subType_Choice_nothing uint64 = iota
	CodebookConfig_codebookType_type2_subType_Choice_TypeII
	CodebookConfig_codebookType_type2_subType_Choice_TypeII_PortSelection
)

type CodebookConfig_codebookType_type2_subType struct {
	Choice               uint64
	TypeII               *CodebookConfig_codebookType_type2_subType_typeII
	TypeII_PortSelection *CodebookConfig_codebookType_type2_subType_typeII_PortSelection
}

func (ie *CodebookConfig_codebookType_type2_subType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_codebookType_type2_subType_Choice_TypeII:
		if err = ie.TypeII.Encode(w); err != nil {
			err = utils.WrapError("Encode TypeII", err)
		}
	case CodebookConfig_codebookType_type2_subType_Choice_TypeII_PortSelection:
		if err = ie.TypeII_PortSelection.Encode(w); err != nil {
			err = utils.WrapError("Encode TypeII_PortSelection", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CodebookConfig_codebookType_type2_subType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_codebookType_type2_subType_Choice_TypeII:
		ie.TypeII = new(CodebookConfig_codebookType_type2_subType_typeII)
		if err = ie.TypeII.Decode(r); err != nil {
			return utils.WrapError("Decode TypeII", err)
		}
	case CodebookConfig_codebookType_type2_subType_Choice_TypeII_PortSelection:
		ie.TypeII_PortSelection = new(CodebookConfig_codebookType_type2_subType_typeII_PortSelection)
		if err = ie.TypeII_PortSelection.Decode(r); err != nil {
			return utils.WrapError("Decode TypeII_PortSelection", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
