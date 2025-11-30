package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type OnDemandSIB_Request_r16 struct {
	OnDemandSIB_RequestProhibitTimer_r16 OnDemandSIB_Request_r16_onDemandSIB_RequestProhibitTimer_r16 `madatory`
}

func (ie *OnDemandSIB_Request_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.OnDemandSIB_RequestProhibitTimer_r16.Encode(w); err != nil {
		return utils.WrapError("Encode OnDemandSIB_RequestProhibitTimer_r16", err)
	}
	return nil
}

func (ie *OnDemandSIB_Request_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.OnDemandSIB_RequestProhibitTimer_r16.Decode(r); err != nil {
		return utils.WrapError("Decode OnDemandSIB_RequestProhibitTimer_r16", err)
	}
	return nil
}
