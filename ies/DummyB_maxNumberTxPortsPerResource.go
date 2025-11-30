package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyB_maxNumberTxPortsPerResource_Enum_p2  aper.Enumerated = 0
	DummyB_maxNumberTxPortsPerResource_Enum_p4  aper.Enumerated = 1
	DummyB_maxNumberTxPortsPerResource_Enum_p8  aper.Enumerated = 2
	DummyB_maxNumberTxPortsPerResource_Enum_p12 aper.Enumerated = 3
	DummyB_maxNumberTxPortsPerResource_Enum_p16 aper.Enumerated = 4
	DummyB_maxNumberTxPortsPerResource_Enum_p24 aper.Enumerated = 5
	DummyB_maxNumberTxPortsPerResource_Enum_p32 aper.Enumerated = 6
)

type DummyB_maxNumberTxPortsPerResource struct {
	Value aper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *DummyB_maxNumberTxPortsPerResource) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode DummyB_maxNumberTxPortsPerResource", err)
	}
	return nil
}

func (ie *DummyB_maxNumberTxPortsPerResource) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode DummyB_maxNumberTxPortsPerResource", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
