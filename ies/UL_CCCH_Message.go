package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UL_CCCH_Message struct {
	Message UL_CCCH_MessageType `madatory`
}

func (ie *UL_CCCH_Message) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Message.Encode(w); err != nil {
		return utils.WrapError("Encode Message", err)
	}
	return nil
}

func (ie *UL_CCCH_Message) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Message.Decode(r); err != nil {
		return utils.WrapError("Decode Message", err)
	}
	return nil
}
