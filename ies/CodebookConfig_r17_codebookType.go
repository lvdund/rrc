package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_r17_codebookType_Choice_nothing uint64 = iota
	CodebookConfig_r17_codebookType_Choice_Type1
	CodebookConfig_r17_codebookType_Choice_Type2
)

type CodebookConfig_r17_codebookType struct {
	Choice uint64
	Type1  *CodebookConfig_r17_codebookType_type1
	Type2  *CodebookConfig_r17_codebookType_type2
}

func (ie *CodebookConfig_r17_codebookType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_r17_codebookType_Choice_Type1:
		if err = ie.Type1.Encode(w); err != nil {
			err = utils.WrapError("Encode Type1", err)
		}
	case CodebookConfig_r17_codebookType_Choice_Type2:
		if err = ie.Type2.Encode(w); err != nil {
			err = utils.WrapError("Encode Type2", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CodebookConfig_r17_codebookType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_r17_codebookType_Choice_Type1:
		ie.Type1 = new(CodebookConfig_r17_codebookType_type1)
		if err = ie.Type1.Decode(r); err != nil {
			return utils.WrapError("Decode Type1", err)
		}
	case CodebookConfig_r17_codebookType_Choice_Type2:
		ie.Type2 = new(CodebookConfig_r17_codebookType_type2)
		if err = ie.Type2.Decode(r); err != nil {
			return utils.WrapError("Decode Type2", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
