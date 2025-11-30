package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_dummy1_Enum_n1  aper.Enumerated = 0
	SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_dummy1_Enum_n2  aper.Enumerated = 1
	SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_dummy1_Enum_n4  aper.Enumerated = 2
	SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_dummy1_Enum_n8  aper.Enumerated = 3
	SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_dummy1_Enum_n16 aper.Enumerated = 4
	SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_dummy1_Enum_n32 aper.Enumerated = 5
	SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_dummy1_Enum_n64 aper.Enumerated = 6
)

type SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_dummy1 struct {
	Value aper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_dummy1) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_dummy1", err)
	}
	return nil
}

func (ie *SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_dummy1) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_dummy1", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
