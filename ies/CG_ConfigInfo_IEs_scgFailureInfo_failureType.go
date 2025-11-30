package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CG_ConfigInfo_IEs_scgFailureInfo_failureType_Enum_t310_Expiry              aper.Enumerated = 0
	CG_ConfigInfo_IEs_scgFailureInfo_failureType_Enum_randomAccessProblem      aper.Enumerated = 1
	CG_ConfigInfo_IEs_scgFailureInfo_failureType_Enum_rlc_MaxNumRetx           aper.Enumerated = 2
	CG_ConfigInfo_IEs_scgFailureInfo_failureType_Enum_synchReconfigFailure_SCG aper.Enumerated = 3
	CG_ConfigInfo_IEs_scgFailureInfo_failureType_Enum_scg_reconfigFailure      aper.Enumerated = 4
	CG_ConfigInfo_IEs_scgFailureInfo_failureType_Enum_srb3_IntegrityFailure    aper.Enumerated = 5
)

type CG_ConfigInfo_IEs_scgFailureInfo_failureType struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *CG_ConfigInfo_IEs_scgFailureInfo_failureType) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode CG_ConfigInfo_IEs_scgFailureInfo_failureType", err)
	}
	return nil
}

func (ie *CG_ConfigInfo_IEs_scgFailureInfo_failureType) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode CG_ConfigInfo_IEs_scgFailureInfo_failureType", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
