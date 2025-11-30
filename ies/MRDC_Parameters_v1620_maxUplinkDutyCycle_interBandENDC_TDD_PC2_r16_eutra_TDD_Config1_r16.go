package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config1_r16_Enum_n20  aper.Enumerated = 0
	MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config1_r16_Enum_n40  aper.Enumerated = 1
	MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config1_r16_Enum_n50  aper.Enumerated = 2
	MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config1_r16_Enum_n60  aper.Enumerated = 3
	MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config1_r16_Enum_n70  aper.Enumerated = 4
	MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config1_r16_Enum_n80  aper.Enumerated = 5
	MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config1_r16_Enum_n90  aper.Enumerated = 6
	MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config1_r16_Enum_n100 aper.Enumerated = 7
)

type MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config1_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config1_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config1_r16", err)
	}
	return nil
}

func (ie *MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config1_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config1_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
