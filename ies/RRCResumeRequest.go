package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RRCResumeRequest struct {
	RrcResumeRequest RRCResumeRequest_IEs `madatory`
}

func (ie *RRCResumeRequest) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.RrcResumeRequest.Encode(w); err != nil {
		return utils.WrapError("Encode RrcResumeRequest", err)
	}
	return nil
}

func (ie *RRCResumeRequest) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.RrcResumeRequest.Decode(r); err != nil {
		return utils.WrapError("Decode RrcResumeRequest", err)
	}
	return nil
}
