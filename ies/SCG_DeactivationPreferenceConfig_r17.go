package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SCG_DeactivationPreferenceConfig_r17 struct {
	Scg_DeactivationPreferenceProhibitTimer_r17 SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17 `madatory`
}

func (ie *SCG_DeactivationPreferenceConfig_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Scg_DeactivationPreferenceProhibitTimer_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Scg_DeactivationPreferenceProhibitTimer_r17", err)
	}
	return nil
}

func (ie *SCG_DeactivationPreferenceConfig_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Scg_DeactivationPreferenceProhibitTimer_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Scg_DeactivationPreferenceProhibitTimer_r17", err)
	}
	return nil
}
