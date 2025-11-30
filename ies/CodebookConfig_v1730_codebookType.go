package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_v1730_codebookType_Choice_nothing uint64 = iota
	CodebookConfig_v1730_codebookType_Choice_Type1
)

type CodebookConfig_v1730_codebookType struct {
	Choice uint64
	Type1  *CodebookConfig_v1730_codebookType_type1
}

func (ie *CodebookConfig_v1730_codebookType) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_v1730_codebookType_Choice_Type1:
		if err = ie.Type1.Encode(w); err != nil {
			err = utils.WrapError("Encode Type1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CodebookConfig_v1730_codebookType) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_v1730_codebookType_Choice_Type1:
		ie.Type1 = new(CodebookConfig_v1730_codebookType_type1)
		if err = ie.Type1.Decode(r); err != nil {
			return utils.WrapError("Decode Type1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
