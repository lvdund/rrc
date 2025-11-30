package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIMO_ParametersPerBand_unifiedJointTCI_r17_maxConfiguredJointTCI_r17_Enum_n8   aper.Enumerated = 0
	MIMO_ParametersPerBand_unifiedJointTCI_r17_maxConfiguredJointTCI_r17_Enum_n12  aper.Enumerated = 1
	MIMO_ParametersPerBand_unifiedJointTCI_r17_maxConfiguredJointTCI_r17_Enum_n16  aper.Enumerated = 2
	MIMO_ParametersPerBand_unifiedJointTCI_r17_maxConfiguredJointTCI_r17_Enum_n24  aper.Enumerated = 3
	MIMO_ParametersPerBand_unifiedJointTCI_r17_maxConfiguredJointTCI_r17_Enum_n32  aper.Enumerated = 4
	MIMO_ParametersPerBand_unifiedJointTCI_r17_maxConfiguredJointTCI_r17_Enum_n48  aper.Enumerated = 5
	MIMO_ParametersPerBand_unifiedJointTCI_r17_maxConfiguredJointTCI_r17_Enum_n64  aper.Enumerated = 6
	MIMO_ParametersPerBand_unifiedJointTCI_r17_maxConfiguredJointTCI_r17_Enum_n128 aper.Enumerated = 7
)

type MIMO_ParametersPerBand_unifiedJointTCI_r17_maxConfiguredJointTCI_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *MIMO_ParametersPerBand_unifiedJointTCI_r17_maxConfiguredJointTCI_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode MIMO_ParametersPerBand_unifiedJointTCI_r17_maxConfiguredJointTCI_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_unifiedJointTCI_r17_maxConfiguredJointTCI_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode MIMO_ParametersPerBand_unifiedJointTCI_r17_maxConfiguredJointTCI_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
