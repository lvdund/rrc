package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MaxCC_PreferenceConfig_r16 struct {
	MaxCC_PreferenceProhibitTimer_r16 MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16 `madatory`
}

func (ie *MaxCC_PreferenceConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.MaxCC_PreferenceProhibitTimer_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxCC_PreferenceProhibitTimer_r16", err)
	}
	return nil
}

func (ie *MaxCC_PreferenceConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.MaxCC_PreferenceProhibitTimer_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxCC_PreferenceProhibitTimer_r16", err)
	}
	return nil
}
