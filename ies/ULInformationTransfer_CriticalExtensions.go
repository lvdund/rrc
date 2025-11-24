package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ULInformationTransfer_CriticalExtensions_Choice_nothing uint64 = iota
	ULInformationTransfer_CriticalExtensions_Choice_UlInformationTransfer
	ULInformationTransfer_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type ULInformationTransfer_CriticalExtensions struct {
	Choice                   uint64
	UlInformationTransfer    *ULInformationTransfer_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *ULInformationTransfer_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ULInformationTransfer_CriticalExtensions_Choice_UlInformationTransfer:
		if err = ie.UlInformationTransfer.Encode(w); err != nil {
			err = utils.WrapError("Encode UlInformationTransfer", err)
		}
	case ULInformationTransfer_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ULInformationTransfer_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ULInformationTransfer_CriticalExtensions_Choice_UlInformationTransfer:
		ie.UlInformationTransfer = new(ULInformationTransfer_IEs)
		if err = ie.UlInformationTransfer.Decode(r); err != nil {
			return utils.WrapError("Decode UlInformationTransfer", err)
		}
	case ULInformationTransfer_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
