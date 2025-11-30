package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PhysicalCellGroupConfig_ackNackFeedbackMode_r16_Enum_joint    aper.Enumerated = 0
	PhysicalCellGroupConfig_ackNackFeedbackMode_r16_Enum_separate aper.Enumerated = 1
)

type PhysicalCellGroupConfig_ackNackFeedbackMode_r16 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PhysicalCellGroupConfig_ackNackFeedbackMode_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PhysicalCellGroupConfig_ackNackFeedbackMode_r16", err)
	}
	return nil
}

func (ie *PhysicalCellGroupConfig_ackNackFeedbackMode_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PhysicalCellGroupConfig_ackNackFeedbackMode_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
