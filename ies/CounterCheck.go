package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CounterCheck struct {
	Rrc_TransactionIdentifier RRC_TransactionIdentifier       `madatory`
	CriticalExtensions        CounterCheck_CriticalExtensions `madatory`
}

func (ie *CounterCheck) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Rrc_TransactionIdentifier.Encode(w); err != nil {
		return utils.WrapError("Encode Rrc_TransactionIdentifier", err)
	}
	if err = ie.CriticalExtensions.Encode(w); err != nil {
		return utils.WrapError("Encode CriticalExtensions", err)
	}
	return nil
}

func (ie *CounterCheck) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Rrc_TransactionIdentifier.Decode(r); err != nil {
		return utils.WrapError("Decode Rrc_TransactionIdentifier", err)
	}
	if err = ie.CriticalExtensions.Decode(r); err != nil {
		return utils.WrapError("Decode CriticalExtensions", err)
	}
	return nil
}
