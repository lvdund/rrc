package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FailureReportSCG_EUTRA_failureType_Enum_t313_Expiry         aper.Enumerated = 0
	FailureReportSCG_EUTRA_failureType_Enum_randomAccessProblem aper.Enumerated = 1
	FailureReportSCG_EUTRA_failureType_Enum_rlc_MaxNumRetx      aper.Enumerated = 2
	FailureReportSCG_EUTRA_failureType_Enum_scg_ChangeFailure   aper.Enumerated = 3
	FailureReportSCG_EUTRA_failureType_Enum_spare4              aper.Enumerated = 4
	FailureReportSCG_EUTRA_failureType_Enum_spare3              aper.Enumerated = 5
	FailureReportSCG_EUTRA_failureType_Enum_spare2              aper.Enumerated = 6
	FailureReportSCG_EUTRA_failureType_Enum_spare1              aper.Enumerated = 7
)

type FailureReportSCG_EUTRA_failureType struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *FailureReportSCG_EUTRA_failureType) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode FailureReportSCG_EUTRA_failureType", err)
	}
	return nil
}

func (ie *FailureReportSCG_EUTRA_failureType) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode FailureReportSCG_EUTRA_failureType", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
