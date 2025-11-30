package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC_Enum_n2 aper.Enumerated = 0
	MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC_Enum_n3 aper.Enumerated = 1
	MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC_Enum_n4 aper.Enumerated = 2
	MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC_Enum_n5 aper.Enumerated = 3
	MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC_Enum_n6 aper.Enumerated = 4
	MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC_Enum_n7 aper.Enumerated = 5
	MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC_Enum_n8 aper.Enumerated = 6
)

type MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC struct {
	Value aper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
