package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_unifiedJointTCI_mTRP_InterCell_BM_r17 struct {
	MaxNumAdditionalPCI_L1_RSRP_r17        int64                                                                                               `lb:1,ub:7,madatory`
	MaxNumSSB_ResourceL1_RSRP_AcrossCC_r17 MIMO_ParametersPerBand_unifiedJointTCI_mTRP_InterCell_BM_r17_maxNumSSB_ResourceL1_RSRP_AcrossCC_r17 `madatory`
}

func (ie *MIMO_ParametersPerBand_unifiedJointTCI_mTRP_InterCell_BM_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.MaxNumAdditionalPCI_L1_RSRP_r17, &uper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumAdditionalPCI_L1_RSRP_r17", err)
	}
	if err = ie.MaxNumSSB_ResourceL1_RSRP_AcrossCC_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumSSB_ResourceL1_RSRP_AcrossCC_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_unifiedJointTCI_mTRP_InterCell_BM_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_MaxNumAdditionalPCI_L1_RSRP_r17 int64
	if tmp_int_MaxNumAdditionalPCI_L1_RSRP_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumAdditionalPCI_L1_RSRP_r17", err)
	}
	ie.MaxNumAdditionalPCI_L1_RSRP_r17 = tmp_int_MaxNumAdditionalPCI_L1_RSRP_r17
	if err = ie.MaxNumSSB_ResourceL1_RSRP_AcrossCC_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumSSB_ResourceL1_RSRP_AcrossCC_r17", err)
	}
	return nil
}
