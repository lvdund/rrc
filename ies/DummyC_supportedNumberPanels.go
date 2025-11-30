package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyC_supportedNumberPanels_Enum_n2 aper.Enumerated = 0
	DummyC_supportedNumberPanels_Enum_n4 aper.Enumerated = 1
)

type DummyC_supportedNumberPanels struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *DummyC_supportedNumberPanels) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode DummyC_supportedNumberPanels", err)
	}
	return nil
}

func (ie *DummyC_supportedNumberPanels) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode DummyC_supportedNumberPanels", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
