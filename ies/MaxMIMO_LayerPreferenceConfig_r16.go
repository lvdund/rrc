package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MaxMIMO_LayerPreferenceConfig_r16 struct {
	MaxMIMO_LayerPreferenceProhibitTimer_r16 MaxMIMO_LayerPreferenceConfig_r16_maxMIMO_LayerPreferenceProhibitTimer_r16 `madatory`
}

func (ie *MaxMIMO_LayerPreferenceConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.MaxMIMO_LayerPreferenceProhibitTimer_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxMIMO_LayerPreferenceProhibitTimer_r16", err)
	}
	return nil
}

func (ie *MaxMIMO_LayerPreferenceConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.MaxMIMO_LayerPreferenceProhibitTimer_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxMIMO_LayerPreferenceProhibitTimer_r16", err)
	}
	return nil
}
