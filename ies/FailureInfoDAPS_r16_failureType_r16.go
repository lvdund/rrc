package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FailureInfoDAPS_r16_failureType_r16_Enum_daps_failure aper.Enumerated = 0
	FailureInfoDAPS_r16_failureType_r16_Enum_spare3       aper.Enumerated = 1
	FailureInfoDAPS_r16_failureType_r16_Enum_spare2       aper.Enumerated = 2
	FailureInfoDAPS_r16_failureType_r16_Enum_spare1       aper.Enumerated = 3
)

type FailureInfoDAPS_r16_failureType_r16 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *FailureInfoDAPS_r16_failureType_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode FailureInfoDAPS_r16_failureType_r16", err)
	}
	return nil
}

func (ie *FailureInfoDAPS_r16_failureType_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode FailureInfoDAPS_r16_failureType_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
