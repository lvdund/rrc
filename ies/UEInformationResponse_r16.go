package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UEInformationResponse_r16 struct {
	Rrc_TransactionIdentifier RRC_TransactionIdentifier                    `madatory`
	CriticalExtensions        UEInformationResponse_r16_CriticalExtensions `madatory`
}

func (ie *UEInformationResponse_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Rrc_TransactionIdentifier.Encode(w); err != nil {
		return utils.WrapError("Encode Rrc_TransactionIdentifier", err)
	}
	if err = ie.CriticalExtensions.Encode(w); err != nil {
		return utils.WrapError("Encode CriticalExtensions", err)
	}
	return nil
}

func (ie *UEInformationResponse_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Rrc_TransactionIdentifier.Decode(r); err != nil {
		return utils.WrapError("Decode Rrc_TransactionIdentifier", err)
	}
	if err = ie.CriticalExtensions.Decode(r); err != nil {
		return utils.WrapError("Decode CriticalExtensions", err)
	}
	return nil
}
