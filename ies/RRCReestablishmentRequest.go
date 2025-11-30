package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RRCReestablishmentRequest struct {
	RrcReestablishmentRequest RRCReestablishmentRequest_IEs `madatory`
}

func (ie *RRCReestablishmentRequest) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.RrcReestablishmentRequest.Encode(w); err != nil {
		return utils.WrapError("Encode RrcReestablishmentRequest", err)
	}
	return nil
}

func (ie *RRCReestablishmentRequest) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.RrcReestablishmentRequest.Decode(r); err != nil {
		return utils.WrapError("Decode RrcReestablishmentRequest", err)
	}
	return nil
}
