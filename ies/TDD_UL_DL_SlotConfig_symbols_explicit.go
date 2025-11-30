package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type TDD_UL_DL_SlotConfig_symbols_explicit struct {
	NrofDownlinkSymbols *int64 `lb:1,ub:maxNrofSymbols_1,optional`
	NrofUplinkSymbols   *int64 `lb:1,ub:maxNrofSymbols_1,optional`
}

func (ie *TDD_UL_DL_SlotConfig_symbols_explicit) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.NrofDownlinkSymbols != nil, ie.NrofUplinkSymbols != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.NrofDownlinkSymbols != nil {
		if err = w.WriteInteger(*ie.NrofDownlinkSymbols, &aper.Constraint{Lb: 1, Ub: maxNrofSymbols_1}, false); err != nil {
			return utils.WrapError("Encode NrofDownlinkSymbols", err)
		}
	}
	if ie.NrofUplinkSymbols != nil {
		if err = w.WriteInteger(*ie.NrofUplinkSymbols, &aper.Constraint{Lb: 1, Ub: maxNrofSymbols_1}, false); err != nil {
			return utils.WrapError("Encode NrofUplinkSymbols", err)
		}
	}
	return nil
}

func (ie *TDD_UL_DL_SlotConfig_symbols_explicit) Decode(r *aper.AperReader) error {
	var err error
	var NrofDownlinkSymbolsPresent bool
	if NrofDownlinkSymbolsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NrofUplinkSymbolsPresent bool
	if NrofUplinkSymbolsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if NrofDownlinkSymbolsPresent {
		var tmp_int_NrofDownlinkSymbols int64
		if tmp_int_NrofDownlinkSymbols, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxNrofSymbols_1}, false); err != nil {
			return utils.WrapError("Decode NrofDownlinkSymbols", err)
		}
		ie.NrofDownlinkSymbols = &tmp_int_NrofDownlinkSymbols
	}
	if NrofUplinkSymbolsPresent {
		var tmp_int_NrofUplinkSymbols int64
		if tmp_int_NrofUplinkSymbols, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxNrofSymbols_1}, false); err != nil {
			return utils.WrapError("Decode NrofUplinkSymbols", err)
		}
		ie.NrofUplinkSymbols = &tmp_int_NrofUplinkSymbols
	}
	return nil
}
