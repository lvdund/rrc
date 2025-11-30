package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyI_supportedSRS_TxPortSwitch_Enum_t1r2      aper.Enumerated = 0
	DummyI_supportedSRS_TxPortSwitch_Enum_t1r4      aper.Enumerated = 1
	DummyI_supportedSRS_TxPortSwitch_Enum_t2r4      aper.Enumerated = 2
	DummyI_supportedSRS_TxPortSwitch_Enum_t1r4_t2r4 aper.Enumerated = 3
	DummyI_supportedSRS_TxPortSwitch_Enum_tr_equal  aper.Enumerated = 4
)

type DummyI_supportedSRS_TxPortSwitch struct {
	Value aper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *DummyI_supportedSRS_TxPortSwitch) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode DummyI_supportedSRS_TxPortSwitch", err)
	}
	return nil
}

func (ie *DummyI_supportedSRS_TxPortSwitch) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode DummyI_supportedSRS_TxPortSwitch", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
