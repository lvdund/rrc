package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasResultNR_triggeredEvent_r17_firstTriggeredEvent_Enum_condFirstEvent  aper.Enumerated = 0
	MeasResultNR_triggeredEvent_r17_firstTriggeredEvent_Enum_condSecondEvent aper.Enumerated = 1
)

type MeasResultNR_triggeredEvent_r17_firstTriggeredEvent struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *MeasResultNR_triggeredEvent_r17_firstTriggeredEvent) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode MeasResultNR_triggeredEvent_r17_firstTriggeredEvent", err)
	}
	return nil
}

func (ie *MeasResultNR_triggeredEvent_r17_firstTriggeredEvent) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode MeasResultNR_triggeredEvent_r17_firstTriggeredEvent", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
