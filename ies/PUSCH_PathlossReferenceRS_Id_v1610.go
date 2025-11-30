package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_PathlossReferenceRS_Id_v1610 struct {
	Value uint64 `lb:maxNrofPUSCH_PathlossReferenceRSs,ub:maxNrofPUSCH_PathlossReferenceRSs_1_r16,madatory`
}

func (ie *PUSCH_PathlossReferenceRS_Id_v1610) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: maxNrofPUSCH_PathlossReferenceRSs, Ub: maxNrofPUSCH_PathlossReferenceRSs_1_r16}, false); err != nil {
		return utils.WrapError("Encode PUSCH_PathlossReferenceRS_Id_v1610", err)
	}
	return nil
}

func (ie *PUSCH_PathlossReferenceRS_Id_v1610) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: maxNrofPUSCH_PathlossReferenceRSs, Ub: maxNrofPUSCH_PathlossReferenceRSs_1_r16}, false); err != nil {
		return utils.WrapError("Decode PUSCH_PathlossReferenceRS_Id_v1610", err)
	}
	ie.Value = uint64(v)
	return nil
}
