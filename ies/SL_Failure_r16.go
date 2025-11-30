package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_Failure_r16 struct {
	Sl_DestinationIdentity_r16 SL_DestinationIdentity_r16    `madatory`
	Sl_Failure_r16             SL_Failure_r16_sl_Failure_r16 `madatory`
}

func (ie *SL_Failure_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Sl_DestinationIdentity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_DestinationIdentity_r16", err)
	}
	if err = ie.Sl_Failure_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_Failure_r16", err)
	}
	return nil
}

func (ie *SL_Failure_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Sl_DestinationIdentity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_DestinationIdentity_r16", err)
	}
	if err = ie.Sl_Failure_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_Failure_r16", err)
	}
	return nil
}
