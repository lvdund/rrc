package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParameters_v1540_srs_TxSwitch struct {
	SupportedSRS_TxPortSwitch BandParameters_v1540_srs_TxSwitch_supportedSRS_TxPortSwitch `madatory`
	TxSwitchImpactToRx        *int64                                                      `lb:1,ub:32,optional`
	TxSwitchWithAnotherBand   *int64                                                      `lb:1,ub:32,optional`
}

func (ie *BandParameters_v1540_srs_TxSwitch) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.TxSwitchImpactToRx != nil, ie.TxSwitchWithAnotherBand != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.SupportedSRS_TxPortSwitch.Encode(w); err != nil {
		return utils.WrapError("Encode SupportedSRS_TxPortSwitch", err)
	}
	if ie.TxSwitchImpactToRx != nil {
		if err = w.WriteInteger(*ie.TxSwitchImpactToRx, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode TxSwitchImpactToRx", err)
		}
	}
	if ie.TxSwitchWithAnotherBand != nil {
		if err = w.WriteInteger(*ie.TxSwitchWithAnotherBand, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode TxSwitchWithAnotherBand", err)
		}
	}
	return nil
}

func (ie *BandParameters_v1540_srs_TxSwitch) Decode(r *uper.UperReader) error {
	var err error
	var TxSwitchImpactToRxPresent bool
	if TxSwitchImpactToRxPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TxSwitchWithAnotherBandPresent bool
	if TxSwitchWithAnotherBandPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.SupportedSRS_TxPortSwitch.Decode(r); err != nil {
		return utils.WrapError("Decode SupportedSRS_TxPortSwitch", err)
	}
	if TxSwitchImpactToRxPresent {
		var tmp_int_TxSwitchImpactToRx int64
		if tmp_int_TxSwitchImpactToRx, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode TxSwitchImpactToRx", err)
		}
		ie.TxSwitchImpactToRx = &tmp_int_TxSwitchImpactToRx
	}
	if TxSwitchWithAnotherBandPresent {
		var tmp_int_TxSwitchWithAnotherBand int64
		if tmp_int_TxSwitchWithAnotherBand, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode TxSwitchWithAnotherBand", err)
		}
		ie.TxSwitchWithAnotherBand = &tmp_int_TxSwitchWithAnotherBand
	}
	return nil
}
