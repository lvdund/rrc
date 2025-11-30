package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RMTC_Config_r16_ref_SCS_CP_r16_Enum_kHz15     aper.Enumerated = 0
	RMTC_Config_r16_ref_SCS_CP_r16_Enum_kHz30     aper.Enumerated = 1
	RMTC_Config_r16_ref_SCS_CP_r16_Enum_kHz60_NCP aper.Enumerated = 2
	RMTC_Config_r16_ref_SCS_CP_r16_Enum_kHz60_ECP aper.Enumerated = 3
)

type RMTC_Config_r16_ref_SCS_CP_r16 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *RMTC_Config_r16_ref_SCS_CP_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode RMTC_Config_r16_ref_SCS_CP_r16", err)
	}
	return nil
}

func (ie *RMTC_Config_r16_ref_SCS_CP_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode RMTC_Config_r16_ref_SCS_CP_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
