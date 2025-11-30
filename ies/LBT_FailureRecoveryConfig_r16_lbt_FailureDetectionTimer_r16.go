package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	LBT_FailureRecoveryConfig_r16_lbt_FailureDetectionTimer_r16_Enum_ms10  aper.Enumerated = 0
	LBT_FailureRecoveryConfig_r16_lbt_FailureDetectionTimer_r16_Enum_ms20  aper.Enumerated = 1
	LBT_FailureRecoveryConfig_r16_lbt_FailureDetectionTimer_r16_Enum_ms40  aper.Enumerated = 2
	LBT_FailureRecoveryConfig_r16_lbt_FailureDetectionTimer_r16_Enum_ms80  aper.Enumerated = 3
	LBT_FailureRecoveryConfig_r16_lbt_FailureDetectionTimer_r16_Enum_ms160 aper.Enumerated = 4
	LBT_FailureRecoveryConfig_r16_lbt_FailureDetectionTimer_r16_Enum_ms320 aper.Enumerated = 5
)

type LBT_FailureRecoveryConfig_r16_lbt_FailureDetectionTimer_r16 struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *LBT_FailureRecoveryConfig_r16_lbt_FailureDetectionTimer_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode LBT_FailureRecoveryConfig_r16_lbt_FailureDetectionTimer_r16", err)
	}
	return nil
}

func (ie *LBT_FailureRecoveryConfig_r16_lbt_FailureDetectionTimer_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode LBT_FailureRecoveryConfig_r16_lbt_FailureDetectionTimer_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
