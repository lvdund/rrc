package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_r16_codebookType_Choice_nothing uint64 = iota
	CodebookConfig_r16_codebookType_Choice_Type2
)

type CodebookConfig_r16_codebookType struct {
	Choice uint64
	Type2  *CodebookConfig_r16_codebookType_type2
}

func (ie *CodebookConfig_r16_codebookType) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_r16_codebookType_Choice_Type2:
		if err = ie.Type2.Encode(w); err != nil {
			err = utils.WrapError("Encode Type2", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CodebookConfig_r16_codebookType) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_r16_codebookType_Choice_Type2:
		ie.Type2 = new(CodebookConfig_r16_codebookType_type2)
		if err = ie.Type2.Decode(r); err != nil {
			return utils.WrapError("Decode Type2", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
