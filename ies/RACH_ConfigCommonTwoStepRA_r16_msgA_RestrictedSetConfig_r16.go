package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigCommonTwoStepRA_r16_msgA_RestrictedSetConfig_r16_Enum_unrestrictedSet    aper.Enumerated = 0
	RACH_ConfigCommonTwoStepRA_r16_msgA_RestrictedSetConfig_r16_Enum_restrictedSetTypeA aper.Enumerated = 1
	RACH_ConfigCommonTwoStepRA_r16_msgA_RestrictedSetConfig_r16_Enum_restrictedSetTypeB aper.Enumerated = 2
)

type RACH_ConfigCommonTwoStepRA_r16_msgA_RestrictedSetConfig_r16 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *RACH_ConfigCommonTwoStepRA_r16_msgA_RestrictedSetConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigCommonTwoStepRA_r16_msgA_RestrictedSetConfig_r16", err)
	}
	return nil
}

func (ie *RACH_ConfigCommonTwoStepRA_r16_msgA_RestrictedSetConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigCommonTwoStepRA_r16_msgA_RestrictedSetConfig_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
