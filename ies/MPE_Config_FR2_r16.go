package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MPE_Config_FR2_r16 struct {
	Mpe_ProhibitTimer_r16 MPE_Config_FR2_r16_mpe_ProhibitTimer_r16 `madatory`
	Mpe_Threshold_r16     MPE_Config_FR2_r16_mpe_Threshold_r16     `madatory`
}

func (ie *MPE_Config_FR2_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Mpe_ProhibitTimer_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Mpe_ProhibitTimer_r16", err)
	}
	if err = ie.Mpe_Threshold_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Mpe_Threshold_r16", err)
	}
	return nil
}

func (ie *MPE_Config_FR2_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Mpe_ProhibitTimer_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Mpe_ProhibitTimer_r16", err)
	}
	if err = ie.Mpe_Threshold_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Mpe_Threshold_r16", err)
	}
	return nil
}
