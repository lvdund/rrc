package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FailureReportSCG_failureType_Enum_t310_Expiry             aper.Enumerated = 0
	FailureReportSCG_failureType_Enum_randomAccessProblem     aper.Enumerated = 1
	FailureReportSCG_failureType_Enum_rlc_MaxNumRetx          aper.Enumerated = 2
	FailureReportSCG_failureType_Enum_synchReconfigFailureSCG aper.Enumerated = 3
	FailureReportSCG_failureType_Enum_scg_ReconfigFailure     aper.Enumerated = 4
	FailureReportSCG_failureType_Enum_srb3_IntegrityFailure   aper.Enumerated = 5
	FailureReportSCG_failureType_Enum_other_r16               aper.Enumerated = 6
	FailureReportSCG_failureType_Enum_spare1                  aper.Enumerated = 7
)

type FailureReportSCG_failureType struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *FailureReportSCG_failureType) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode FailureReportSCG_failureType", err)
	}
	return nil
}

func (ie *FailureReportSCG_failureType) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode FailureReportSCG_failureType", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
