package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyE_maxNumberTxPortsPerResource_Enum_p4  aper.Enumerated = 0
	DummyE_maxNumberTxPortsPerResource_Enum_p8  aper.Enumerated = 1
	DummyE_maxNumberTxPortsPerResource_Enum_p12 aper.Enumerated = 2
	DummyE_maxNumberTxPortsPerResource_Enum_p16 aper.Enumerated = 3
	DummyE_maxNumberTxPortsPerResource_Enum_p24 aper.Enumerated = 4
	DummyE_maxNumberTxPortsPerResource_Enum_p32 aper.Enumerated = 5
)

type DummyE_maxNumberTxPortsPerResource struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *DummyE_maxNumberTxPortsPerResource) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode DummyE_maxNumberTxPortsPerResource", err)
	}
	return nil
}

func (ie *DummyE_maxNumberTxPortsPerResource) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode DummyE_maxNumberTxPortsPerResource", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
