package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReestablishmentRequest struct {
	RrcReestablishmentRequest RRCReestablishmentRequest_IEs `madatory`
}

func (ie *RRCReestablishmentRequest) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.RrcReestablishmentRequest.Encode(w); err != nil {
		return utils.WrapError("Encode RrcReestablishmentRequest", err)
	}
	return nil
}

func (ie *RRCReestablishmentRequest) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.RrcReestablishmentRequest.Decode(r); err != nil {
		return utils.WrapError("Decode RrcReestablishmentRequest", err)
	}
	return nil
}
