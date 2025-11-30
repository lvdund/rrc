package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MCCH_Config_r17_mcch_WindowDuration_r17_Enum_sl2   aper.Enumerated = 0
	MCCH_Config_r17_mcch_WindowDuration_r17_Enum_sl4   aper.Enumerated = 1
	MCCH_Config_r17_mcch_WindowDuration_r17_Enum_sl8   aper.Enumerated = 2
	MCCH_Config_r17_mcch_WindowDuration_r17_Enum_sl10  aper.Enumerated = 3
	MCCH_Config_r17_mcch_WindowDuration_r17_Enum_sl20  aper.Enumerated = 4
	MCCH_Config_r17_mcch_WindowDuration_r17_Enum_sl40  aper.Enumerated = 5
	MCCH_Config_r17_mcch_WindowDuration_r17_Enum_sl80  aper.Enumerated = 6
	MCCH_Config_r17_mcch_WindowDuration_r17_Enum_sl160 aper.Enumerated = 7
)

type MCCH_Config_r17_mcch_WindowDuration_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *MCCH_Config_r17_mcch_WindowDuration_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode MCCH_Config_r17_mcch_WindowDuration_r17", err)
	}
	return nil
}

func (ie *MCCH_Config_r17_mcch_WindowDuration_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode MCCH_Config_r17_mcch_WindowDuration_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
