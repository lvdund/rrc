package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MinSchedulingOffsetPreferenceConfig_r16 struct {
	MinSchedulingOffsetPreferenceProhibitTimer_r16 MinSchedulingOffsetPreferenceConfig_r16_minSchedulingOffsetPreferenceProhibitTimer_r16 `madatory`
}

func (ie *MinSchedulingOffsetPreferenceConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.MinSchedulingOffsetPreferenceProhibitTimer_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MinSchedulingOffsetPreferenceProhibitTimer_r16", err)
	}
	return nil
}

func (ie *MinSchedulingOffsetPreferenceConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.MinSchedulingOffsetPreferenceProhibitTimer_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MinSchedulingOffsetPreferenceProhibitTimer_r16", err)
	}
	return nil
}
