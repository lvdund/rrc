package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ULInformationTransferIRAT_r16_CriticalExtensions_Choice_nothing uint64 = iota
	ULInformationTransferIRAT_r16_CriticalExtensions_Choice_C1
	ULInformationTransferIRAT_r16_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type ULInformationTransferIRAT_r16_CriticalExtensions struct {
	Choice                   uint64
	C1                       *ULInformationTransferIRAT_r16_CriticalExtensions_C1
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *ULInformationTransferIRAT_r16_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ULInformationTransferIRAT_r16_CriticalExtensions_Choice_C1:
		if err = ie.C1.Encode(w); err != nil {
			err = utils.WrapError("Encode C1", err)
		}
	case ULInformationTransferIRAT_r16_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ULInformationTransferIRAT_r16_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ULInformationTransferIRAT_r16_CriticalExtensions_Choice_C1:
		ie.C1 = new(ULInformationTransferIRAT_r16_CriticalExtensions_C1)
		if err = ie.C1.Decode(r); err != nil {
			return utils.WrapError("Decode C1", err)
		}
	case ULInformationTransferIRAT_r16_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
