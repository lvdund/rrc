package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SemiStaticChannelAccessConfig_r16 struct {
	Period SemiStaticChannelAccessConfig_r16_period `madatory`
}

func (ie *SemiStaticChannelAccessConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Period.Encode(w); err != nil {
		return utils.WrapError("Encode Period", err)
	}
	return nil
}

func (ie *SemiStaticChannelAccessConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Period.Decode(r); err != nil {
		return utils.WrapError("Decode Period", err)
	}
	return nil
}
