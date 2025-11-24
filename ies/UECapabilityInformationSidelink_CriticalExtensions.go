package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UECapabilityInformationSidelink_CriticalExtensions_Choice_nothing uint64 = iota
	UECapabilityInformationSidelink_CriticalExtensions_Choice_UeCapabilityInformationSidelink_r16
	UECapabilityInformationSidelink_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type UECapabilityInformationSidelink_CriticalExtensions struct {
	Choice                              uint64
	UeCapabilityInformationSidelink_r16 *UECapabilityInformationSidelink_r16_IEs
	CriticalExtensionsFuture            interface{} `madatory`
}

func (ie *UECapabilityInformationSidelink_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UECapabilityInformationSidelink_CriticalExtensions_Choice_UeCapabilityInformationSidelink_r16:
		if err = ie.UeCapabilityInformationSidelink_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode UeCapabilityInformationSidelink_r16", err)
		}
	case UECapabilityInformationSidelink_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UECapabilityInformationSidelink_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UECapabilityInformationSidelink_CriticalExtensions_Choice_UeCapabilityInformationSidelink_r16:
		ie.UeCapabilityInformationSidelink_r16 = new(UECapabilityInformationSidelink_r16_IEs)
		if err = ie.UeCapabilityInformationSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode UeCapabilityInformationSidelink_r16", err)
		}
	case UECapabilityInformationSidelink_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
