package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_MeasConfigInfo_r16 struct {
	Sl_DestinationIndex_r16 SL_DestinationIndex_r16 `madatory`
	Sl_MeasConfig_r16       SL_MeasConfig_r16       `madatory`
}

func (ie *SL_MeasConfigInfo_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Sl_DestinationIndex_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_DestinationIndex_r16", err)
	}
	if err = ie.Sl_MeasConfig_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_MeasConfig_r16", err)
	}
	return nil
}

func (ie *SL_MeasConfigInfo_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Sl_DestinationIndex_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_DestinationIndex_r16", err)
	}
	if err = ie.Sl_MeasConfig_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_MeasConfig_r16", err)
	}
	return nil
}
