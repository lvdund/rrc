package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ReleasePreference_r16 struct {
	PreferredRRC_State_r16 ReleasePreference_r16_preferredRRC_State_r16 `madatory`
}

func (ie *ReleasePreference_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.PreferredRRC_State_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PreferredRRC_State_r16", err)
	}
	return nil
}

func (ie *ReleasePreference_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.PreferredRRC_State_r16.Decode(r); err != nil {
		return utils.WrapError("Decode PreferredRRC_State_r16", err)
	}
	return nil
}
