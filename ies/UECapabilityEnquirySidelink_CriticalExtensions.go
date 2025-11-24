package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UECapabilityEnquirySidelink_CriticalExtensions_Choice_nothing uint64 = iota
	UECapabilityEnquirySidelink_CriticalExtensions_Choice_UeCapabilityEnquirySidelink_r16
	UECapabilityEnquirySidelink_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type UECapabilityEnquirySidelink_CriticalExtensions struct {
	Choice                          uint64
	UeCapabilityEnquirySidelink_r16 *UECapabilityEnquirySidelink_r16_IEs
	CriticalExtensionsFuture        interface{} `madatory`
}

func (ie *UECapabilityEnquirySidelink_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UECapabilityEnquirySidelink_CriticalExtensions_Choice_UeCapabilityEnquirySidelink_r16:
		if err = ie.UeCapabilityEnquirySidelink_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode UeCapabilityEnquirySidelink_r16", err)
		}
	case UECapabilityEnquirySidelink_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UECapabilityEnquirySidelink_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UECapabilityEnquirySidelink_CriticalExtensions_Choice_UeCapabilityEnquirySidelink_r16:
		ie.UeCapabilityEnquirySidelink_r16 = new(UECapabilityEnquirySidelink_r16_IEs)
		if err = ie.UeCapabilityEnquirySidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode UeCapabilityEnquirySidelink_r16", err)
		}
	case UECapabilityEnquirySidelink_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
