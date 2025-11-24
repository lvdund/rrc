package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PTRS_DensityRecommendationUL struct {
	FrequencyDensity1 int64 `lb:1,ub:276,madatory`
	FrequencyDensity2 int64 `lb:1,ub:276,madatory`
	TimeDensity1      int64 `lb:0,ub:29,madatory`
	TimeDensity2      int64 `lb:0,ub:29,madatory`
	TimeDensity3      int64 `lb:0,ub:29,madatory`
	SampleDensity1    int64 `lb:1,ub:276,madatory`
	SampleDensity2    int64 `lb:1,ub:276,madatory`
	SampleDensity3    int64 `lb:1,ub:276,madatory`
	SampleDensity4    int64 `lb:1,ub:276,madatory`
	SampleDensity5    int64 `lb:1,ub:276,madatory`
}

func (ie *PTRS_DensityRecommendationUL) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.FrequencyDensity1, &uper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("WriteInteger FrequencyDensity1", err)
	}
	if err = w.WriteInteger(ie.FrequencyDensity2, &uper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("WriteInteger FrequencyDensity2", err)
	}
	if err = w.WriteInteger(ie.TimeDensity1, &uper.Constraint{Lb: 0, Ub: 29}, false); err != nil {
		return utils.WrapError("WriteInteger TimeDensity1", err)
	}
	if err = w.WriteInteger(ie.TimeDensity2, &uper.Constraint{Lb: 0, Ub: 29}, false); err != nil {
		return utils.WrapError("WriteInteger TimeDensity2", err)
	}
	if err = w.WriteInteger(ie.TimeDensity3, &uper.Constraint{Lb: 0, Ub: 29}, false); err != nil {
		return utils.WrapError("WriteInteger TimeDensity3", err)
	}
	if err = w.WriteInteger(ie.SampleDensity1, &uper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("WriteInteger SampleDensity1", err)
	}
	if err = w.WriteInteger(ie.SampleDensity2, &uper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("WriteInteger SampleDensity2", err)
	}
	if err = w.WriteInteger(ie.SampleDensity3, &uper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("WriteInteger SampleDensity3", err)
	}
	if err = w.WriteInteger(ie.SampleDensity4, &uper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("WriteInteger SampleDensity4", err)
	}
	if err = w.WriteInteger(ie.SampleDensity5, &uper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("WriteInteger SampleDensity5", err)
	}
	return nil
}

func (ie *PTRS_DensityRecommendationUL) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_FrequencyDensity1 int64
	if tmp_int_FrequencyDensity1, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("ReadInteger FrequencyDensity1", err)
	}
	ie.FrequencyDensity1 = tmp_int_FrequencyDensity1
	var tmp_int_FrequencyDensity2 int64
	if tmp_int_FrequencyDensity2, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("ReadInteger FrequencyDensity2", err)
	}
	ie.FrequencyDensity2 = tmp_int_FrequencyDensity2
	var tmp_int_TimeDensity1 int64
	if tmp_int_TimeDensity1, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 29}, false); err != nil {
		return utils.WrapError("ReadInteger TimeDensity1", err)
	}
	ie.TimeDensity1 = tmp_int_TimeDensity1
	var tmp_int_TimeDensity2 int64
	if tmp_int_TimeDensity2, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 29}, false); err != nil {
		return utils.WrapError("ReadInteger TimeDensity2", err)
	}
	ie.TimeDensity2 = tmp_int_TimeDensity2
	var tmp_int_TimeDensity3 int64
	if tmp_int_TimeDensity3, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 29}, false); err != nil {
		return utils.WrapError("ReadInteger TimeDensity3", err)
	}
	ie.TimeDensity3 = tmp_int_TimeDensity3
	var tmp_int_SampleDensity1 int64
	if tmp_int_SampleDensity1, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("ReadInteger SampleDensity1", err)
	}
	ie.SampleDensity1 = tmp_int_SampleDensity1
	var tmp_int_SampleDensity2 int64
	if tmp_int_SampleDensity2, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("ReadInteger SampleDensity2", err)
	}
	ie.SampleDensity2 = tmp_int_SampleDensity2
	var tmp_int_SampleDensity3 int64
	if tmp_int_SampleDensity3, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("ReadInteger SampleDensity3", err)
	}
	ie.SampleDensity3 = tmp_int_SampleDensity3
	var tmp_int_SampleDensity4 int64
	if tmp_int_SampleDensity4, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("ReadInteger SampleDensity4", err)
	}
	ie.SampleDensity4 = tmp_int_SampleDensity4
	var tmp_int_SampleDensity5 int64
	if tmp_int_SampleDensity5, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("ReadInteger SampleDensity5", err)
	}
	ie.SampleDensity5 = tmp_int_SampleDensity5
	return nil
}
