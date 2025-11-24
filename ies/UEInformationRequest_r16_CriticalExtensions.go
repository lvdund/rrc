package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UEInformationRequest_r16_CriticalExtensions_Choice_nothing uint64 = iota
	UEInformationRequest_r16_CriticalExtensions_Choice_UeInformationRequest_r16
	UEInformationRequest_r16_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type UEInformationRequest_r16_CriticalExtensions struct {
	Choice                   uint64
	UeInformationRequest_r16 *UEInformationRequest_r16_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *UEInformationRequest_r16_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UEInformationRequest_r16_CriticalExtensions_Choice_UeInformationRequest_r16:
		if err = ie.UeInformationRequest_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode UeInformationRequest_r16", err)
		}
	case UEInformationRequest_r16_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UEInformationRequest_r16_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UEInformationRequest_r16_CriticalExtensions_Choice_UeInformationRequest_r16:
		ie.UeInformationRequest_r16 = new(UEInformationRequest_r16_IEs)
		if err = ie.UeInformationRequest_r16.Decode(r); err != nil {
			return utils.WrapError("Decode UeInformationRequest_r16", err)
		}
	case UEInformationRequest_r16_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
