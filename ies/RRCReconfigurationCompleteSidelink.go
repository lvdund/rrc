package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfigurationCompleteSidelink struct {
	Rrc_TransactionIdentifier_r16 RRC_TransactionIdentifier                             `madatory`
	CriticalExtensions            RRCReconfigurationCompleteSidelink_CriticalExtensions `madatory`
}

func (ie *RRCReconfigurationCompleteSidelink) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Rrc_TransactionIdentifier_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Rrc_TransactionIdentifier_r16", err)
	}
	if err = ie.CriticalExtensions.Encode(w); err != nil {
		return utils.WrapError("Encode CriticalExtensions", err)
	}
	return nil
}

func (ie *RRCReconfigurationCompleteSidelink) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Rrc_TransactionIdentifier_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Rrc_TransactionIdentifier_r16", err)
	}
	if err = ie.CriticalExtensions.Decode(r); err != nil {
		return utils.WrapError("Decode CriticalExtensions", err)
	}
	return nil
}
