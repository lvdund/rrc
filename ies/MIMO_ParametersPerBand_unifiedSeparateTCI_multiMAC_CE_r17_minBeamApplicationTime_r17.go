package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n1   aper.Enumerated = 0
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n2   aper.Enumerated = 1
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n4   aper.Enumerated = 2
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n7   aper.Enumerated = 3
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n14  aper.Enumerated = 4
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n28  aper.Enumerated = 5
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n42  aper.Enumerated = 6
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n56  aper.Enumerated = 7
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n70  aper.Enumerated = 8
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n84  aper.Enumerated = 9
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n98  aper.Enumerated = 10
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n112 aper.Enumerated = 11
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n224 aper.Enumerated = 12
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n336 aper.Enumerated = 13
)

type MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17 struct {
	Value aper.Enumerated `lb:0,ub:13,madatory`
}

func (ie *MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("Encode MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("Decode MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
