package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MaxBW_PreferenceConfig_r16 struct {
	MaxBW_PreferenceProhibitTimer_r16 MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16 `madatory`
}

func (ie *MaxBW_PreferenceConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.MaxBW_PreferenceProhibitTimer_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxBW_PreferenceProhibitTimer_r16", err)
	}
	return nil
}

func (ie *MaxBW_PreferenceConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.MaxBW_PreferenceProhibitTimer_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxBW_PreferenceProhibitTimer_r16", err)
	}
	return nil
}
