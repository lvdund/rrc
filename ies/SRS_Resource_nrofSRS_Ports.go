package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_Resource_nrofSRS_Ports_Enum_port1  aper.Enumerated = 0
	SRS_Resource_nrofSRS_Ports_Enum_ports2 aper.Enumerated = 1
	SRS_Resource_nrofSRS_Ports_Enum_ports4 aper.Enumerated = 2
)

type SRS_Resource_nrofSRS_Ports struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *SRS_Resource_nrofSRS_Ports) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode SRS_Resource_nrofSRS_Ports", err)
	}
	return nil
}

func (ie *SRS_Resource_nrofSRS_Ports) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode SRS_Resource_nrofSRS_Ports", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
