package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_mTRP_PUSCH_CSI_RS_r17 struct {
	MaxNumPeriodicSRS_r17          int64 `lb:1,ub:8,madatory`
	MaxNumAperiodicSRS_r17         int64 `lb:1,ub:8,madatory`
	MaxNumSP_SRS_r17               int64 `lb:0,ub:8,madatory`
	NumSRS_ResourcePerCC_r17       int64 `lb:1,ub:16,madatory`
	NumSRS_ResourceNonCodebook_r17 int64 `lb:1,ub:2,madatory`
}

func (ie *MIMO_ParametersPerBand_mTRP_PUSCH_CSI_RS_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.MaxNumPeriodicSRS_r17, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumPeriodicSRS_r17", err)
	}
	if err = w.WriteInteger(ie.MaxNumAperiodicSRS_r17, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumAperiodicSRS_r17", err)
	}
	if err = w.WriteInteger(ie.MaxNumSP_SRS_r17, &uper.Constraint{Lb: 0, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumSP_SRS_r17", err)
	}
	if err = w.WriteInteger(ie.NumSRS_ResourcePerCC_r17, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger NumSRS_ResourcePerCC_r17", err)
	}
	if err = w.WriteInteger(ie.NumSRS_ResourceNonCodebook_r17, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteInteger NumSRS_ResourceNonCodebook_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_mTRP_PUSCH_CSI_RS_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_MaxNumPeriodicSRS_r17 int64
	if tmp_int_MaxNumPeriodicSRS_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumPeriodicSRS_r17", err)
	}
	ie.MaxNumPeriodicSRS_r17 = tmp_int_MaxNumPeriodicSRS_r17
	var tmp_int_MaxNumAperiodicSRS_r17 int64
	if tmp_int_MaxNumAperiodicSRS_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumAperiodicSRS_r17", err)
	}
	ie.MaxNumAperiodicSRS_r17 = tmp_int_MaxNumAperiodicSRS_r17
	var tmp_int_MaxNumSP_SRS_r17 int64
	if tmp_int_MaxNumSP_SRS_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumSP_SRS_r17", err)
	}
	ie.MaxNumSP_SRS_r17 = tmp_int_MaxNumSP_SRS_r17
	var tmp_int_NumSRS_ResourcePerCC_r17 int64
	if tmp_int_NumSRS_ResourcePerCC_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger NumSRS_ResourcePerCC_r17", err)
	}
	ie.NumSRS_ResourcePerCC_r17 = tmp_int_NumSRS_ResourcePerCC_r17
	var tmp_int_NumSRS_ResourceNonCodebook_r17 int64
	if tmp_int_NumSRS_ResourceNonCodebook_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadInteger NumSRS_ResourceNonCodebook_r17", err)
	}
	ie.NumSRS_ResourceNonCodebook_r17 = tmp_int_NumSRS_ResourceNonCodebook_r17
	return nil
}
