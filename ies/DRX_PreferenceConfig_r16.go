package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DRX_PreferenceConfig_r16 struct {
	Drx_PreferenceProhibitTimer_r16 DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16 `madatory`
}

func (ie *DRX_PreferenceConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Drx_PreferenceProhibitTimer_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Drx_PreferenceProhibitTimer_r16", err)
	}
	return nil
}

func (ie *DRX_PreferenceConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Drx_PreferenceProhibitTimer_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Drx_PreferenceProhibitTimer_r16", err)
	}
	return nil
}
