package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigGenericTwoStepRA_r16_msgB_ResponseWindow_v1700_Enum_sl240  aper.Enumerated = 0
	RACH_ConfigGenericTwoStepRA_r16_msgB_ResponseWindow_v1700_Enum_sl640  aper.Enumerated = 1
	RACH_ConfigGenericTwoStepRA_r16_msgB_ResponseWindow_v1700_Enum_sl960  aper.Enumerated = 2
	RACH_ConfigGenericTwoStepRA_r16_msgB_ResponseWindow_v1700_Enum_sl1280 aper.Enumerated = 3
	RACH_ConfigGenericTwoStepRA_r16_msgB_ResponseWindow_v1700_Enum_sl1920 aper.Enumerated = 4
	RACH_ConfigGenericTwoStepRA_r16_msgB_ResponseWindow_v1700_Enum_sl2560 aper.Enumerated = 5
)

type RACH_ConfigGenericTwoStepRA_r16_msgB_ResponseWindow_v1700 struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *RACH_ConfigGenericTwoStepRA_r16_msgB_ResponseWindow_v1700) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigGenericTwoStepRA_r16_msgB_ResponseWindow_v1700", err)
	}
	return nil
}

func (ie *RACH_ConfigGenericTwoStepRA_r16_msgB_ResponseWindow_v1700) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigGenericTwoStepRA_r16_msgB_ResponseWindow_v1700", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
