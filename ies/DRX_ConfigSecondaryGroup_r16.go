package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DRX_ConfigSecondaryGroup_r16 struct {
	Drx_onDurationTimer_r16 DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16 `lb:1,ub:31,madatory`
	Drx_InactivityTimer_r16 DRX_ConfigSecondaryGroup_r16_drx_InactivityTimer_r16 `madatory`
}

func (ie *DRX_ConfigSecondaryGroup_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Drx_onDurationTimer_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Drx_onDurationTimer_r16", err)
	}
	if err = ie.Drx_InactivityTimer_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Drx_InactivityTimer_r16", err)
	}
	return nil
}

func (ie *DRX_ConfigSecondaryGroup_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Drx_onDurationTimer_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Drx_onDurationTimer_r16", err)
	}
	if err = ie.Drx_InactivityTimer_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Drx_InactivityTimer_r16", err)
	}
	return nil
}
