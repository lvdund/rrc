package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_Resources_maxNumberAperiodicSRS_PerBWP_Enum_n1  aper.Enumerated = 0
	SRS_Resources_maxNumberAperiodicSRS_PerBWP_Enum_n2  aper.Enumerated = 1
	SRS_Resources_maxNumberAperiodicSRS_PerBWP_Enum_n4  aper.Enumerated = 2
	SRS_Resources_maxNumberAperiodicSRS_PerBWP_Enum_n8  aper.Enumerated = 3
	SRS_Resources_maxNumberAperiodicSRS_PerBWP_Enum_n16 aper.Enumerated = 4
)

type SRS_Resources_maxNumberAperiodicSRS_PerBWP struct {
	Value aper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *SRS_Resources_maxNumberAperiodicSRS_PerBWP) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode SRS_Resources_maxNumberAperiodicSRS_PerBWP", err)
	}
	return nil
}

func (ie *SRS_Resources_maxNumberAperiodicSRS_PerBWP) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode SRS_Resources_maxNumberAperiodicSRS_PerBWP", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
