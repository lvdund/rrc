package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLC_Config_Choice_nothing uint64 = iota
	RLC_Config_Choice_Am
	RLC_Config_Choice_Um_Bi_Directional
	RLC_Config_Choice_Um_Uni_Directional_UL
	RLC_Config_Choice_Um_Uni_Directional_DL
)

type RLC_Config struct {
	Choice                uint64
	Am                    *RLC_Config_am
	Um_Bi_Directional     *RLC_Config_um_Bi_Directional
	Um_Uni_Directional_UL *RLC_Config_um_Uni_Directional_UL
	Um_Uni_Directional_DL *RLC_Config_um_Uni_Directional_DL
}

func (ie *RLC_Config) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RLC_Config_Choice_Am:
		if err = ie.Am.Encode(w); err != nil {
			err = utils.WrapError("Encode Am", err)
		}
	case RLC_Config_Choice_Um_Bi_Directional:
		if err = ie.Um_Bi_Directional.Encode(w); err != nil {
			err = utils.WrapError("Encode Um_Bi_Directional", err)
		}
	case RLC_Config_Choice_Um_Uni_Directional_UL:
		if err = ie.Um_Uni_Directional_UL.Encode(w); err != nil {
			err = utils.WrapError("Encode Um_Uni_Directional_UL", err)
		}
	case RLC_Config_Choice_Um_Uni_Directional_DL:
		if err = ie.Um_Uni_Directional_DL.Encode(w); err != nil {
			err = utils.WrapError("Encode Um_Uni_Directional_DL", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RLC_Config) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RLC_Config_Choice_Am:
		ie.Am = new(RLC_Config_am)
		if err = ie.Am.Decode(r); err != nil {
			return utils.WrapError("Decode Am", err)
		}
	case RLC_Config_Choice_Um_Bi_Directional:
		ie.Um_Bi_Directional = new(RLC_Config_um_Bi_Directional)
		if err = ie.Um_Bi_Directional.Decode(r); err != nil {
			return utils.WrapError("Decode Um_Bi_Directional", err)
		}
	case RLC_Config_Choice_Um_Uni_Directional_UL:
		ie.Um_Uni_Directional_UL = new(RLC_Config_um_Uni_Directional_UL)
		if err = ie.Um_Uni_Directional_UL.Decode(r); err != nil {
			return utils.WrapError("Decode Um_Uni_Directional_UL", err)
		}
	case RLC_Config_Choice_Um_Uni_Directional_DL:
		ie.Um_Uni_Directional_DL = new(RLC_Config_um_Uni_Directional_DL)
		if err = ie.Um_Uni_Directional_DL.Decode(r); err != nil {
			return utils.WrapError("Decode Um_Uni_Directional_DL", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
