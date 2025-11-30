package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_explicit_r16 struct {
	NrofDownlinkSymbols_r16 *int64 `lb:1,ub:maxNrofSymbols_1,optional`
	NrofUplinkSymbols_r16   *int64 `lb:1,ub:maxNrofSymbols_1,optional`
}

func (ie *TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_explicit_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.NrofDownlinkSymbols_r16 != nil, ie.NrofUplinkSymbols_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.NrofDownlinkSymbols_r16 != nil {
		if err = w.WriteInteger(*ie.NrofDownlinkSymbols_r16, &aper.Constraint{Lb: 1, Ub: maxNrofSymbols_1}, false); err != nil {
			return utils.WrapError("Encode NrofDownlinkSymbols_r16", err)
		}
	}
	if ie.NrofUplinkSymbols_r16 != nil {
		if err = w.WriteInteger(*ie.NrofUplinkSymbols_r16, &aper.Constraint{Lb: 1, Ub: maxNrofSymbols_1}, false); err != nil {
			return utils.WrapError("Encode NrofUplinkSymbols_r16", err)
		}
	}
	return nil
}

func (ie *TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_explicit_r16) Decode(r *aper.AperReader) error {
	var err error
	var NrofDownlinkSymbols_r16Present bool
	if NrofDownlinkSymbols_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NrofUplinkSymbols_r16Present bool
	if NrofUplinkSymbols_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if NrofDownlinkSymbols_r16Present {
		var tmp_int_NrofDownlinkSymbols_r16 int64
		if tmp_int_NrofDownlinkSymbols_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxNrofSymbols_1}, false); err != nil {
			return utils.WrapError("Decode NrofDownlinkSymbols_r16", err)
		}
		ie.NrofDownlinkSymbols_r16 = &tmp_int_NrofDownlinkSymbols_r16
	}
	if NrofUplinkSymbols_r16Present {
		var tmp_int_NrofUplinkSymbols_r16 int64
		if tmp_int_NrofUplinkSymbols_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxNrofSymbols_1}, false); err != nil {
			return utils.WrapError("Decode NrofUplinkSymbols_r16", err)
		}
		ie.NrofUplinkSymbols_r16 = &tmp_int_NrofUplinkSymbols_r16
	}
	return nil
}
