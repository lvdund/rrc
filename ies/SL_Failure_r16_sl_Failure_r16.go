package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_Failure_r16_sl_Failure_r16_Enum_rlf             aper.Enumerated = 0
	SL_Failure_r16_sl_Failure_r16_Enum_configFailure   aper.Enumerated = 1
	SL_Failure_r16_sl_Failure_r16_Enum_drxReject_v1710 aper.Enumerated = 2
	SL_Failure_r16_sl_Failure_r16_Enum_spare5          aper.Enumerated = 3
	SL_Failure_r16_sl_Failure_r16_Enum_spare4          aper.Enumerated = 4
	SL_Failure_r16_sl_Failure_r16_Enum_spare3          aper.Enumerated = 5
	SL_Failure_r16_sl_Failure_r16_Enum_spare2          aper.Enumerated = 6
	SL_Failure_r16_sl_Failure_r16_Enum_spare1          aper.Enumerated = 7
)

type SL_Failure_r16_sl_Failure_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SL_Failure_r16_sl_Failure_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SL_Failure_r16_sl_Failure_r16", err)
	}
	return nil
}

func (ie *SL_Failure_r16_sl_Failure_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SL_Failure_r16_sl_Failure_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
