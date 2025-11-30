package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyI_txSwitchImpactToRx_Enum_true aper.Enumerated = 0
)

type DummyI_txSwitchImpactToRx struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *DummyI_txSwitchImpactToRx) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode DummyI_txSwitchImpactToRx", err)
	}
	return nil
}

func (ie *DummyI_txSwitchImpactToRx) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode DummyI_txSwitchImpactToRx", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
