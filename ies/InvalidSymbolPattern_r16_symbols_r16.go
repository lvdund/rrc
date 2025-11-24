package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	InvalidSymbolPattern_r16_symbols_r16_Choice_nothing uint64 = iota
	InvalidSymbolPattern_r16_symbols_r16_Choice_OneSlot
	InvalidSymbolPattern_r16_symbols_r16_Choice_TwoSlots
)

type InvalidSymbolPattern_r16_symbols_r16 struct {
	Choice   uint64
	OneSlot  uper.BitString `lb:14,ub:14,madatory`
	TwoSlots uper.BitString `lb:28,ub:28,madatory`
}

func (ie *InvalidSymbolPattern_r16_symbols_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case InvalidSymbolPattern_r16_symbols_r16_Choice_OneSlot:
		if err = w.WriteBitString(ie.OneSlot.Bytes, uint(ie.OneSlot.NumBits), &uper.Constraint{Lb: 14, Ub: 14}, false); err != nil {
			err = utils.WrapError("Encode OneSlot", err)
		}
	case InvalidSymbolPattern_r16_symbols_r16_Choice_TwoSlots:
		if err = w.WriteBitString(ie.TwoSlots.Bytes, uint(ie.TwoSlots.NumBits), &uper.Constraint{Lb: 28, Ub: 28}, false); err != nil {
			err = utils.WrapError("Encode TwoSlots", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *InvalidSymbolPattern_r16_symbols_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case InvalidSymbolPattern_r16_symbols_r16_Choice_OneSlot:
		var tmp_bs_OneSlot []byte
		var n_OneSlot uint
		if tmp_bs_OneSlot, n_OneSlot, err = r.ReadBitString(&uper.Constraint{Lb: 14, Ub: 14}, false); err != nil {
			return utils.WrapError("Decode OneSlot", err)
		}
		ie.OneSlot = uper.BitString{
			Bytes:   tmp_bs_OneSlot,
			NumBits: uint64(n_OneSlot),
		}
	case InvalidSymbolPattern_r16_symbols_r16_Choice_TwoSlots:
		var tmp_bs_TwoSlots []byte
		var n_TwoSlots uint
		if tmp_bs_TwoSlots, n_TwoSlots, err = r.ReadBitString(&uper.Constraint{Lb: 28, Ub: 28}, false); err != nil {
			return utils.WrapError("Decode TwoSlots", err)
		}
		ie.TwoSlots = uper.BitString{
			Bytes:   tmp_bs_TwoSlots,
			NumBits: uint64(n_TwoSlots),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
