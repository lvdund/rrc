package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16_failureType_r16_Enum_scg_lbtFailure_r16             aper.Enumerated = 0
	CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16_failureType_r16_Enum_beamFailureRecoveryFailure_r16 aper.Enumerated = 1
	CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16_failureType_r16_Enum_t312_Expiry_r16                aper.Enumerated = 2
	CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16_failureType_r16_Enum_bh_RLF_r16                     aper.Enumerated = 3
	CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16_failureType_r16_Enum_beamFailure_r17                aper.Enumerated = 4
	CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16_failureType_r16_Enum_spare3                         aper.Enumerated = 5
	CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16_failureType_r16_Enum_spare2                         aper.Enumerated = 6
	CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16_failureType_r16_Enum_spare1                         aper.Enumerated = 7
)

type CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16_failureType_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16_failureType_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16_failureType_r16", err)
	}
	return nil
}

func (ie *CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16_failureType_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16_failureType_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
