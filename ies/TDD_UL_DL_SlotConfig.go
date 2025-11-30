package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type TDD_UL_DL_SlotConfig struct {
	SlotIndex TDD_UL_DL_SlotIndex           `madatory`
	Symbols   *TDD_UL_DL_SlotConfig_symbols `lb:1,ub:maxNrofSymbols_1,optional`
}

func (ie *TDD_UL_DL_SlotConfig) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Symbols != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.SlotIndex.Encode(w); err != nil {
		return utils.WrapError("Encode SlotIndex", err)
	}
	if ie.Symbols != nil {
		if err = ie.Symbols.Encode(w); err != nil {
			return utils.WrapError("Encode Symbols", err)
		}
	}
	return nil
}

func (ie *TDD_UL_DL_SlotConfig) Decode(r *aper.AperReader) error {
	var err error
	var SymbolsPresent bool
	if SymbolsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.SlotIndex.Decode(r); err != nil {
		return utils.WrapError("Decode SlotIndex", err)
	}
	if SymbolsPresent {
		ie.Symbols = new(TDD_UL_DL_SlotConfig_symbols)
		if err = ie.Symbols.Decode(r); err != nil {
			return utils.WrapError("Decode Symbols", err)
		}
	}
	return nil
}
