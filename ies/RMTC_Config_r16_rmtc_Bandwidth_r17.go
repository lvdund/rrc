package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RMTC_Config_r16_rmtc_Bandwidth_r17_Enum_mhz100  aper.Enumerated = 0
	RMTC_Config_r16_rmtc_Bandwidth_r17_Enum_mhz400  aper.Enumerated = 1
	RMTC_Config_r16_rmtc_Bandwidth_r17_Enum_mhz800  aper.Enumerated = 2
	RMTC_Config_r16_rmtc_Bandwidth_r17_Enum_mhz1600 aper.Enumerated = 3
	RMTC_Config_r16_rmtc_Bandwidth_r17_Enum_mhz2000 aper.Enumerated = 4
)

type RMTC_Config_r16_rmtc_Bandwidth_r17 struct {
	Value aper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *RMTC_Config_r16_rmtc_Bandwidth_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode RMTC_Config_r16_rmtc_Bandwidth_r17", err)
	}
	return nil
}

func (ie *RMTC_Config_r16_rmtc_Bandwidth_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode RMTC_Config_r16_rmtc_Bandwidth_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
