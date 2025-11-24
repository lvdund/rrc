package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UECapabilityInformation_CriticalExtensions_Choice_nothing uint64 = iota
	UECapabilityInformation_CriticalExtensions_Choice_UeCapabilityInformation
	UECapabilityInformation_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type UECapabilityInformation_CriticalExtensions struct {
	Choice                   uint64
	UeCapabilityInformation  *UECapabilityInformation_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *UECapabilityInformation_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UECapabilityInformation_CriticalExtensions_Choice_UeCapabilityInformation:
		if err = ie.UeCapabilityInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode UeCapabilityInformation", err)
		}
	case UECapabilityInformation_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UECapabilityInformation_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UECapabilityInformation_CriticalExtensions_Choice_UeCapabilityInformation:
		ie.UeCapabilityInformation = new(UECapabilityInformation_IEs)
		if err = ie.UeCapabilityInformation.Decode(r); err != nil {
			return utils.WrapError("Decode UeCapabilityInformation", err)
		}
	case UECapabilityInformation_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
