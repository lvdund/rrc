package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_mTRP_inter_Cell_r17 struct {
	MaxNumAdditionalPCI_Case1_r17 int64 `lb:1,ub:7,madatory`
	MaxNumAdditionalPCI_Case2_r17 int64 `lb:0,ub:7,madatory`
}

func (ie *MIMO_ParametersPerBand_mTRP_inter_Cell_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.MaxNumAdditionalPCI_Case1_r17, &aper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumAdditionalPCI_Case1_r17", err)
	}
	if err = w.WriteInteger(ie.MaxNumAdditionalPCI_Case2_r17, &aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumAdditionalPCI_Case2_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_mTRP_inter_Cell_r17) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_MaxNumAdditionalPCI_Case1_r17 int64
	if tmp_int_MaxNumAdditionalPCI_Case1_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumAdditionalPCI_Case1_r17", err)
	}
	ie.MaxNumAdditionalPCI_Case1_r17 = tmp_int_MaxNumAdditionalPCI_Case1_r17
	var tmp_int_MaxNumAdditionalPCI_Case2_r17 int64
	if tmp_int_MaxNumAdditionalPCI_Case2_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumAdditionalPCI_Case2_r17", err)
	}
	ie.MaxNumAdditionalPCI_Case2_r17 = tmp_int_MaxNumAdditionalPCI_Case2_r17
	return nil
}
