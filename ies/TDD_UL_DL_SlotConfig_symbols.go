package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	TDD_UL_DL_SlotConfig_symbols_Choice_nothing uint64 = iota
	TDD_UL_DL_SlotConfig_symbols_Choice_AllDownlink
	TDD_UL_DL_SlotConfig_symbols_Choice_AllUplink
	TDD_UL_DL_SlotConfig_symbols_Choice_Explicit
)

type TDD_UL_DL_SlotConfig_symbols struct {
	Choice      uint64
	AllDownlink aper.NULL `madatory`
	AllUplink   aper.NULL `madatory`
	Explicit    *TDD_UL_DL_SlotConfig_symbols_explicit
}

func (ie *TDD_UL_DL_SlotConfig_symbols) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TDD_UL_DL_SlotConfig_symbols_Choice_AllDownlink:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode AllDownlink", err)
		}
	case TDD_UL_DL_SlotConfig_symbols_Choice_AllUplink:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode AllUplink", err)
		}
	case TDD_UL_DL_SlotConfig_symbols_Choice_Explicit:
		if err = ie.Explicit.Encode(w); err != nil {
			err = utils.WrapError("Encode Explicit", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *TDD_UL_DL_SlotConfig_symbols) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TDD_UL_DL_SlotConfig_symbols_Choice_AllDownlink:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode AllDownlink", err)
		}
	case TDD_UL_DL_SlotConfig_symbols_Choice_AllUplink:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode AllUplink", err)
		}
	case TDD_UL_DL_SlotConfig_symbols_Choice_Explicit:
		ie.Explicit = new(TDD_UL_DL_SlotConfig_symbols_explicit)
		if err = ie.Explicit.Decode(r); err != nil {
			return utils.WrapError("Decode Explicit", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
