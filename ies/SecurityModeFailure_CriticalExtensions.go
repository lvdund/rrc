package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SecurityModeFailure_CriticalExtensions_Choice_nothing uint64 = iota
	SecurityModeFailure_CriticalExtensions_Choice_SecurityModeFailure
	SecurityModeFailure_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type SecurityModeFailure_CriticalExtensions struct {
	Choice                   uint64
	SecurityModeFailure      *SecurityModeFailure_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *SecurityModeFailure_CriticalExtensions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SecurityModeFailure_CriticalExtensions_Choice_SecurityModeFailure:
		if err = ie.SecurityModeFailure.Encode(w); err != nil {
			err = utils.WrapError("Encode SecurityModeFailure", err)
		}
	case SecurityModeFailure_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SecurityModeFailure_CriticalExtensions) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SecurityModeFailure_CriticalExtensions_Choice_SecurityModeFailure:
		ie.SecurityModeFailure = new(SecurityModeFailure_IEs)
		if err = ie.SecurityModeFailure.Decode(r); err != nil {
			return utils.WrapError("Decode SecurityModeFailure", err)
		}
	case SecurityModeFailure_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
