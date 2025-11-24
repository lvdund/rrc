package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SecurityModeComplete_CriticalExtensions_Choice_nothing uint64 = iota
	SecurityModeComplete_CriticalExtensions_Choice_SecurityModeComplete
	SecurityModeComplete_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type SecurityModeComplete_CriticalExtensions struct {
	Choice                   uint64
	SecurityModeComplete     *SecurityModeComplete_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *SecurityModeComplete_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SecurityModeComplete_CriticalExtensions_Choice_SecurityModeComplete:
		if err = ie.SecurityModeComplete.Encode(w); err != nil {
			err = utils.WrapError("Encode SecurityModeComplete", err)
		}
	case SecurityModeComplete_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SecurityModeComplete_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SecurityModeComplete_CriticalExtensions_Choice_SecurityModeComplete:
		ie.SecurityModeComplete = new(SecurityModeComplete_IEs)
		if err = ie.SecurityModeComplete.Decode(r); err != nil {
			return utils.WrapError("Decode SecurityModeComplete", err)
		}
	case SecurityModeComplete_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
