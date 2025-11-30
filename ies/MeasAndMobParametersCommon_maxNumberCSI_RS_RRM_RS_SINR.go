package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR_Enum_n4  aper.Enumerated = 0
	MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR_Enum_n8  aper.Enumerated = 1
	MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR_Enum_n16 aper.Enumerated = 2
	MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR_Enum_n32 aper.Enumerated = 3
	MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR_Enum_n64 aper.Enumerated = 4
	MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR_Enum_n96 aper.Enumerated = 5
)

type MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR", err)
	}
	return nil
}

func (ie *MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
