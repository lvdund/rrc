package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MBS_RNTI_SpecificConfig_r17_groupCommon_RNTI_r17_Choice_nothing uint64 = iota
	MBS_RNTI_SpecificConfig_r17_groupCommon_RNTI_r17_Choice_G_RNTI
	MBS_RNTI_SpecificConfig_r17_groupCommon_RNTI_r17_Choice_G_CS_RNTI
)

type MBS_RNTI_SpecificConfig_r17_groupCommon_RNTI_r17 struct {
	Choice    uint64
	G_RNTI    *RNTI_Value
	G_CS_RNTI *RNTI_Value
}

func (ie *MBS_RNTI_SpecificConfig_r17_groupCommon_RNTI_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MBS_RNTI_SpecificConfig_r17_groupCommon_RNTI_r17_Choice_G_RNTI:
		if err = ie.G_RNTI.Encode(w); err != nil {
			err = utils.WrapError("Encode G_RNTI", err)
		}
	case MBS_RNTI_SpecificConfig_r17_groupCommon_RNTI_r17_Choice_G_CS_RNTI:
		if err = ie.G_CS_RNTI.Encode(w); err != nil {
			err = utils.WrapError("Encode G_CS_RNTI", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MBS_RNTI_SpecificConfig_r17_groupCommon_RNTI_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MBS_RNTI_SpecificConfig_r17_groupCommon_RNTI_r17_Choice_G_RNTI:
		ie.G_RNTI = new(RNTI_Value)
		if err = ie.G_RNTI.Decode(r); err != nil {
			return utils.WrapError("Decode G_RNTI", err)
		}
	case MBS_RNTI_SpecificConfig_r17_groupCommon_RNTI_r17_Choice_G_CS_RNTI:
		ie.G_CS_RNTI = new(RNTI_Value)
		if err = ie.G_CS_RNTI.Decode(r); err != nil {
			return utils.WrapError("Decode G_CS_RNTI", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
