package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SCCH_Message struct {
	Message SCCH_MessageType `madatory`
}

func (ie *SCCH_Message) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Message.Encode(w); err != nil {
		return utils.WrapError("Encode Message", err)
	}
	return nil
}

func (ie *SCCH_Message) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Message.Decode(r); err != nil {
		return utils.WrapError("Decode Message", err)
	}
	return nil
}
