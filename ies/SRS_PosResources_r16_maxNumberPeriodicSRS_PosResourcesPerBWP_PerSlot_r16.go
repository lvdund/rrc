package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_PosResources_r16_maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_Enum_n1  aper.Enumerated = 0
	SRS_PosResources_r16_maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_Enum_n2  aper.Enumerated = 1
	SRS_PosResources_r16_maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_Enum_n3  aper.Enumerated = 2
	SRS_PosResources_r16_maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_Enum_n4  aper.Enumerated = 3
	SRS_PosResources_r16_maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_Enum_n5  aper.Enumerated = 4
	SRS_PosResources_r16_maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_Enum_n6  aper.Enumerated = 5
	SRS_PosResources_r16_maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_Enum_n8  aper.Enumerated = 6
	SRS_PosResources_r16_maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_Enum_n10 aper.Enumerated = 7
	SRS_PosResources_r16_maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_Enum_n12 aper.Enumerated = 8
	SRS_PosResources_r16_maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16_Enum_n14 aper.Enumerated = 9
)

type SRS_PosResources_r16_maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16 struct {
	Value aper.Enumerated `lb:0,ub:9,madatory`
}

func (ie *SRS_PosResources_r16_maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("Encode SRS_PosResources_r16_maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16", err)
	}
	return nil
}

func (ie *SRS_PosResources_r16_maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("Decode SRS_PosResources_r16_maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
