package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UEAssistanceInformation_CriticalExtensions_Choice_nothing uint64 = iota
	UEAssistanceInformation_CriticalExtensions_Choice_UeAssistanceInformation
	UEAssistanceInformation_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type UEAssistanceInformation_CriticalExtensions struct {
	Choice                   uint64
	UeAssistanceInformation  *UEAssistanceInformation_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *UEAssistanceInformation_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UEAssistanceInformation_CriticalExtensions_Choice_UeAssistanceInformation:
		if err = ie.UeAssistanceInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode UeAssistanceInformation", err)
		}
	case UEAssistanceInformation_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UEAssistanceInformation_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UEAssistanceInformation_CriticalExtensions_Choice_UeAssistanceInformation:
		ie.UeAssistanceInformation = new(UEAssistanceInformation_IEs)
		if err = ie.UeAssistanceInformation.Decode(r); err != nil {
			return utils.WrapError("Decode UeAssistanceInformation", err)
		}
	case UEAssistanceInformation_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
