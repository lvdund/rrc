package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_Config_subslotLengthForPUCCH_r16_Choice_nothing uint64 = iota
	PUCCH_Config_subslotLengthForPUCCH_r16_Choice_NormalCP_r16
	PUCCH_Config_subslotLengthForPUCCH_r16_Choice_ExtendedCP_r16
)

type PUCCH_Config_subslotLengthForPUCCH_r16 struct {
	Choice         uint64
	NormalCP_r16   *PUCCH_Config_subslotLengthForPUCCH_r16_normalCP_r16
	ExtendedCP_r16 *PUCCH_Config_subslotLengthForPUCCH_r16_extendedCP_r16
}

func (ie *PUCCH_Config_subslotLengthForPUCCH_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_Config_subslotLengthForPUCCH_r16_Choice_NormalCP_r16:
		if err = ie.NormalCP_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode NormalCP_r16", err)
		}
	case PUCCH_Config_subslotLengthForPUCCH_r16_Choice_ExtendedCP_r16:
		if err = ie.ExtendedCP_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode ExtendedCP_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PUCCH_Config_subslotLengthForPUCCH_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_Config_subslotLengthForPUCCH_r16_Choice_NormalCP_r16:
		ie.NormalCP_r16 = new(PUCCH_Config_subslotLengthForPUCCH_r16_normalCP_r16)
		if err = ie.NormalCP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode NormalCP_r16", err)
		}
	case PUCCH_Config_subslotLengthForPUCCH_r16_Choice_ExtendedCP_r16:
		ie.ExtendedCP_r16 = new(PUCCH_Config_subslotLengthForPUCCH_r16_extendedCP_r16)
		if err = ie.ExtendedCP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ExtendedCP_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
