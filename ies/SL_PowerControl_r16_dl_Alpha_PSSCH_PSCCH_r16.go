package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_PowerControl_r16_dl_Alpha_PSSCH_PSCCH_r16_Enum_alpha0  aper.Enumerated = 0
	SL_PowerControl_r16_dl_Alpha_PSSCH_PSCCH_r16_Enum_alpha04 aper.Enumerated = 1
	SL_PowerControl_r16_dl_Alpha_PSSCH_PSCCH_r16_Enum_alpha05 aper.Enumerated = 2
	SL_PowerControl_r16_dl_Alpha_PSSCH_PSCCH_r16_Enum_alpha06 aper.Enumerated = 3
	SL_PowerControl_r16_dl_Alpha_PSSCH_PSCCH_r16_Enum_alpha07 aper.Enumerated = 4
	SL_PowerControl_r16_dl_Alpha_PSSCH_PSCCH_r16_Enum_alpha08 aper.Enumerated = 5
	SL_PowerControl_r16_dl_Alpha_PSSCH_PSCCH_r16_Enum_alpha09 aper.Enumerated = 6
	SL_PowerControl_r16_dl_Alpha_PSSCH_PSCCH_r16_Enum_alpha1  aper.Enumerated = 7
)

type SL_PowerControl_r16_dl_Alpha_PSSCH_PSCCH_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SL_PowerControl_r16_dl_Alpha_PSSCH_PSCCH_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SL_PowerControl_r16_dl_Alpha_PSSCH_PSCCH_r16", err)
	}
	return nil
}

func (ie *SL_PowerControl_r16_dl_Alpha_PSSCH_PSCCH_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SL_PowerControl_r16_dl_Alpha_PSSCH_PSCCH_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
