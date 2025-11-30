package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17_Enum_p00 aper.Enumerated = 0
	UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17_Enum_p05 aper.Enumerated = 1
	UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17_Enum_p10 aper.Enumerated = 2
	UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17_Enum_p15 aper.Enumerated = 3
	UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17_Enum_p20 aper.Enumerated = 4
	UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17_Enum_p25 aper.Enumerated = 5
	UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17_Enum_p30 aper.Enumerated = 6
	UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17_Enum_p40 aper.Enumerated = 7
	UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17_Enum_p50 aper.Enumerated = 8
	UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17_Enum_p60 aper.Enumerated = 9
	UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17_Enum_p70 aper.Enumerated = 10
	UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17_Enum_p75 aper.Enumerated = 11
	UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17_Enum_p80 aper.Enumerated = 12
	UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17_Enum_p85 aper.Enumerated = 13
	UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17_Enum_p90 aper.Enumerated = 14
	UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17_Enum_p95 aper.Enumerated = 15
)

type UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17", err)
	}
	return nil
}

func (ie *UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
