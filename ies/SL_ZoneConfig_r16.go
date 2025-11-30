package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_ZoneConfig_r16 struct {
	Sl_ZoneLength_r16 SL_ZoneConfig_r16_sl_ZoneLength_r16 `madatory`
}

func (ie *SL_ZoneConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Sl_ZoneLength_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_ZoneLength_r16", err)
	}
	return nil
}

func (ie *SL_ZoneConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Sl_ZoneLength_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_ZoneLength_r16", err)
	}
	return nil
}
