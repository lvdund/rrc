package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UAC_AccessCategory1_SelectionAssistanceInfo_Enum_a aper.Enumerated = 0
	UAC_AccessCategory1_SelectionAssistanceInfo_Enum_b aper.Enumerated = 1
	UAC_AccessCategory1_SelectionAssistanceInfo_Enum_c aper.Enumerated = 2
)

type UAC_AccessCategory1_SelectionAssistanceInfo struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *UAC_AccessCategory1_SelectionAssistanceInfo) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode UAC_AccessCategory1_SelectionAssistanceInfo", err)
	}
	return nil
}

func (ie *UAC_AccessCategory1_SelectionAssistanceInfo) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode UAC_AccessCategory1_SelectionAssistanceInfo", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
