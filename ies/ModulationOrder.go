package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ModulationOrder_Enum_bpsk_halfpi aper.Enumerated = 0
	ModulationOrder_Enum_bpsk        aper.Enumerated = 1
	ModulationOrder_Enum_qpsk        aper.Enumerated = 2
	ModulationOrder_Enum_qam16       aper.Enumerated = 3
	ModulationOrder_Enum_qam64       aper.Enumerated = 4
	ModulationOrder_Enum_qam256      aper.Enumerated = 5
)

type ModulationOrder struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *ModulationOrder) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode ModulationOrder", err)
	}
	return nil
}

func (ie *ModulationOrder) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode ModulationOrder", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
