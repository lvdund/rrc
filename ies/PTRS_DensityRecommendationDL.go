package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PTRS_DensityRecommendationDL struct {
	FrequencyDensity1 int64 `lb:1,ub:276,madatory`
	FrequencyDensity2 int64 `lb:1,ub:276,madatory`
	TimeDensity1      int64 `lb:0,ub:29,madatory`
	TimeDensity2      int64 `lb:0,ub:29,madatory`
	TimeDensity3      int64 `lb:0,ub:29,madatory`
}

func (ie *PTRS_DensityRecommendationDL) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.FrequencyDensity1, &aper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("WriteInteger FrequencyDensity1", err)
	}
	if err = w.WriteInteger(ie.FrequencyDensity2, &aper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("WriteInteger FrequencyDensity2", err)
	}
	if err = w.WriteInteger(ie.TimeDensity1, &aper.Constraint{Lb: 0, Ub: 29}, false); err != nil {
		return utils.WrapError("WriteInteger TimeDensity1", err)
	}
	if err = w.WriteInteger(ie.TimeDensity2, &aper.Constraint{Lb: 0, Ub: 29}, false); err != nil {
		return utils.WrapError("WriteInteger TimeDensity2", err)
	}
	if err = w.WriteInteger(ie.TimeDensity3, &aper.Constraint{Lb: 0, Ub: 29}, false); err != nil {
		return utils.WrapError("WriteInteger TimeDensity3", err)
	}
	return nil
}

func (ie *PTRS_DensityRecommendationDL) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_FrequencyDensity1 int64
	if tmp_int_FrequencyDensity1, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("ReadInteger FrequencyDensity1", err)
	}
	ie.FrequencyDensity1 = tmp_int_FrequencyDensity1
	var tmp_int_FrequencyDensity2 int64
	if tmp_int_FrequencyDensity2, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("ReadInteger FrequencyDensity2", err)
	}
	ie.FrequencyDensity2 = tmp_int_FrequencyDensity2
	var tmp_int_TimeDensity1 int64
	if tmp_int_TimeDensity1, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 29}, false); err != nil {
		return utils.WrapError("ReadInteger TimeDensity1", err)
	}
	ie.TimeDensity1 = tmp_int_TimeDensity1
	var tmp_int_TimeDensity2 int64
	if tmp_int_TimeDensity2, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 29}, false); err != nil {
		return utils.WrapError("ReadInteger TimeDensity2", err)
	}
	ie.TimeDensity2 = tmp_int_TimeDensity2
	var tmp_int_TimeDensity3 int64
	if tmp_int_TimeDensity3, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 29}, false); err != nil {
		return utils.WrapError("ReadInteger TimeDensity3", err)
	}
	ie.TimeDensity3 = tmp_int_TimeDensity3
	return nil
}
