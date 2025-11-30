package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BeamMeasConfigIdle_NR_r16_reportQuantityRS_Indexes_r16_Enum_rsrp aper.Enumerated = 0
	BeamMeasConfigIdle_NR_r16_reportQuantityRS_Indexes_r16_Enum_rsrq aper.Enumerated = 1
	BeamMeasConfigIdle_NR_r16_reportQuantityRS_Indexes_r16_Enum_both aper.Enumerated = 2
)

type BeamMeasConfigIdle_NR_r16_reportQuantityRS_Indexes_r16 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *BeamMeasConfigIdle_NR_r16_reportQuantityRS_Indexes_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode BeamMeasConfigIdle_NR_r16_reportQuantityRS_Indexes_r16", err)
	}
	return nil
}

func (ie *BeamMeasConfigIdle_NR_r16_reportQuantityRS_Indexes_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode BeamMeasConfigIdle_NR_r16_reportQuantityRS_Indexes_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
