package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UuMessageTransferSidelink_r17 struct {
	CriticalExtensions UuMessageTransferSidelink_r17_CriticalExtensions `madatory`
}

func (ie *UuMessageTransferSidelink_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.CriticalExtensions.Encode(w); err != nil {
		return utils.WrapError("Encode CriticalExtensions", err)
	}
	return nil
}

func (ie *UuMessageTransferSidelink_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.CriticalExtensions.Decode(r); err != nil {
		return utils.WrapError("Decode CriticalExtensions", err)
	}
	return nil
}
