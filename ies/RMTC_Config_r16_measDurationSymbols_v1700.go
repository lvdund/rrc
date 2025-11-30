package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RMTC_Config_r16_measDurationSymbols_v1700_Enum_sym140  aper.Enumerated = 0
	RMTC_Config_r16_measDurationSymbols_v1700_Enum_sym560  aper.Enumerated = 1
	RMTC_Config_r16_measDurationSymbols_v1700_Enum_sym1120 aper.Enumerated = 2
)

type RMTC_Config_r16_measDurationSymbols_v1700 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *RMTC_Config_r16_measDurationSymbols_v1700) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode RMTC_Config_r16_measDurationSymbols_v1700", err)
	}
	return nil
}

func (ie *RMTC_Config_r16_measDurationSymbols_v1700) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode RMTC_Config_r16_measDurationSymbols_v1700", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
