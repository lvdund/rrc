package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SuccessHO_Config_r17_thresholdPercentageT312_r17_Enum_p20    aper.Enumerated = 0
	SuccessHO_Config_r17_thresholdPercentageT312_r17_Enum_p40    aper.Enumerated = 1
	SuccessHO_Config_r17_thresholdPercentageT312_r17_Enum_p60    aper.Enumerated = 2
	SuccessHO_Config_r17_thresholdPercentageT312_r17_Enum_p80    aper.Enumerated = 3
	SuccessHO_Config_r17_thresholdPercentageT312_r17_Enum_spare4 aper.Enumerated = 4
	SuccessHO_Config_r17_thresholdPercentageT312_r17_Enum_spare3 aper.Enumerated = 5
	SuccessHO_Config_r17_thresholdPercentageT312_r17_Enum_spare2 aper.Enumerated = 6
	SuccessHO_Config_r17_thresholdPercentageT312_r17_Enum_spare1 aper.Enumerated = 7
)

type SuccessHO_Config_r17_thresholdPercentageT312_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SuccessHO_Config_r17_thresholdPercentageT312_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SuccessHO_Config_r17_thresholdPercentageT312_r17", err)
	}
	return nil
}

func (ie *SuccessHO_Config_r17_thresholdPercentageT312_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SuccessHO_Config_r17_thresholdPercentageT312_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
