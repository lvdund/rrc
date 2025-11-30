package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BCCH_Config_modificationPeriodCoeff_Enum_n2  aper.Enumerated = 0
	BCCH_Config_modificationPeriodCoeff_Enum_n4  aper.Enumerated = 1
	BCCH_Config_modificationPeriodCoeff_Enum_n8  aper.Enumerated = 2
	BCCH_Config_modificationPeriodCoeff_Enum_n16 aper.Enumerated = 3
)

type BCCH_Config_modificationPeriodCoeff struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *BCCH_Config_modificationPeriodCoeff) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode BCCH_Config_modificationPeriodCoeff", err)
	}
	return nil
}

func (ie *BCCH_Config_modificationPeriodCoeff) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode BCCH_Config_modificationPeriodCoeff", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
