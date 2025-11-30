package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MsgA_PUSCH_Resource_r16_msgA_Alpha_r16_Enum_alpha0  aper.Enumerated = 0
	MsgA_PUSCH_Resource_r16_msgA_Alpha_r16_Enum_alpha04 aper.Enumerated = 1
	MsgA_PUSCH_Resource_r16_msgA_Alpha_r16_Enum_alpha05 aper.Enumerated = 2
	MsgA_PUSCH_Resource_r16_msgA_Alpha_r16_Enum_alpha06 aper.Enumerated = 3
	MsgA_PUSCH_Resource_r16_msgA_Alpha_r16_Enum_alpha07 aper.Enumerated = 4
	MsgA_PUSCH_Resource_r16_msgA_Alpha_r16_Enum_alpha08 aper.Enumerated = 5
	MsgA_PUSCH_Resource_r16_msgA_Alpha_r16_Enum_alpha09 aper.Enumerated = 6
	MsgA_PUSCH_Resource_r16_msgA_Alpha_r16_Enum_alpha1  aper.Enumerated = 7
)

type MsgA_PUSCH_Resource_r16_msgA_Alpha_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *MsgA_PUSCH_Resource_r16_msgA_Alpha_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode MsgA_PUSCH_Resource_r16_msgA_Alpha_r16", err)
	}
	return nil
}

func (ie *MsgA_PUSCH_Resource_r16_msgA_Alpha_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode MsgA_PUSCH_Resource_r16_msgA_Alpha_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
