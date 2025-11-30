package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_nothing uint64 = iota
	TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_AllDownlink_r16
	TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_AllUplink_r16
	TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_Explicit_r16
	TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_Explicit_IAB_MT_r16
)

type TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16 struct {
	Choice              uint64
	AllDownlink_r16     aper.NULL `madatory`
	AllUplink_r16       aper.NULL `madatory`
	Explicit_r16        *TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_explicit_r16
	Explicit_IAB_MT_r16 *TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_explicit_IAB_MT_r16
}

func (ie *TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_AllDownlink_r16:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode AllDownlink_r16", err)
		}
	case TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_AllUplink_r16:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode AllUplink_r16", err)
		}
	case TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_Explicit_r16:
		if err = ie.Explicit_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Explicit_r16", err)
		}
	case TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_Explicit_IAB_MT_r16:
		if err = ie.Explicit_IAB_MT_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Explicit_IAB_MT_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_AllDownlink_r16:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode AllDownlink_r16", err)
		}
	case TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_AllUplink_r16:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode AllUplink_r16", err)
		}
	case TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_Explicit_r16:
		ie.Explicit_r16 = new(TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_explicit_r16)
		if err = ie.Explicit_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Explicit_r16", err)
		}
	case TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_Explicit_IAB_MT_r16:
		ie.Explicit_IAB_MT_r16 = new(TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_explicit_IAB_MT_r16)
		if err = ie.Explicit_IAB_MT_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Explicit_IAB_MT_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
