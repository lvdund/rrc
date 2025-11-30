package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UAC_AC1_SelectAssistInfo_r16_Enum_a             aper.Enumerated = 0
	UAC_AC1_SelectAssistInfo_r16_Enum_b             aper.Enumerated = 1
	UAC_AC1_SelectAssistInfo_r16_Enum_c             aper.Enumerated = 2
	UAC_AC1_SelectAssistInfo_r16_Enum_notConfigured aper.Enumerated = 3
)

type UAC_AC1_SelectAssistInfo_r16 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *UAC_AC1_SelectAssistInfo_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode UAC_AC1_SelectAssistInfo_r16", err)
	}
	return nil
}

func (ie *UAC_AC1_SelectAssistInfo_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode UAC_AC1_SelectAssistInfo_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
