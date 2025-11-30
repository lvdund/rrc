package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FailureReportSCG_failureType_v1610_Enum_scg_lbtFailure_r16             aper.Enumerated = 0
	FailureReportSCG_failureType_v1610_Enum_beamFailureRecoveryFailure_r16 aper.Enumerated = 1
	FailureReportSCG_failureType_v1610_Enum_t312_Expiry_r16                aper.Enumerated = 2
	FailureReportSCG_failureType_v1610_Enum_bh_RLF_r16                     aper.Enumerated = 3
	FailureReportSCG_failureType_v1610_Enum_beamFailure_r17                aper.Enumerated = 4
	FailureReportSCG_failureType_v1610_Enum_spare3                         aper.Enumerated = 5
	FailureReportSCG_failureType_v1610_Enum_spare2                         aper.Enumerated = 6
	FailureReportSCG_failureType_v1610_Enum_spare1                         aper.Enumerated = 7
)

type FailureReportSCG_failureType_v1610 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *FailureReportSCG_failureType_v1610) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode FailureReportSCG_failureType_v1610", err)
	}
	return nil
}

func (ie *FailureReportSCG_failureType_v1610) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode FailureReportSCG_failureType_v1610", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
