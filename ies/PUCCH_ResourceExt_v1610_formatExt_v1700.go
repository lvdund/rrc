package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_ResourceExt_v1610_formatExt_v1700 struct {
	NrofPRBs_r17 int64 `lb:1,ub:16,madatory`
}

func (ie *PUCCH_ResourceExt_v1610_formatExt_v1700) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.NrofPRBs_r17, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger NrofPRBs_r17", err)
	}
	return nil
}

func (ie *PUCCH_ResourceExt_v1610_formatExt_v1700) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_NrofPRBs_r17 int64
	if tmp_int_NrofPRBs_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger NrofPRBs_r17", err)
	}
	ie.NrofPRBs_r17 = tmp_int_NrofPRBs_r17
	return nil
}
