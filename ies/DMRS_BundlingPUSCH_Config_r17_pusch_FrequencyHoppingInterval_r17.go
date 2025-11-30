package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17_Enum_s2  aper.Enumerated = 0
	DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17_Enum_s4  aper.Enumerated = 1
	DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17_Enum_s5  aper.Enumerated = 2
	DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17_Enum_s6  aper.Enumerated = 3
	DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17_Enum_s8  aper.Enumerated = 4
	DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17_Enum_s10 aper.Enumerated = 5
	DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17_Enum_s12 aper.Enumerated = 6
	DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17_Enum_s14 aper.Enumerated = 7
	DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17_Enum_s16 aper.Enumerated = 8
	DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17_Enum_s20 aper.Enumerated = 9
)

type DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17 struct {
	Value aper.Enumerated `lb:0,ub:9,madatory`
}

func (ie *DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("Encode DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17", err)
	}
	return nil
}

func (ie *DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("Decode DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
