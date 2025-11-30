package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA_failureTypeEUTRA_Enum_t313_Expiry         aper.Enumerated = 0
	CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA_failureTypeEUTRA_Enum_randomAccessProblem aper.Enumerated = 1
	CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA_failureTypeEUTRA_Enum_rlc_MaxNumRetx      aper.Enumerated = 2
	CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA_failureTypeEUTRA_Enum_scg_ChangeFailure   aper.Enumerated = 3
)

type CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA_failureTypeEUTRA struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA_failureTypeEUTRA) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA_failureTypeEUTRA", err)
	}
	return nil
}

func (ie *CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA_failureTypeEUTRA) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA_failureTypeEUTRA", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
