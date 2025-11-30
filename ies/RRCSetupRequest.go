package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RRCSetupRequest struct {
	RrcSetupRequest RRCSetupRequest_IEs `madatory`
}

func (ie *RRCSetupRequest) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.RrcSetupRequest.Encode(w); err != nil {
		return utils.WrapError("Encode RrcSetupRequest", err)
	}
	return nil
}

func (ie *RRCSetupRequest) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.RrcSetupRequest.Decode(r); err != nil {
		return utils.WrapError("Decode RrcSetupRequest", err)
	}
	return nil
}
