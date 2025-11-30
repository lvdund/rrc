package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_Choice_nothing uint64 = iota
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_Choice_Two
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_Choice_MoreThanTwo
)

type CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts struct {
	Choice      uint64
	Two         *CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_two
	MoreThanTwo *CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo
}

func (ie *CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_Choice_Two:
		if err = ie.Two.Encode(w); err != nil {
			err = utils.WrapError("Encode Two", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_Choice_MoreThanTwo:
		if err = ie.MoreThanTwo.Encode(w); err != nil {
			err = utils.WrapError("Encode MoreThanTwo", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_Choice_Two:
		ie.Two = new(CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_two)
		if err = ie.Two.Decode(r); err != nil {
			return utils.WrapError("Decode Two", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_Choice_MoreThanTwo:
		ie.MoreThanTwo = new(CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo)
		if err = ie.MoreThanTwo.Decode(r); err != nil {
			return utils.WrapError("Decode MoreThanTwo", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
