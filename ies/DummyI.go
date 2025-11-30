package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DummyI struct {
	SupportedSRS_TxPortSwitch DummyI_supportedSRS_TxPortSwitch `madatory`
	TxSwitchImpactToRx        *DummyI_txSwitchImpactToRx       `optional`
}

func (ie *DummyI) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.TxSwitchImpactToRx != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.SupportedSRS_TxPortSwitch.Encode(w); err != nil {
		return utils.WrapError("Encode SupportedSRS_TxPortSwitch", err)
	}
	if ie.TxSwitchImpactToRx != nil {
		if err = ie.TxSwitchImpactToRx.Encode(w); err != nil {
			return utils.WrapError("Encode TxSwitchImpactToRx", err)
		}
	}
	return nil
}

func (ie *DummyI) Decode(r *aper.AperReader) error {
	var err error
	var TxSwitchImpactToRxPresent bool
	if TxSwitchImpactToRxPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.SupportedSRS_TxPortSwitch.Decode(r); err != nil {
		return utils.WrapError("Decode SupportedSRS_TxPortSwitch", err)
	}
	if TxSwitchImpactToRxPresent {
		ie.TxSwitchImpactToRx = new(DummyI_txSwitchImpactToRx)
		if err = ie.TxSwitchImpactToRx.Decode(r); err != nil {
			return utils.WrapError("Decode TxSwitchImpactToRx", err)
		}
	}
	return nil
}
