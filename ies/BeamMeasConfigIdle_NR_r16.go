package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BeamMeasConfigIdle_NR_r16 struct {
	ReportQuantityRS_Indexes_r16  BeamMeasConfigIdle_NR_r16_reportQuantityRS_Indexes_r16 `madatory`
	MaxNrofRS_IndexesToReport_r16 int64                                                  `lb:1,ub:maxNrofIndexesToReport,madatory`
	IncludeBeamMeasurements_r16   bool                                                   `madatory`
}

func (ie *BeamMeasConfigIdle_NR_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.ReportQuantityRS_Indexes_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ReportQuantityRS_Indexes_r16", err)
	}
	if err = w.WriteInteger(ie.MaxNrofRS_IndexesToReport_r16, &aper.Constraint{Lb: 1, Ub: maxNrofIndexesToReport}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNrofRS_IndexesToReport_r16", err)
	}
	if err = w.WriteBoolean(ie.IncludeBeamMeasurements_r16); err != nil {
		return utils.WrapError("WriteBoolean IncludeBeamMeasurements_r16", err)
	}
	return nil
}

func (ie *BeamMeasConfigIdle_NR_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.ReportQuantityRS_Indexes_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ReportQuantityRS_Indexes_r16", err)
	}
	var tmp_int_MaxNrofRS_IndexesToReport_r16 int64
	if tmp_int_MaxNrofRS_IndexesToReport_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxNrofIndexesToReport}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNrofRS_IndexesToReport_r16", err)
	}
	ie.MaxNrofRS_IndexesToReport_r16 = tmp_int_MaxNrofRS_IndexesToReport_r16
	var tmp_bool_IncludeBeamMeasurements_r16 bool
	if tmp_bool_IncludeBeamMeasurements_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean IncludeBeamMeasurements_r16", err)
	}
	ie.IncludeBeamMeasurements_r16 = tmp_bool_IncludeBeamMeasurements_r16
	return nil
}
