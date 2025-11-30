package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCResumeComplete_v1610_IEs_mobilityState_r16_Enum_normal aper.Enumerated = 0
	RRCResumeComplete_v1610_IEs_mobilityState_r16_Enum_medium aper.Enumerated = 1
	RRCResumeComplete_v1610_IEs_mobilityState_r16_Enum_high   aper.Enumerated = 2
	RRCResumeComplete_v1610_IEs_mobilityState_r16_Enum_spare  aper.Enumerated = 3
)

type RRCResumeComplete_v1610_IEs_mobilityState_r16 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *RRCResumeComplete_v1610_IEs_mobilityState_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode RRCResumeComplete_v1610_IEs_mobilityState_r16", err)
	}
	return nil
}

func (ie *RRCResumeComplete_v1610_IEs_mobilityState_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode RRCResumeComplete_v1610_IEs_mobilityState_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
