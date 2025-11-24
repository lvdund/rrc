package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DLInformationTransfer_CriticalExtensions_Choice_nothing uint64 = iota
	DLInformationTransfer_CriticalExtensions_Choice_DlInformationTransfer
	DLInformationTransfer_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type DLInformationTransfer_CriticalExtensions struct {
	Choice                   uint64
	DlInformationTransfer    *DLInformationTransfer_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *DLInformationTransfer_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DLInformationTransfer_CriticalExtensions_Choice_DlInformationTransfer:
		if err = ie.DlInformationTransfer.Encode(w); err != nil {
			err = utils.WrapError("Encode DlInformationTransfer", err)
		}
	case DLInformationTransfer_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DLInformationTransfer_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DLInformationTransfer_CriticalExtensions_Choice_DlInformationTransfer:
		ie.DlInformationTransfer = new(DLInformationTransfer_IEs)
		if err = ie.DlInformationTransfer.Decode(r); err != nil {
			return utils.WrapError("Decode DlInformationTransfer", err)
		}
	case DLInformationTransfer_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
