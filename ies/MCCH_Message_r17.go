package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MCCH_Message_r17 struct {
	Message MCCH_MessageType_r17 `madatory`
}

func (ie *MCCH_Message_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Message.Encode(w); err != nil {
		return utils.WrapError("Encode Message", err)
	}
	return nil
}

func (ie *MCCH_Message_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Message.Decode(r); err != nil {
		return utils.WrapError("Decode Message", err)
	}
	return nil
}
