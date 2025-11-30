package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type TDD_UL_DL_SlotConfig_IAB_MT_r16 struct {
	SlotIndex_r16      TDD_UL_DL_SlotIndex                                 `madatory`
	Symbols_IAB_MT_r16 *TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16 `lb:1,ub:maxNrofSymbols_1,optional`
}

func (ie *TDD_UL_DL_SlotConfig_IAB_MT_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Symbols_IAB_MT_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.SlotIndex_r16.Encode(w); err != nil {
		return utils.WrapError("Encode SlotIndex_r16", err)
	}
	if ie.Symbols_IAB_MT_r16 != nil {
		if err = ie.Symbols_IAB_MT_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Symbols_IAB_MT_r16", err)
		}
	}
	return nil
}

func (ie *TDD_UL_DL_SlotConfig_IAB_MT_r16) Decode(r *aper.AperReader) error {
	var err error
	var Symbols_IAB_MT_r16Present bool
	if Symbols_IAB_MT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.SlotIndex_r16.Decode(r); err != nil {
		return utils.WrapError("Decode SlotIndex_r16", err)
	}
	if Symbols_IAB_MT_r16Present {
		ie.Symbols_IAB_MT_r16 = new(TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16)
		if err = ie.Symbols_IAB_MT_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Symbols_IAB_MT_r16", err)
		}
	}
	return nil
}
