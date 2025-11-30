package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17_Enum_n1  aper.Enumerated = 0
	CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17_Enum_n2  aper.Enumerated = 1
	CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17_Enum_n3  aper.Enumerated = 2
	CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17_Enum_n4  aper.Enumerated = 3
	CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17_Enum_n6  aper.Enumerated = 4
	CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17_Enum_n8  aper.Enumerated = 5
	CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17_Enum_n16 aper.Enumerated = 6
	CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17_Enum_n32 aper.Enumerated = 7
)

type CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17", err)
	}
	return nil
}

func (ie *CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
