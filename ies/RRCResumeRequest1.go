package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RRCResumeRequest1 struct {
	RrcResumeRequest1 RRCResumeRequest1_IEs `madatory`
}

func (ie *RRCResumeRequest1) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.RrcResumeRequest1.Encode(w); err != nil {
		return utils.WrapError("Encode RrcResumeRequest1", err)
	}
	return nil
}

func (ie *RRCResumeRequest1) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.RrcResumeRequest1.Decode(r); err != nil {
		return utils.WrapError("Decode RrcResumeRequest1", err)
	}
	return nil
}
