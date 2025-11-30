package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_PathlossReferenceRS_Id_v1610 struct {
	Value uint64 `lb:maxNrofPUCCH_PathlossReferenceRSs,ub:maxNrofPUCCH_PathlossReferenceRSs_1_r16,madatory`
}

func (ie *PUCCH_PathlossReferenceRS_Id_v1610) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: maxNrofPUCCH_PathlossReferenceRSs, Ub: maxNrofPUCCH_PathlossReferenceRSs_1_r16}, false); err != nil {
		return utils.WrapError("Encode PUCCH_PathlossReferenceRS_Id_v1610", err)
	}
	return nil
}

func (ie *PUCCH_PathlossReferenceRS_Id_v1610) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: maxNrofPUCCH_PathlossReferenceRSs, Ub: maxNrofPUCCH_PathlossReferenceRSs_1_r16}, false); err != nil {
		return utils.WrapError("Decode PUCCH_PathlossReferenceRS_Id_v1610", err)
	}
	ie.Value = uint64(v)
	return nil
}
