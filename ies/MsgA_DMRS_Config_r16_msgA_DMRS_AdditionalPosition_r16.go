package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MsgA_DMRS_Config_r16_msgA_DMRS_AdditionalPosition_r16_Enum_pos0 aper.Enumerated = 0
	MsgA_DMRS_Config_r16_msgA_DMRS_AdditionalPosition_r16_Enum_pos1 aper.Enumerated = 1
	MsgA_DMRS_Config_r16_msgA_DMRS_AdditionalPosition_r16_Enum_pos3 aper.Enumerated = 2
)

type MsgA_DMRS_Config_r16_msgA_DMRS_AdditionalPosition_r16 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *MsgA_DMRS_Config_r16_msgA_DMRS_AdditionalPosition_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode MsgA_DMRS_Config_r16_msgA_DMRS_AdditionalPosition_r16", err)
	}
	return nil
}

func (ie *MsgA_DMRS_Config_r16_msgA_DMRS_AdditionalPosition_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode MsgA_DMRS_Config_r16_msgA_DMRS_AdditionalPosition_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
