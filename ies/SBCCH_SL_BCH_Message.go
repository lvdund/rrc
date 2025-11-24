package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SBCCH_SL_BCH_Message struct {
	Message SBCCH_SL_BCH_MessageType `madatory`
}

func (ie *SBCCH_SL_BCH_Message) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Message.Encode(w); err != nil {
		return utils.WrapError("Encode Message", err)
	}
	return nil
}

func (ie *SBCCH_SL_BCH_Message) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Message.Decode(r); err != nil {
		return utils.WrapError("Decode Message", err)
	}
	return nil
}
