package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCSetupRequest struct {
	RrcSetupRequest RRCSetupRequest_IEs `madatory`
}

func (ie *RRCSetupRequest) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.RrcSetupRequest.Encode(w); err != nil {
		return utils.WrapError("Encode RrcSetupRequest", err)
	}
	return nil
}

func (ie *RRCSetupRequest) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.RrcSetupRequest.Decode(r); err != nil {
		return utils.WrapError("Decode RrcSetupRequest", err)
	}
	return nil
}
