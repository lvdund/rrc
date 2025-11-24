package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_mpe_Mitigation_r17 struct {
	MaxNumP_MPR_RI_pairs_r17 int64                                                      `lb:1,ub:4,madatory`
	MaxNumConfRS_r17         MIMO_ParametersPerBand_mpe_Mitigation_r17_maxNumConfRS_r17 `madatory`
}

func (ie *MIMO_ParametersPerBand_mpe_Mitigation_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.MaxNumP_MPR_RI_pairs_r17, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumP_MPR_RI_pairs_r17", err)
	}
	if err = ie.MaxNumConfRS_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumConfRS_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_mpe_Mitigation_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_MaxNumP_MPR_RI_pairs_r17 int64
	if tmp_int_MaxNumP_MPR_RI_pairs_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumP_MPR_RI_pairs_r17", err)
	}
	ie.MaxNumP_MPR_RI_pairs_r17 = tmp_int_MaxNumP_MPR_RI_pairs_r17
	if err = ie.MaxNumConfRS_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumConfRS_r17", err)
	}
	return nil
}
