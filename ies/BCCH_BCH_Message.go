package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BCCH_BCH_Message struct {
	Message BCCH_BCH_MessageType `madatory`
}

func (ie *BCCH_BCH_Message) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Message.Encode(w); err != nil {
		return utils.WrapError("Encode Message", err)
	}
	return nil
}

func (ie *BCCH_BCH_Message) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Message.Decode(r); err != nil {
		return utils.WrapError("Decode Message", err)
	}
	return nil
}
