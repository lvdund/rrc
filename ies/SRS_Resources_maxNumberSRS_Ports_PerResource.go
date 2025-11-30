package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_Resources_maxNumberSRS_Ports_PerResource_Enum_n1 aper.Enumerated = 0
	SRS_Resources_maxNumberSRS_Ports_PerResource_Enum_n2 aper.Enumerated = 1
	SRS_Resources_maxNumberSRS_Ports_PerResource_Enum_n4 aper.Enumerated = 2
)

type SRS_Resources_maxNumberSRS_Ports_PerResource struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *SRS_Resources_maxNumberSRS_Ports_PerResource) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode SRS_Resources_maxNumberSRS_Ports_PerResource", err)
	}
	return nil
}

func (ie *SRS_Resources_maxNumberSRS_Ports_PerResource) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode SRS_Resources_maxNumberSRS_Ports_PerResource", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
