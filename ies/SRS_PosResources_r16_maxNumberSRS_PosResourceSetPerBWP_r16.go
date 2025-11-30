package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16_Enum_n1  aper.Enumerated = 0
	SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16_Enum_n2  aper.Enumerated = 1
	SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16_Enum_n4  aper.Enumerated = 2
	SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16_Enum_n8  aper.Enumerated = 3
	SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16_Enum_n12 aper.Enumerated = 4
	SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16_Enum_n16 aper.Enumerated = 5
)

type SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16 struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16", err)
	}
	return nil
}

func (ie *SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
