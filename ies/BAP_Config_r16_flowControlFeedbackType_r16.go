package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BAP_Config_r16_flowControlFeedbackType_r16_Enum_perBH_RLC_Channel aper.Enumerated = 0
	BAP_Config_r16_flowControlFeedbackType_r16_Enum_perRoutingID      aper.Enumerated = 1
	BAP_Config_r16_flowControlFeedbackType_r16_Enum_both              aper.Enumerated = 2
)

type BAP_Config_r16_flowControlFeedbackType_r16 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *BAP_Config_r16_flowControlFeedbackType_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode BAP_Config_r16_flowControlFeedbackType_r16", err)
	}
	return nil
}

func (ie *BAP_Config_r16_flowControlFeedbackType_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode BAP_Config_r16_flowControlFeedbackType_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
